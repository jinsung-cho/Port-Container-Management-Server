package controller

type EqOperateInfo struct {
	ID          int    `json:"id"`
	InspEqNo    string `json:"inspEqNo"`
	InspAuto    string `json:"inspAuto"`
	InspName    string `json:"inspName,omitempty"`
	InspLoc     string `json:"inspLoc,omitempty"`
	InspContact string `json:"inspContact,omitempty"`
}

type EqStateInfo struct {
	ID           int    `json:"id"`
	InspEqNo     string `json:"inspEqNo"`
	InspEqStatus string `json:"inspEqStatus"`
}
