package models

import (
	"encoding/json"
)

const (
	statusOK    = "OK"
	statusFail  = "Fail"
	okMessage   = "Operation is Successful"
	failMessage = "Operation is Fail"
)

/*RetSimple for Return Message*/
type RetSimple struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Extra   string `json:"extra, omitempty"`
}

/*StatusOK make a Successful Return Message*/
func (r *RetSimple) StatusOK() {
	r.Status = statusOK
	r.Message = okMessage
}

/*StatusFail make a Successful Return Message*/
func (r *RetSimple) StatusFail(err string) {
	r.Status = failMessage
	if len(err) == 0 {
		r.Message = failMessage
	} else {
		r.Message = err
	}
}

/*GetJSONStream convert Panda array to Json*/
func (r *RetSimple) GetJSONStream() []byte {
	if jsonData, err := json.Marshal(r); err == nil {
		return jsonData
	}
	return []byte{}
}

/*RetPanda for Return Message*/
type RetPanda struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Extra   string  `json:"extra, omitempty"`
	Pandas  []Panda `json:"pandas, omitempty"`
}

/*StatusOK make a Successful Return Message*/
func (r *RetPanda) StatusOK(p []Panda) {
	r.Status = statusOK
	r.Message = okMessage
	r.Pandas = p
}

/*StatusFail make a Successful Return Message*/
func (r *RetPanda) StatusFail(err string) {
	r.Status = failMessage
	if len(err) == 0 {
		r.Message = failMessage
	} else {
		r.Message = err
	}
}

/*IsOK is status equal OK*/
func (r *RetPanda) IsOK() bool {
	if r.Status == statusOK {
		return true
	}
	return false
}

/*GetJSONStream convert Panda array to Json*/
func (r *RetPanda) GetJSONStream() []byte {
	if jsonData, err := json.Marshal(r); err == nil {
		return jsonData
	}
	return []byte{}
}
