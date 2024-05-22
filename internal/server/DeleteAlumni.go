package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (s *Server) DeleteAlumni(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimSpace(chi.URLParam(r, "alumniID")))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Request body incompatible", http.StatusBadRequest)
		return
	}
	err = s.db.DeleteAlumniQuery(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Database query failed", http.StatusBadRequest)
		return
	}
	resp := make(map[string]string)
	resp["message"] = "Alumni deleted successfully"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("error handling JSON marshal. Err: %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	_, _ = w.Write(jsonResp)
}
