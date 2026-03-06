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
  color: blue;
  border: 1px solid white;
  border-radius: 5px;
  padding: 5px 10px;
  cursor: pointer;
  font-size: 18px;
}
.demo {
}
.output {
  background: #003A5B;
  color: lightgreen;
  padding: 10px;  
  border-radius: 5px;
  font-family: monospace;
  font-size: 16px;
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
    -  **go-restful** - REST-ful framework (part of k8s)
    - **melrōse** - Music programming
    - **proto** - ProtocolBuffers parser
    - **dot** - Graphviz DOT writer
    - **pgtalk** - Postgres access code generator

---
# mission - basics

- interpret Go programs using latest SDK
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
# iota


<table class="demo"><tr><td>

```
  package main
  
  const (
      a = iota
      b
      c
  )
  
  func main() {
      println(a, b, c)
  }
``` 

</td><td class="output" width="30%">
1 2 3
</td></tr></table>

<div align="right"><button onClick="runGi()">run with gi</button></div>

---
# range

- slice, map
- number
- strings
- chan

---
```
package main

func main() {
    for i, v := range []string{"a", "b", "c"} {
        println(i, v)
    }

    for k, v := range map[string]int{"a": 1, "b": 2} {
        println(k, v)
    }

    for i := range 3 {
        println(i)
    }

    for i, c := range "gi" {
        println(i, c)
    }

    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)

    for v := range ch {
        println(v)
    }
}
```

---
# status

---
create a test abstraction that does the heavy lifting of the test

---
![height:600px center](/img/testMain.png)

---

![height:400px center](/img/test_by_program.png)

---
debugging an interpreter that works in the "Reflection space"

---

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

---
![height:400px center](/img/debug_trace_TestFunc.png)

---
method 0:  using delve

![height:500px center](/img/debug_operand_reflect.png)

---
method 1:  tracing statements

![height:500px center](/img/debug_with_trace.png)

---
method 2:  visualize call graph

![height:500px center](/img/TestFunc.dot.svg)

---
method 3:  use structexplorer

![height:500px center](/img/explore_TestFunc.png)

---
method 4:  step through the code with a gi debugger

![height:500px center](/img/gi_step_TestFunc.png)


---
demo `gi step`

