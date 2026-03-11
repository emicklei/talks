---
marp: true
theme: default
paginate: true 
color: "#EEE"
backgroundColor: "#222"
footer: <h3>github.com/emicklei/gi</h3>
header: <h4>Golang Meetup, March 2026, Gi - building a Go interpreter</h4>

---
# Gi - building a Go interpreter
March 2026, Golang Meetup Amsterdam

### Ernest Micklei

Software Artist, Engineer, Architect

<style>
pre,code {
  background: #00527A;
  color: yellow;
}
a {
  color: cyan;
}
img[alt~="center"] {
  display: block;
  margin: 0 auto;
  background-color: transparent;
}
strong {
  color: cyan;    
}
button {
  background: orange;
  color: black;
  border: 1px solid white;
  border-radius: 5px;
  padding: 5px 10px;
  cursor: pointer;
  font-size: 18px;
  font-weight: bold;
}
code[contenteditable]:focus {
  outline: 2px solid orange;
  cursor: text;
}
.output {
  background: gray;
  color: white;
  padding: 10px;  
  border: 2px solid orange;
  border-radius: 5px;
  font-family: monospace;
  font-weight: bold;
  font-size: 22px;
  vertical-align: top;
}
</style>
<script src="slides/play.js"></script>

---
# Who am i

![height:80px center](/img/emicklei_hackers_logo.png)

- using Go since 2011
- platform engineer (day) 
- open-source developer (night)
    - **go-restful** - REST-ful framework (part of k8s)
    - **melrōse** - Music programming
    - **proto** - ProtocolBuffers parser
    - **dot** - Graphviz DOT writer
    - **pgtalk** - Postgres access code generator

---
# foundation

![height:100px](/img/gi-logo.png)

- interpret Go programs using latest SDK (1.26+)
- embedded use (plugins)
- support Go modules
- provide **DAP** (Debug Adapter Protocol) for debugging

---
# mission

- code modification **during debug session**
  - change function body
  - change struct type
  - add package variable | const
  - drop stack frame + resume

---

# Source → executable datastructure, the call graph

---
Building call graphs

![height:400px center](/img/source-to-callgraphs.png)

---
Executing the call graphs

![height:400px center](/img/callgraph-vm-exec.png)

---
from Go AST to a "mirror" AST

```
package main

func main() {
	print(40 + 2)
}
```

Let us look at both **AST**s of this program:

---

![height:600px center](/img/go_ast_42.png)

---

![height:600px center](/img/gi_ast_42.png)

---

Building the call graph from the AST

<br>

![height:200px center](/img/callgraph_42.png)

---

![height:600px center](/img/flow_42.png)

---

# BinaryExpr

```
package main

func main() {
	print(40 + 2)
}
```
<div align="right">
<button data-header-id="binaryexpr">run with gi</button>
</div>

---
# Go language features

---
# Challenges

- var
- iota
- range
- break, goto, continue

Interpreted types

- struct
- method
- as argument in reflection call
- callback from reflection func

---
# var

```
package main

var (
	a = c + b  // == 9
	b = f()    // == 4
	c = f()    // == 5
	d = 3      // == 5 after initialization has finished
)

func f() int {
	d++
	return d
}
func main() {
	print(a,b,c,d)
}
```
<div align="right">
<button data-header-id="var">run with gi</button>
</div>

---
# iota

```
package main

type state int

const (
	a = iota
	b
	c       = 5
	d state = iota
	e
	f = 1
	g
)
func main() {
	print(a, b, c, d, e, f, g)
}
```
<div align="right">
<button data-header-id="iota">run with gi</button>
</div>

---
# range

- slice, array
- map
- number
- string
- chan
- iterator

---

Translate to a simple `for` loop and re-use `ForStmt` ast node type

```
for v := range ch {
    print(v)
}
```
introduce condition and assignment:

```
for len(ch) > 0 {
    v := <-ch
    print(v)
}
```

---

```
func (r RangeStmt) chanFlow(g *graphBuilder) (head Step) {	
	cond := BinaryExpr{ // len(ch) > 0
		op:    token.GTR,
		opPos: r.forPos,
		x:     reflectLenExpr{X: r.x},
		y:     newBasicLit(r.pos(), reflect.ValueOf(0)),
	}
	recv := UnaryExpr{ // <- chan
		opPos: r.pos(),
		op:    token.ARROW,
		x:     r.x,
	}	
	ass := AssignStmt{ // var := <- chan
		tok:    token.DEFINE,
		tokPos: r.pos(),
		lhs:    []Expr{r.key},
		rhs:    []Expr{recv},
	}
	bodyList := append([]Stmt{ass}, r.body.list...)
	body := &BlockStmt{lbracePos: r.body.pos(), list: bodyList}
	return ForStmt{forPos: r.pos(), cond: cond, body: body}.flow(g)
}
```

---
# range slice

```
package main

func main() {
    for i, v := range []string{"a", "b", "c"} {
        println(i, v)
    }
}
```
<div align="right">
<button data-header-id="range-slice">run with gi</button>
</div>

---
# range map

```
package main

func main() {
    for k, v := range map[string]int{"a": 1, "b": 2} {
        println(k, v)
    }
}
```
<div align="right">
<button data-header-id="range-map">run with gi</button>
</div>

---
# range string

```
package main

func main() { 
    for i, c := range "gi" {
        println(i, c)
    }
}
```
<div align="right">
<button data-header-id="range-string">run with gi</button>
</div>

[call graph range string](/img/TestRangeOfString.dot.svg)

---
# range chan

```
package main

func main() {
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)
    for v := range ch {
        print(v)
    }
}
```
<div align="right">
<button data-header-id="range-chan">run with gi</button>
</div>

[call graph range chan](/img/TestChannelRange.dot.svg)

---
# Interpreted types

---
# Struct

```
package main

type Aircraft struct {
	Model string
}
func main() {
	heli := Aircraft{Model:"helicopter"}
	print(heli.Model)
}
```
<div align="right">
<button data-header-id="struct">run with gi</button>
</div>

---

# Underlying literal type

```
package main

type Count int

func main() {
	one := Count(1)
	print(one)
}
```
<div align="right">
<button data-header-id="underlying-literal-type">run with gi</button>
</div>

---

# Alias literal type

```
package main

type Count = int

func main() {
	one := Count(1)
	print(one)
}
```
<div align="right">
<button data-header-id="alias-literal-type">run with gi</button>
</div>

---

# Struct types in `gi`

> type Aircraft struct

```
type StructType struct {
	name      string
	fields    *FieldList // for instantiation
	methods   map[string]*FuncDecl
	embeds    []StructType
}
```

and instances thereof

```
type StructValue struct {
	structType StructType
	fields     map[string]reflect.Value
}
```

--- 

# Extension types 

> type Count int

```
type ExtendedType struct {
	name    Ident
	methods map[string]*FuncDecl
}
```
and instances thereof
```
type ExtendedValue struct {
	typ ExtendedType  // The typ is used for method resolution.
	val reflect.Value // The val field holds the actual reflect.Value.
}
```

---
 # Call SDK function with interpreted type instance
 
 ```
 type Aircraft struct {
    Name string
 }
 
 func main() {
     a := &Aircraft{Name: "Boeing 747"}
     fmt.Printf("%v", a)
 }
 ```
 ---
 
 # Callback from SDK function to interpreted type instance
 
 ```
 type bookReader struct {}
 
 func (b *bookReader) Read(p []byte) (n int, err error) {
     // ...
     return 0, nil
 }
 
  func main() {
      var r io.Reader = new(bookReader)
      _ = io.ReadAll(r)
      
  }
  ```

---
# Testing

.. from unit to integration

---
# composing Go AST

![height:500px center](/img/test_by_ast.png)

---
# abstraction

- not affected by changes in design or implementation
- easy to create new test variations
- easy to compare output of `gi` with `go`

![height:300px center](/img/test_by_program.png)

---
![height:600px center](/img/testMain.png)

---
# Debugging 

.. an interpreter that operates in the "Reflection space"

---
# plus

```
package main

func plus(a int, b int) int {
	return a + b
}
func main() {
	result := plus(2, 3)
	print(result)
}
```
<div align="right">
<button data-header-id="plus">run with gi</button>
</div>

---

Breakpoint on calling the binary expression.

![height:400px center](/img/debug_trace_TestFunc.png)

---
using `delve` in Zed

![height:500px center](/img/debug_operand_reflect.png)

---
trace = true

![height:500px center](/img/debug_with_trace.png)

---
visualize call graph

![height:500px center](/img/TestFunc.dot.svg)

---
use `structexplorer`

![height:500px center](/img/explore_TestFunc.png)

---
step through the code with `gi step`

# demo

---
step through the code with `gi step`

![height:500px center](/img/gi_step_TestFunc.png)

---
# todo

## basics

- recover
- Go modules
- Call SDK functions with **interpreted type instances**
- **Callback** from SDK function to methods of interpreted type instances
- Complete interpreted types and interfaces

## advanced
 
 - all of it

---
# thank you

![height:100px](/img/gi-logo.png)

Open-source project (MIT)

**github.com/emicklei/gi**

Slide deck (Creative Commons)

**github.com/emicklei/talks**

