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
	resp,e:=http.Get("http://utbbs.tbbs.nl")
	//if e!=nil { fmt.Println("ERROR!!!",e.Error()) } else {fmt.Println(c)}
	if e!=nil {goerr(e.Error()); return ""}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil { goerr(err.Error()); return "" }
	return string(body)
}



func gjrequest(action,querystring string) map[string] string{
	// https://api.gamejolt.com/api/game/v1/data-store/?game_id=32&key=test&signature=912ec803b2ce49e4a541068d495ab570
	ng:=netget("https://api.gamejolt.com/api/game/v1/"+action+"/?"+querystring) //game_id=32&key=test&signature=912ec803b2ce49e4a541068d495ab570
	lines:=strings.Split(ng,"\n")
	ret:=map[string] string{}
	for li,ln := range lines{
		if ln!="" {
			vr:=strings.Split(ln,":")
			if len(vr)!=2 {
				myerr(fmt.Sprintf("Game Jolt Parse error in line %d",li))
			} else {
				vr[1] = strings.Replace(vr[1], "\"", "", -1)
				ret[vr[0]]=vr[1]
			}
		}
	}
	if ret["success"]!="true" { gjerr(ret["message"]) }
	return ret
}

func (self *GJUser) qreq(action,querystring string) map[string] string{
	return gjrequest(action,querystring+self.idstring+self.gamestuff)
}



// https://gist.github.com/sergiotapia/8263278
func getMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}
