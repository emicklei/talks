# todo

## basics

- recover
- Go modules
- Call SDK functions with **interpreted type instances**
- **Callback** from SDK function to methods of interpreted type instances
- Interpreted interfaces

## advanced
 
 - all of it

 ---
 Call SDK function with interpreted type instance
 
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
 
 Callback from SDK function to interpreted type instance
 
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
