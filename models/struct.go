package models

/*Panda struct*/
type Panda struct {
	PandaIndex uint32 `sql:"pandaIndex" json:"pandaIndex"`
	Genes      []byte `sql:"genes" json:"genes"`
	Birthtime  uint64 `sql:"birthtime" json:"birthtime"`
	Cooldown   uint64 `sql:"cooldown" json:"cooldown"`
	Rank       uint32 `sql:"rank" json:"rank"`
	MotherID   uint32 `sql:"motherID" json:"motherID"`
	FatherID   uint32 `sql:"fatherID" json:"fatherID"`
	Generation uint16 `sql:"generation" json:"generation"`
	Owner      []byte `sql:"owner" json:"owner"`
	Ownername  string `sql:"ownername" json:"ownername"`
	Photourl   string `sql:"photourl" json:"photourl"`
	Snapurl    string `sql:"sanpurl" json:"sanpurl"`
}
