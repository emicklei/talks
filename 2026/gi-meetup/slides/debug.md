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