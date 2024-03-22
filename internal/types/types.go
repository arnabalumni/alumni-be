package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type BatchDetails struct {
	School        string `json:"school"`
	Department    string `json:"department"`
	Program       string `json:"program"`
	AdmissionYear string `json:"admissionYear"`
}

type DeptData struct {
	DeptId string `json:"id"`
}

type StudentDetails struct {
	StudentId  int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Address    string `json:"address"`
	Email      string `json:"email"`
	Linkedin   string `json:"linkedin"`
}

type Credentials struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashedPassword"`
}

type Admins struct {
	DepartmentId   string `json:"departmentId"`
	IsHod          bool   `json:"isHod"`
	Name           string `json:"name"`
	HashedPassword string `json:"hashedPassword"`
}

type Claims struct {
	DepartmentId string `json:"departmentId"`
	IsHod        bool   `json:"isHod"`
	jwt.RegisteredClaims
}

type Hod struct {
	School     string `json:"schoolName"`
	Department string `json:"departmentName"`
	Name       string `json:"name"`
}

type FullStudentDetails struct {
	StudentDetails
	BatchDetails
}
