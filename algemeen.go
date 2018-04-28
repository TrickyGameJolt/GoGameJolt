package gj

import "fmt"

var debug = true

func chat(msg string){
    fmt.PrintLn("DEBUG>\t",debug)
}

func init(){
   chat("GJ THERE!")
}
