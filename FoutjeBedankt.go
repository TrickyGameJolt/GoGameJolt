package gj

import(
	"fmt"
)


// Used to store the registered error function
type Erf struct{
	// This field must contain the function in question
	EFUN func(reportedby,errmessage string)
}

var cErf,dErf *Erf

// By default the GJ API will just print the error onto the output with a simple fmt.Print kind of command.
// However if you have a full graphic game or GUI based application, you may not want to go down that road.
// Here you can register an alternate function to throw the error to the user.
// Please note, the API will NOT crash anything. It will only display the error and then resume the program like nothing happened.
func RegError(verf *Erf){
	if verf==nil { algemene_fout("This shitty API library","Hey you! What are you trying to do? Trying to crash this innocent API library? Haven't you got anything better to do? No nils for error function registration, please!"); return }
	cErf = verf
}

// Set the errorhanding back to default
func DefaultError(){
	RegError(dErf)
}

// Returns the current error routine
func CurrentError() *Erf{
	return cErf
}


func algemene_fout(reportedby,message string) { cErf.EFUN(reportedby,message) }

func gjerr(message string) { algemene_fout("Game Jolt",message) }
func goerr(message string) { algemene_fout("Go Language",message) } 

func init(){
	RegError(&Erf{func(reportedby,errmessage string){
		fmt.Sprint("GAME JOLT API ERROR!\nReported by: %s\nError message: %s\n",reportedby,errmessage)
	}})
	dErf = cErf
}

