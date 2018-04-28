package gj

import (
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
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

// Pings to the GJ server
// It is adviced to do this every 30 seconds
// Please note that it might be possible this is brought under a "go" sequence in future versions, and by that time this will be an empty function to prevent errors. This is however not a promise ;)
func Ping(){
	}


func gjrequest(querystring string) map[string] string{
	// https://api.gamejolt.com/api/game/v1/data-store/?game_id=32&key=test&signature=912ec803b2ce49e4a541068d495ab570
	ng:=netget("https://api.gamejolt.com/api/game/v1/data-store/?"+querystring) //game_id=32&key=test&signature=912ec803b2ce49e4a541068d495ab570
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
