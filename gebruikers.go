/*
  gebruikers.go
  Go Game Jolt -- User Management
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

//
package gj


// This type is used to log in a user, and also contains the methods to send all the data a user can send or receive from Game Jolt through this API.
type GJUser struct{
	userid string
	token string
	gameid string
	gamekey string
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
	ret.gamekey=privatekey //getMD5Hash(privatekey)
	ret.idstring = "&username="+username+"&user_token="+token
	ret.gamestuff = "&game_id="+gameid //+"&signature="+ret.gamesig
	d:=ret.qreq("users/auth","")
	ret.LoggedIn = d["success"]=="true"
	return ret
}

