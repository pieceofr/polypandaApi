package models

/*Panda struct*/
type Panda struct {
	PandaIndex uint32 `sql:"pandaIndex`
	Genes      []byte `sql:"genes"`
	Birthtime  uint64 `sql:"birthtime"`
	Cooldown   uint64 `sql:"cooldown"`
	Rank       uint32 `sql:"rank"`
	MotherID   uint32 `sql:"motherID"`
	FatherID   uint32 `sql:"fatherID"`
	Generation uint16 `sql:"generation"`
	Owner      []byte `sql:"owner"`
	Ownername  string `sql:"ownername"`
	Photourl   string `sql:"photourl"`
	Snapurl    string `sql:"sanpurl"`
}
