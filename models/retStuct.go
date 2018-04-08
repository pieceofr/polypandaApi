package models

import (
	"encoding/json"
	"strconv"
)

const (
	statusFail  = "Fail"
	okMessage   = "200 OK"
	failMessage = "Operation Fail"
)

/*RetSimple for Return Message*/
type RetSimple struct {
	Code      StatusCode `json:"code"`
	Message   string     `json:"message"`
	ExtraCode int        `json:"extraCode, omitempty"`
	Extra     string     `json:"extra, omitempty"`
}

/*StatusOK make a Successful Return Message*/
func (r *RetSimple) StatusOK() {
	r.Code = St200OK
	r.Message = okMessage
}

/*StatusFail make a Successful Return Message*/
func (r *RetSimple) StatusFail(code StatusCode) {
	r.Code = code
	r.Message = "fail code:" + strconv.Itoa(int(code))
}

/*SetStatus set status of RetPanda, extracode = 0 to omit the attribute */
func (r *RetSimple) SetStatus(code StatusCode, msg string, extracode int) {
	r.Code = code
	r.Message = msg
	if extracode != 0 {
		r.ExtraCode = extracode
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
	Code      StatusCode `json:"code"`
	Message   string     `json:"message"`
	ExtraCode int        `json:"extraCode, omitempty"`
	Extra     string     `json:"extra, omitempty"`
	Pandas    []Panda    `json:"pandas, omitempty"`
}

/*StatusOK make a Successful Return Message*/
func (r *RetPanda) StatusOK(p []Panda) {
	r.Code = St200OK
	r.Message = okMessage
	r.Pandas = p
}

/*StatusFail make a Successful Return Message*/
func (r *RetPanda) StatusFail(code StatusCode) {
	r.Code = code
	r.Message = "fail code:" + strconv.Itoa(int(code))
}

/*SetStatus set status of RetPanda, extracode = 0 to omit the attribute */
func (r *RetPanda) SetStatus(code StatusCode, msg string, extracode int) {
	r.Code = code
	r.Message = msg
	if extracode != 0 {
		r.ExtraCode = extracode
	}
}

/*GetJSONStream convert Panda array to Json*/
func (r *RetPanda) GetJSONStream() []byte {
	if jsonData, err := json.Marshal(r); err == nil {
		return jsonData
	}
	return []byte{}
}
