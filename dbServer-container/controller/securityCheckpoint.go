package controller

import (
	"db-server/model"
	"db-server/util"
	"encoding/json"
	"net/http"
)

func GetCheckpoint(w http.ResponseWriter, r *http.Request) {
	db := model.DBConn()

	rows, err := db.Query("SELECT id, inspEqNo, inspAuto, inspName, inspLoc, inspContact FROM EqOperateInfo")
	if util.CheckHttpError(w, err, "Check DB Connection") {
		return
	}
	defer rows.Close()

	var eqOperateInfos []EqOperateInfo
	for rows.Next() {
		var eoi EqOperateInfo
		err = rows.Scan(&eoi.ID, &eoi.InspEqNo, &eoi.InspAuto, &eoi.InspName, &eoi.InspLoc, &eoi.InspContact)
		if util.CheckHttpError(w, err, "Check DB Scan") {
			return
		}
		eqOperateInfos = append(eqOperateInfos, eoi)
	}
	err = rows.Err()
	if util.CheckHttpError(w, err, "Check DB rows") {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eqOperateInfos)
}

func GetCheckpointState(w http.ResponseWriter, r *http.Request) {
	db := model.DBConn()

	rows, err := db.Query("SELECT id, inspEqNo, inspEqStatus FROM EqStateInfo")
	if util.CheckHttpError(w, err, "Check DB Connection") {
		return
	}
	defer rows.Close()

	var eqStateInfos []EqStateInfo
	for rows.Next() {
		var esi EqStateInfo
		err = rows.Scan(&esi.ID, &esi.InspEqNo, &esi.InspEqStatus)
		if util.CheckHttpError(w, err, "Check DB Scan") {
			return
		}
		eqStateInfos = append(eqStateInfos, esi)
	}
	err = rows.Err()
	if util.CheckHttpError(w, err, "Check DB rows") {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eqStateInfos)
}
