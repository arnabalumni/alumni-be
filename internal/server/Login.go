package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login API called")

	var temp struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		log.Printf("Invalid request payload: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	storedCredentials, err := s.db.LoginQuery(temp.Username)

	if err != nil {
		log.Printf("Error while fetching credentials: %v", err)
		http.Error(w, "Error while fetching credentials", http.StatusInternalServerError)
		return
	}

	if len(storedCredentials) == 0 {
		http.Error(w, "Invalid Username or Password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedCredentials[0].HashedPassword), []byte(temp.Password))
	// fmt.Println(err) // nil means it is a match
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}
	// flow where user is authenticated

	resp := map[string]string{"status": "success"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error handling JSON marshal: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}

}
