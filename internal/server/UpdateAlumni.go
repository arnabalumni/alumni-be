package server

import (
	"ausAlumniServer/internal/types"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) UpdateAlumni(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Alumni API called")
	var updatedStudentDetail types.StudentDetails

	err := json.NewDecoder(r.Body).Decode(&updatedStudentDetail)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Request body incompatible", http.StatusBadRequest)
	}
	fmt.Println(updatedStudentDetail)
	err = s.db.UpdateAlumni(updatedStudentDetail)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Database error", http.StatusBadGateway)
		return
	}
	jsonResp, err := json.Marshal(updatedStudentDetail)
	if err != nil {
		fmt.Printf("error handling JSON marshal. Err: %v", err)
		http.Error(w, "Internal JSON parsing error", http.StatusBadGateway)
		return
	}
	_, _ = w.Write(jsonResp)
}
