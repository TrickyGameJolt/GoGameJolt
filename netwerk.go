/*
  netwerk.go
  Go Game Jolt -- Network calls
  version: 18.04.29
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

import (
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

func netget(url string) string{
	chat("netget("+url+")")
	resp,e:=http.Get(url)
	//if e!=nil { fmt.Println("ERROR!!!",e.Error()) } else {fmt.Println(c)}
	if e!=nil {goerr(e.Error()); return ""}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil { goerr(err.Error()); return "" }
	return string(body)
}



func gjrequest(action,querystring,privatekey string) map[string] string{
	// https://api.gamejolt.com/api/game/v1/data-store/?game_id=32&key=test&signature=912ec803b2ce49e4a541068d495ab570
	url:="https://api.gamejolt.com/api/game/v1/"+action+"/?"+querystring
	url+="&signature="+getMD5Hash(url+privatekey)
	ng:=netget(url) //game_id=32&key=test&signature=912ec803b2ce49e4a541068d495ab570
	chat("GJ returned:\n"+ng+"\nEND RETURN")
	lines:=strings.Split(ng,"\n")
	ret:=map[string] string{}
	for li,ln := range lines{
		if ln!="" {
			vr:=strings.Split(ln,":")
			if len(vr)!=2 {
				myerr(fmt.Sprintf("Game Jolt Parse error in line %d",li))
			} else {
				value:= myTrim(strings.Replace(vr[1], "\"", "", -1))
				key:=vr[0]
				kid:=0
				_,ok:=ret[key]
				for ok { 
					kid++
					key=fmt.Sprintf("%s%d",vr[0],kid)
					_,ok=ret[key]
				}
				ret[key]=value
			}
		}
	}
	if debug {
		for k,v := range ret {
			chat(fmt.Sprintf("\t%s = '%s'",k,v))
		}
	}

	if ret["success"]!="true" { chat("Dit ging niet goed"); gjerr(ret["message"]) }
	return ret
}

func (self *GJUser) qreq(action,querystring string) map[string] string{
	return gjrequest(action,querystring+self.idstring+self.gamestuff,self.gamekey)
}



// https://gist.github.com/sergiotapia/8263278
func getMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}
