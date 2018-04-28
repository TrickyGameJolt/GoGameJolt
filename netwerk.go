package gj

import (
	"net/http"
	"io/ioutil"
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
