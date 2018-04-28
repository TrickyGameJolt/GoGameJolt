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

func (me *GJUser) OpenSession(){
	me.qreq("sessions/close","")
}
