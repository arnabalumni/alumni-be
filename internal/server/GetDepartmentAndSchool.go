package server

import (
	"ausAlumniServer/internal/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *Server) GetDepartmentAndSchool(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get School API called")

	var deptData types.DeptData
	err := json.NewDecoder(r.Body).Decode(&deptData)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Request body incompatible", http.StatusBadRequest)
		return
	}

	deptAndSchoolName, err := s.db.GetDepartmentAndSchoolQuery(deptData.DeptId)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Database query failed", http.StatusBadRequest)
		return
	}

	jsonResp, err := json.Marshal(deptAndSchoolName)
	if err != nil {
		log.Printf("Error handling JSON marshal: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// fmt.Println(deptAndSchoolName)
	for key, value := range deptAndSchoolName {
		fmt.Printf("%s: %s\n", key, value)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
