package models

type Acc struct {
	Name    string `json:"name"`
	NotKill bool   `json:"not_kill"`
}
type PayLoadWSS struct {
	MessPre string `json:"mess_pre"`
	Mess    *Mess  `json:"mess"`
	ListAcc []*Acc `json:"list_acc"`
	AccNext string `json:"acc_next"`
}
