package controllers

import (
	"encoding/json"
	"polypanda/apiserver/models"
)

const (
	statusOK    = "OK"
	statusFail  = "Fail"
	okMessage   = "Operation is Successful"
	failMessage = "Operation is Fail"
)

/*RetJSON the type that support Json Format*/
type RetJSON interface {
	statusOK()
	statusFail(err string)
	getJSONStream()
}

/*RetSimple for Return Message*/
type RetSimple struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Extra   string `json:"extra, omitempty"`
}

/*StatusOK make a Successful Return Message*/
func (r *RetSimple) statusOK() {
	r.Status = statusOK
	r.Message = okMessage
}

/*StatusFail make a Successful Return Message*/
func (r *RetSimple) statusFail(err string) {
	r.Status = failMessage
	if len(err) == 0 {
		r.Message = failMessage
	} else {
		r.Message = err
	}
}

/*PandgetJSONStream convert Panda array to Json*/
func (r *RetSimple) getJSONStream() []byte {
	if jsonData, err := json.Marshal(r); err == nil {
		return jsonData
	}
	return []byte{}
}

/*RetPanda for Return Message*/
type RetPanda struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Extra   string         `json:"extra, omitempty"`
	Pandas  []models.Panda `json:"pandas, omitempty"`
}

/*StatusOK make a Successful Return Message*/
func (r *RetPanda) statusOK(p []models.Panda) {
	r.Status = statusOK
	r.Message = okMessage
	r.Pandas = p
}

/*StatusFail make a Successful Return Message*/
func (r *RetPanda) statusFail(err string) {
	r.Status = failMessage
	if len(err) == 0 {
		r.Message = failMessage
	} else {
		r.Message = err
	}
}

func (r *RetPanda) isOK() bool {
	if r.Status == statusOK {
		return true
	}
	return false
}

/*PandgetJSONStream convert Panda array to Json*/
func (r *RetPanda) getJSONStream() []byte {
	if jsonData, err := json.Marshal(r); err == nil {
		return jsonData
	}
	return []byte{}
}
