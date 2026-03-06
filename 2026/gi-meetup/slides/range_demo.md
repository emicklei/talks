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
