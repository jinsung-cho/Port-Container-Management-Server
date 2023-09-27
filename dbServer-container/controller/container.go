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

	rows, err := db.Query("SELECT id, inspEqNo, cntrNo, truckNo FROM Information WHERE cntrNo = $1", cntrNoHeader)
	if util.CheckHttpError(w, err, "Check DB Connection") {
		return
	}

	var informations []Information
	for rows.Next() {
		var info Information
		err = rows.Scan(&info.ID, &info.InspEqNo, &info.CntrNo, &info.TruckNo)
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

	rows, err := db.Query("SELECT id, inspEqNo, cntrNo, truckNo FROM Information")
	if util.CheckHttpError(w, err, "Check DB Connection") {
		return
	}

	var informations []Information
	for rows.Next() {
		var info Information
		err = rows.Scan(&info.ID, &info.InspEqNo, &info.CntrNo, &info.TruckNo)
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
	fmt.Fprintf(w, "CreateContainer")
}

func GetAllContainerSpec(w http.ResponseWriter, r *http.Request) {
	db := model.DBConn()

	rows, err := db.Query("SELECT * FROM Spec")
	if util.CheckHttpError(w, err, "Check DB Connection") {
		return
	}
	defer rows.Close()

	var specs []Spec
	for rows.Next() {
		var spec Spec
		err = rows.Scan(&spec.ID, &spec.InspEqNo, &spec.InspNo, &spec.InspStartTime, &spec.InspEndTime, &spec.PckMatch, &spec.InspRsltCD, &spec.DetectionCnt, &spec.FaultCD, &spec.InspRsltImgDir)
		if util.CheckHttpError(w, err, "Check DB Scan") {
			return
		}
		specs = append(specs, spec)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(specs)
}

func GetContainerReview(w http.ResponseWriter, r *http.Request) {
	inspNoHeader := r.Header.Get("inspNo")
	db := model.DBConn()

	rows, err := db.Query("SELECT r.remarkId AS remark_id, r.inspRemark, r.informationId FROM Spec s JOIN Remarks r ON s.id = r.informationId WHERE s.inspNo = $1;", inspNoHeader)
	if util.CheckHttpError(w, err, "Check DB Connection") {
		return
	}
	defer rows.Close()

	var remarks []Remark
	for rows.Next() {
		var r Remark
		err = rows.Scan(&r.RemarkID, &r.InspRemark, &r.InformationID)
		if util.CheckHttpError(w, err, "Check DB Scan") {
			return
		}
		remarks = append(remarks, r)
	}
	err = rows.Err()
	if util.CheckHttpError(w, err, "Check DB rows") {
		return
	}

	if len(remarks) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "No Reviews"}`))
		return
	}

	jsonResponse, err := json.Marshal(remarks)
	if util.CheckHttpError(w, err, "Check JSON") {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func AppendContainerReview(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "AppendContainerReview")
}