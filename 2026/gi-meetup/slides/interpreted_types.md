# Interpreted types

---
# Struct

```
package main

type Aircraft struct {
	Model string
}
func main() {
	heli := Aircraft{Model:"helicopter"}
	print(heli.Model)
}
```
<div align="right">
<button data-header-id="interpreted-types">run with gi</button>
</div>

---

# Underlying literal type

```
package main

type Count int

func main() {
	one := Count(1)
	print(one)
}
```
<div align="right">
<button data-header-id="underlying-literal-type">run with gi</button>
</div>

---

# Alias literal type

```
package main

type Count = int

func main() {
	one := Count(1)
	print(one)
}
```
<div align="right">
<button data-header-id="alias-literal-type">run with gi</button>
</div>

---

# Struct types in `gi`

> type Aircraft struct

```
type StructType struct {
	name      string
	fields    *FieldList // for instantiation
	methods   map[string]*FuncDecl
	embeds    []StructType
}
```

and instances thereof

```
type StructValue struct {
	structType StructType
	fields     map[string]reflect.Value
}
```

--- 

# Extension types 

> type Count int

```
type ExtendedType struct {
	name    Ident
	methods map[string]*FuncDecl
}
```
and instances thereof
```
type ExtendedValue struct {
	typ ExtendedType  // The typ is used for method resolution.
	val reflect.Value // The val field holds the actual reflect.Value.
}
```

---
 # Call SDK function with interpreted type instance
 
 ```
 type Aircraft struct {
    Name string
 }
 
 func main() {
     a := &Aircraft{Name: "Boeing 747"}
     fmt.Printf("%v", a)
 }
 ```
 ---
 
 # Callback from SDK function to interpreted type instance
 
 ```
 type bookReader struct {}
 
 func (b *bookReader) Read(p []byte) (n int, err error) {
     // ...
     return 0, nil
 }
 
  func main() {
      var r io.Reader = new(bookReader)
      _ = io.ReadAll(r)
      
  }
  ```
