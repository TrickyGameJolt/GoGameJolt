package gj

// Pings to the GJ server
// It is adviced to do this every 30 seconds
// Please note that it might be possible this is brought under a "go" sequence in future versions, and by that time this will be an empty function to prevent errors. This is however not a promise ;)
func (me *GJUser) Ping(){
	me.qreq("sessions/ping","")
	}

