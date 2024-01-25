package database

import (
	"ausAlumniServer/internal/types"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (s *service) StudentQuery(batchDetail types.BatchDetails) ([]types.StudentDetails, error) {
	query := `SELECT gs.student_id, gs.name, gs.occupation, gs.current_address, gs.email, gs.linkedin
			FROM graduated_students gs
			JOIN programs pr ON gs.program_id = pr.program_id
			JOIN departments dp ON pr.department_id = dp.department_id
			JOIN schools sc ON dp.school_id = sc.school_id
			WHERE sc.name = $1
			AND dp.name = $2
			AND pr.name = $3
			AND gs.admission_year = $4;
			`

	batchDetailAdmissionYear, err := strconv.ParseInt(batchDetail.AdmissionYear, 10, 64) // base 10, up to 64 bits
	fmt.Println(batchDetailAdmissionYear)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error in Parsing admission_year: %v", err))
		return nil, err
	}
	fmt.Println(query)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println(batchDetail)
	rows, err := s.db.QueryContext(ctx, query, batchDetail.School, batchDetail.Department, batchDetail.Program+" Program", batchDetailAdmissionYear)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error in QueryContext: %v", err))
		return nil, err
	}
	defer rows.Close()
	var results []types.StudentDetails
	for rows.Next() {
		fmt.Println("in rows.Next()")
		var result types.StudentDetails
		err = rows.Scan(&result.ID, &result.Name, &result.Occupation, &result.Address, &result.Email, &result.Linkedin)
		if err != nil {
			log.Fatalf(fmt.Sprintf("Error in Scan in StudentsQuery: %v", err))
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}
