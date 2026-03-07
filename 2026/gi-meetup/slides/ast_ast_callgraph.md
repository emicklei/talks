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

```
package main

func main() {
	print(40 + 2)
}
```

<div align="right"><button onClick="runGi()">run with gi</button></div>
