package controller

import (
	"db-server/model"
	"db-server/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetContainer(w http.ResponseWriter, r *http.Request) {
	cntrNoHeader := r.Header.Get("cntrNo")
	db := model.DBConn()

	rows, err := db.Query("SELECT id, inspEqNo, cntrNo, truckNo, typeNo, qDate FROM PreInformation WHERE cntrNo = $1", cntrNoHeader)
	if util.CheckHttpError(w, err, "Check DB Connection") {
		return
	}

	var informations []PreInformation
	for rows.Next() {
		var info PreInformation
		err = rows.Scan(&info.ID, &info.InspEqNo, &info.CntrNo, &info.TruckNo, &info.TypeNo, &info.QDate)
		if util.CheckHttpError(w, err, "Check DB Scan") {
			return
		}
		informations = append(informations, info)
	}
	err = rows.Err()
	if util.CheckHttpError(w, err, "Check DB Rows") {
		return
	}

	if len(informations) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "No matching data"}`))
		return
	}

	jsonResponse, err := json.Marshal(informations)
	if util.CheckHttpError(w, err, "Check JSON") {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetAllContainers(w http.ResponseWriter, r *http.Request) {
	db := model.DBConn()

	rows, err := db.Query("SELECT id, inspEqNo, cntrNo, truckNo, typeNo, qDate FROM PreInformation")
	if util.CheckHttpError(w, err, "Check DB Connection") {
		return
	}

	var informations []PreInformation
	for rows.Next() {
		var info PreInformation
		err = rows.Scan(&info.ID, &info.InspEqNo, &info.CntrNo, &info.TruckNo, &info.TypeNo, &info.QDate)
		if util.CheckHttpError(w, err, "Check DB Scan") {
			return
		}
		informations = append(informations, info)
	}
	err = rows.Err()
	if util.CheckHttpError(w, err, "Check DB Rows") {
		return
	}

	if len(informations) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "No matching data"}`))
		return
	}

	jsonResponse, err := json.Marshal(informations)
	if util.CheckHttpError(w, err, "Check JSON") {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func CreateContainer(w http.ResponseWriter, r *http.Request) {
	db := model.DBConn()

	var spec ContainerSpec
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&spec)
	if util.CheckHttpError(w, err, "Decoding JSON error") {
		return
	}

	query := `
		INSERT INTO ContainerSpec 
		(inspEqNo, inspNo, inspStartTime, inspEndTime, pckMatch, inspRsltCD, detectionCnt, faultCD, inspRsltImgDir, qDate) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	err = db.QueryRow(query, spec.InspEqNo, spec.InspNo, spec.InspStartTime, spec.InspEndTime, spec.PckMatch, spec.InspRsltCD, spec.DetectionCnt, spec.FaultCD, spec.InspRsltImgDir, spec.QDate).Scan(&spec.ID)
	if util.CheckHttpError(w, err, "Save DB error") {
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetAllContainerSpec(w http.ResponseWriter, r *http.Request) {
	db := model.DBConn()

	rows, err := db.Query("SELECT * FROM ContainerSpec")
	if util.CheckHttpError(w, err, "Check DB Connection") {
		return
	}
	defer rows.Close()

	var specs []ContainerSpec
	for rows.Next() {
		var spec ContainerSpec
		err = rows.Scan(&spec.ID, &spec.InspEqNo, &spec.InspNo, &spec.InspStartTime, &spec.InspEndTime, &spec.PckMatch, &spec.InspRsltCD, &spec.DetectionCnt, &spec.FaultCD, &spec.InspRsltImgDir, &spec.QDate)
		if util.CheckHttpError(w, err, "Check DB Scan") {
			return
		}
		specs = append(specs, spec)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(specs)
}
