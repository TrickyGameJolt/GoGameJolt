package gj


// This type is used to log in a user, and also contains the methods to send all the data a user can send or receive from Game Jolt through this API.
type GJUser struct{
	userid string
	token string
	gameid string
	gamesig string
	// contains 'true' if succesfully logged in.
	LoggedIn bool
	idstring string 
	gamestuff string
}



// Logs a user into a game and returns the user record
func Login(gameid,privatekey,username,token string) *GJUser {
	ret:=&GJUser{}
	ret.userid=username
	ret.token=token
	ret.gameid=gameid
	ret.gamesig=getMD5Hash(privatekey)
	ret.idstring = "&username="+username+"&user_token="+token
	ret.gamestuff = "&game_id="+gameid+"&signature="+ret.gamesig
	d:=ret.qreq("users/auth","")
	ret.LoggedIn = d["success"]=="true"
	return ret
}

