package gj

import "fmt"

var debug = true

func chat(msg string){
    fmt.Println("DEBUG>\t",msg)
}

func init(){
   chat("GJ THERE!")
}
