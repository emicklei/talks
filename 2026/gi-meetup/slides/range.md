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

---

```
func (r RangeStmt) chanFlow(g *graphBuilder) (head Step) {	
	cond := BinaryExpr{ // len(ch) > 0
		op:    token.GTR,
		opPos: r.forPos,
		x:     reflectLenExpr{X: r.x},
		y:     newBasicLit(r.pos(), reflect.ValueOf(0)),
	}
	recv := UnaryExpr{ // <- chan
		opPos: r.pos(),
		op:    token.ARROW,
		x:     r.x,
	}	
	ass := AssignStmt{ // var := <- chan
		tok:    token.DEFINE,
		tokPos: r.pos(),
		lhs:    []Expr{r.key},
		rhs:    []Expr{recv},
	}
	bodyList := append([]Stmt{ass}, r.body.list...)
	body := &BlockStmt{lbracePos: r.body.pos(), list: bodyList}
	return ForStmt{forPos: r.pos(), cond: cond, body: body}.flow(g)
}
```