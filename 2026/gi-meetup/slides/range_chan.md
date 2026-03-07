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
