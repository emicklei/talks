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

Software Artist, cloudfork.com

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
# mission - basics

![height:100px](/img/gi-logo.png)

- interpret Go programs using latest SDK (1.26+)
- embedded use (plugins)
- support Go modules
- provide **DAP** (Debug Adapter Protocol) for debugging

---
# mission - advanced

- code modification **during debug session**
  - change function
  - change struct type
  - add package variable | const
  - drop stack frame + resume

---
Source -> executable datastructure, the call graph

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
# Go language feature challenges

- const
- iota
- range
- break

Interpreted types and interfaces

- struct
- method
- as argument in reflection call
- callback from reflection func

---
# const

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
<button data-header-id="const">run with gi</button>
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

Translate to a simple `for` loop and re-use `ForStmt` ast node type

[call graph range string](/img/TestRangeOfString.dot.svg)

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

---
# Interpreted types

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
<button data-header-id="interpreted-types">run with gi</button>
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

# Struct types

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
# Unit testing by composing AST

![height:500px center](/img/test_by_ast.png)

---
# Test abstraction

- not affected by changes in design or implementation
- easy to create new test variations
- easy to compare output of `gi` with `go`

![height:300px center](/img/test_by_program.png)

---
![height:600px center](/img/testMain.png)

---
## plus

Debugging an interpreter that works in the "Reflection space"

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
method 0:  using `delve` in Zed

![height:500px center](/img/debug_operand_reflect.png)

---
method 1:  trace = true

![height:500px center](/img/debug_with_trace.png)

---
method 2:  visualize call graph

![height:500px center](/img/TestFunc.dot.svg)

---
method 3:  use `structexplorer`

![height:500px center](/img/explore_TestFunc.png)

---
method 4:  step through the code with `gi step`

# demo

---
method 4:  step through the code with `gi step`

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

