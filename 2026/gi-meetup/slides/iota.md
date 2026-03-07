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
