package gj

import "fmt"

var debug = true

func chat(msg){
    fmt.PrintLn("DEBUG>\t",debug)
}

func init(){
   chat("GJ THERE!")
}
