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