# range

- slice, array
- map
- number
- string
- chan
- iterator

---

Translate to a simple `for` loop and re-use `ForStmt` ast node type

```
for v := range ch {
    print(v)
}
```
introduce condition and assignment:

```
for len(ch) > 0 {
    v := <-ch
    print(v)
}
```