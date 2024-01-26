package types

type BatchDetails struct {
	School        string `json:"school"`
	Department    string `json:"department"`
	Program       string `json:"program"`
	AdmissionYear string `json:"admission_year"`
}

type StudentDetails struct {
	ID         int    `json:"id"`
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
