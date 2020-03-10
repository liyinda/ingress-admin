package main 
import ( 
    "encoding/json" 
    "fmt" 
) 
func main ( ) { 
    var jsonBlob = [ ] byte ( ` [ 
        { "Name" : "Platypus" , "Order" : "Monotremata" } , 
        { "Name" : "Quoll" ,     "Order" : "Dasyuromorphia" } 
    ] ` ) 
    type Animal struct { 
        Name  string 
        Order string 
    } 
    var animals [ ] Animal 
    err := json. Unmarshal ( jsonBlob , & animals ) 
    if err != nil { 
        fmt. Println ( "error:" , err ) 
    }
    test := string(jsonBlob) 
    fmt. Println (  test ) 
    fmt. Println (  "222222222222222" ) 
    fmt. Printf ( "%+v" , animals ) 
}
