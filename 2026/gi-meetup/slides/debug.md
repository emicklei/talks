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
