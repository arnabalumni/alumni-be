package server

import (
	"ausAlumniServer/internal/types"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) Students(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Students API called")
	// body, err := io.ReadAll(r.Body)
	// fmt.Println(string(body))
	var batchDetail types.BatchDetails

	err := json.NewDecoder(r.Body).Decode(&batchDetail)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Request body incompatible", http.StatusBadRequest)
		return
	}
	resp, err := s.db.StudentQuery(batchDetail)
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
