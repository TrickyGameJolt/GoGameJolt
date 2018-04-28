/*
  FoutjeBedankt.go
  Go Game Jolt
  version: 18.04.28
  Copyright (C) 2018 Jeroen P. Broks
  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.
  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:
  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.
*/
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
func myerr(message string) { algemene_fout("Go Game Jolt",message) }

func init(){
	RegError(&Erf{func(reportedby,errmessage string){
		fmt.Sprint("GAME JOLT API ERROR!\nReported by: %s\nError message: %s\n",reportedby,errmessage)
	}})
	dErf = cErf
}

