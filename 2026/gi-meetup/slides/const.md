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