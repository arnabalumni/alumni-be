package server

import (
	"ausAlumniServer/internal/types"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) AddAlumni(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Alumni API called")
	var fullStudentDetails types.FullStudentDetails

	err := json.NewDecoder(r.Body).Decode(&fullStudentDetails)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Request body incompatible", http.StatusBadRequest)
		return
	}

	err = s.db.AddAlumni(fullStudentDetails)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Database error", http.StatusBadGateway)
		return
	}

	resp := make(map[string]string)
	resp["message"] = "Alumni successfully added"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("error handling JSON marshal. Err: %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	_, _ = w.Write(jsonResp)
}
