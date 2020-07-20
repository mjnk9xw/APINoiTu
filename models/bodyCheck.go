package models

type BodyCheck struct {
	PreSeq string `json:"pre_seq"`
	Seq    string `json:"seq"`
}

type BoolResponse struct {
	Data      bool   `json:"data"`
	ErrorMess string `json:"error_mess"`
}
