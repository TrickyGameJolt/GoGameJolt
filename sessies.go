/*
  sessies.go
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

// Pings to the GJ server
// It is adviced to do this every 30 seconds
// Please note that it might be possible this is brought under a "go" sequence in future versions, and by that time this will be an empty function to prevent errors. This is however not a promise ;)
func (me *GJUser) Ping(){
	me.qreq("sessions/ping","")
	}

// Opens a game session for a particular user. Allows you to tell Game Jolt that a user is playing your game. You must ping the session to keep it active and you must close it when you're done with it. Note that you can only have one open session at a time. If you try to open a new session while one is running, the system will close out your current one before opening a new one.
func (me *GJUser) OpenSession(){
	me.qreq("sessions/open","")
}

// Closes the active session.

func (me *GJUser) CloseSession(){
	me.qreq("sessions/close","")
}
