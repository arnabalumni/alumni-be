package server

import (
	"ausAlumniServer/internal/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Students API called")

	var batchDetail types.BatchDetails

	err := json.NewDecoder(r.Body).Decode(&batchDetail)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := s.db.StudentQuery(batchDetail)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	_, _ = w.Write(jsonResp)
}
