# iota


<table class="demo"><tr><td>

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

</td><td class="output" width="30%">
1 2 3
</td></tr></table>

<div align="right"><button onClick="runGi()">run with gi</button></div>
