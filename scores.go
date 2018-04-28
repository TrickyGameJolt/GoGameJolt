/*
  scores.go
  Go Game Jolt -- Scores management
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


// Submit score as a guest
func SubmitGuestScore(guestname,gameid,privatekey,score,sort,table_id string) bool {
	qs:="score="+score+"&sort="+sort
	if table_id!="" { qs+="&table_id"+table_id }
	qs+="guest="+guestname+"&game_id="+gameid+"&signature="+getMD5Hash(privatekey)
	r:=gjrequest("scores/add",qs)
	return r["success"]=="true"
}


// Submits a score as a user.
// If you give up an empty string for table_id, Game Jolt will use the "primary" table. The only reason why you MUST at least put in something even if it's an empty string is because Go does not support optional parameters, and the dev crew also showed no interest to implement the feature (for reasons that are downright stupid if you ask me).
func (me *GJUser) SubmitScore(score,sort,table_id string) bool {
	qs:="score="+score+"&sort="+sort
	if table_id!="" { qs+="&table_id"+table_id }
	r:=me.qreq("scores/add",qs)
	return r["success"]=="true"
}



func fetchscore(username,token,gameid,limit,table_id string) map[string] string{
	qs:=""
	if username!="" { 
		qs+="username="+username+"&user_token="+token 
	}
	if limit!="" {
		if qs!="" { qs+="&" }
		qs+="limit="+limit
	}
	if table_id!="" {
		if qs!="" { qs+="&" }
		qs+="table_id="+table_id
	}
	if qs!="" { qs+="&" }
	qs+="game_id="+gameid
	return gjrequest("scores",qs)
}


// Fetches the score for a user
// Only use this if you want to fetch the scores for that specific user!
func (me *GJUser) FetchScore(limit,table_id string) map[string] string{
	return fetchscore(me.userid,me.token,me.gameid,limit,table_id)
}

// Fetches scores in general
func FetchScore(gameid,limit,table_id string) map[string] string{
	return fetchscore("","",gameid,limit,table_id)
}
