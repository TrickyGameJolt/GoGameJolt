package gj

import "fmt"

var debug = true

func chat(msg string){
    fmt.Println("DEBUG>\t",debug)
}

func init(){
   chat("GJ THERE!")
}
