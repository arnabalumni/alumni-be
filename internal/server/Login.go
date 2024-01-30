package server

import (
	"ausAlumniServer/internal/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	// create JWT token
	claims := types.Claims{
		DepartmentId: storedCredentials[0].DepartmentId,
		IsHod:        storedCredentials[0].IsHod,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "ausAlumniServer",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Printf("Error while signing token: %v", err)
		http.Error(w, "Error while signing token", http.StatusInternalServerError)
		return
	}

	resp := map[string]string{"token": signedToken, "status": "success"}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error handling JSON marshal: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
