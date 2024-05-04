package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) GetBuildDepartmentsData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Students API called")

	// err := json.NewDecoder(r.Body).Decode(&batchDetail)
	// if err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, "Request body incompatible", http.StatusBadRequest)
	// 	return
	// }
	resp, err := s.db.BuildDepartmentsData()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal database error", http.StatusBadGateway)
		return
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("error handling JSON marshal. Err: %v", err)
		http.Error(w, "Internal JSON parsing error", http.StatusBadGateway)
		return
	}

	_, _ = w.Write(jsonResp)
}
