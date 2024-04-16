package server

import (
	"ausAlumniServer/internal/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thedatashed/xlsxreader"
)

func ExtractToken(authHeader string) (string, error) {
	// Split the header to separate the 'Bearer' prefix and the token itself
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("authorization header format must be 'Bearer {token}'")
	}
	return parts[1], nil
}

func ValidateToken(signedToken string) (*types.Claims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &types.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*types.Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func (s *Server) AddAlumniBulk(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Alumni Bulk API called")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File", err)
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	mimeType := handler.Header.Get("Content-Type")
	if mimeType != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		fmt.Println("Unsupported file type:", mimeType)
		http.Error(w, "Unsupported file type", http.StatusUnsupportedMediaType)
		return
	}
	tokenString, err := ExtractToken(r.Header.Get("Authorization"))
	if err != nil {
		fmt.Println("Token extraction error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	claims, err := ValidateToken(tokenString)
	if err != nil {
		fmt.Println("Token validation error:", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file contents:", err)
		http.Error(w, "Error reading file contents", http.StatusInternalServerError)
		return
	}

	// Process the file directly from the request
	xl, err := xlsxreader.NewReader(fileBytes)
	if err != nil {
		fmt.Println("Error creating new xlsx reader:", err)
		http.Error(w, "Error processing the file", http.StatusInternalServerError)
		return
	}

	// Assuming you want to process the first sheet
	sheetName := xl.Sheets[0]
	fmt.Printf("Processing sheet: %s\n", sheetName)

	departmentAndSchoolName, err := s.db.GetDepartmentAndSchoolQuery(claims.DepartmentId)
	if err != nil {
		fmt.Println("Error retrieving department and school", err)
		http.Error(w, "Error processing the file", http.StatusInternalServerError)
		return
	}
	// Iterate on the rows of data
	for row := range xl.ReadRows(sheetName) {
		if !claims.IsHod || (row.Cells[5].Value == departmentAndSchoolName["schoolName"] && row.Cells[6].Value == departmentAndSchoolName["departmentName"]) {
			var fullStudentDetails types.FullStudentDetails
			fullStudentDetails.Name = row.Cells[0].Value
			fullStudentDetails.Occupation = row.Cells[1].Value
			fullStudentDetails.Address = row.Cells[2].Value
			fullStudentDetails.Email = row.Cells[3].Value
			fullStudentDetails.Linkedin = row.Cells[4].Value
			fullStudentDetails.School = row.Cells[5].Value
			fullStudentDetails.Department = row.Cells[6].Value
			fullStudentDetails.Program = row.Cells[7].Value
			fullStudentDetails.AdmissionYear = row.Cells[8].Value
			s.db.AddAlumni(fullStudentDetails)
		}
	}
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	resp := make(map[string]string)
	resp["message"] = "Alumni successfully added"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error marshalling response JSON:", err)
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResp)
}
