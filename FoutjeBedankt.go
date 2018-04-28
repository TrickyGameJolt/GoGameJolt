package gj

import(
	"fmt"
)


// Used to store the registered error function
type Erf struct{
	// This field must contain the function in question
	EFUN func(reportedby,errmessage string)
}

var cErf *Erf

// By default the GJ API will just print the error onto the output with a simple fmt.Print kind of command.
// However if you have a full graphic game or GUI based application, you may not want to go down that road.
// Here you can register an alternate function to throw the error to the user.
// Please note, the API will NOT crash anything. It will only display the error and then resume the program like nothing happened.
func RegError(verf *Erf){
	cErf = verf
}


func init(){
	RegError(&Erf{func(reportedby,errmessage string){
		fmt.Sprint("GAME JOLT API ERROR!\nReported by: %s\nError message: %s\n",reportedby,errmessage)
	}})
}
