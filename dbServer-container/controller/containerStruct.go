package controller

type Information struct {
	ID       int    `json:"id"`
	InspEqNo string `json:"inspEqNo"`
	CntrNo   string `json:"cntrNo"`
	TruckNo  string `json:"truckNo"`
}

type Spec struct {
	ID             int    `json:"id"`
	InspEqNo       string `json:"inspEqNo"`
	InspNo         string `json:"inspNo"`
	InspStartTime  string `json:"inspStartTime,omitempty"`
	InspEndTime    string `json:"inspEndTime,omitempty"`
	PckMatch       string `json:"pckMatch"`
	InspRsltCD     string `json:"inspRsltCD"`
	DetectionCnt   string `json:"detectionCnt"`
	FaultCD        string `json:"faultCD"`
	InspRsltImgDir string `json:"inspRsltImgDir,omitempty"`
}

type Remark struct {
	RemarkID      int    `json:"remarkId"`
	InspRemark    string `json:"inspRemark"`
	InformationID int    `json:"informationId"`
}
