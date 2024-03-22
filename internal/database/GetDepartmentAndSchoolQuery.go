package database

import (
	"context"
	"log"
	"time"
)

func (s *service) GetDepartmentAndSchoolQuery(deptId string) (map[string]string, error) {
	query := `SELECT s.name AS school_name, d.name AS department_name
				FROM schools s
				JOIN departments d ON s.school_id = d.school_id
				WHERE d.department_id = $1;`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := s.db.QueryContext(ctx, query, deptId)
	if err != nil {
		log.Printf("Error getting department and school")
		return map[string]string{}, err
	}

	var departmentAndSchoolName []map[string]string
	for rows.Next() {
		var departmentName string
		var schoolName string

		err = rows.Scan(&schoolName, &departmentName)
		if err != nil {
			log.Printf("Error getting department and school")
			return map[string]string{}, err
		}
		departmentAndSchool := map[string]string{"schoolName": schoolName, "departmentName": departmentName}
		departmentAndSchoolName = append(departmentAndSchoolName, departmentAndSchool)
	}
	// fmt.Println(departmentNames)
	return departmentAndSchoolName[0], nil
}
