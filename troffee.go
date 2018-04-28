package gj

// Awards a trophy and returns "true" if succesful!
func (me *GJUser) AwardTrophy(id string) bool{
	r:=me.qreq("trophies/add-achieved","trophy_id="+id)
	return r["success"]=="true" // temp line
}

