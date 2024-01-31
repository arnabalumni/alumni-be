package server

import (
	"ausAlumniServer/internal/types"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) GenerateCreds(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Generate Creds API called")
	// body, err := io.ReadAll(r.Body)
	// fmt.Println(string(body))
	var hodDetail types.Hod

	err := json.NewDecoder(r.Body).Decode(&hodDetail)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	cleanName := strings.ReplaceAll(hodDetail.Name, " ", "")
	cleanName = strings.ToLower(cleanName)

	username := fmt.Sprintf("%s%d", cleanName, rnd.Intn(10000))
	password := uuid.New().String()[:10]

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error while hashing the password", http.StatusInternalServerError)
		return
	}

	err = s.db.AddAdmin(hodDetail, username, string(hashedPassword))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	resp := make(map[string]string)
	resp["username"] = username
	resp["password"] = password

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("error handling JSON marshal. Err: %v", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	_, _ = w.Write(jsonResp)
}
