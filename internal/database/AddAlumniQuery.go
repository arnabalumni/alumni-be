package database

import (
	"ausAlumniServer/internal/types"
	"context"
	"log"
	"strconv"
	"time"
)

func (s *service) AddAlumni(alumniDetails types.FullStudentDetails) error {
	AdmissionYear, err := strconv.ParseInt(alumniDetails.AdmissionYear, 10, 64) // base 10, up to 64 bits
	if err != nil {
		log.Printf("Error in Parsing admission_year: %v", err)
		return err
	}
	query := `INSERT INTO graduated_students 
		(name, program_id, admission_year, occupation, current_address, email, linkedin) 
		VALUES 
		($1, 
			(SELECT program_id FROM programs WHERE name = $2 AND department_id = 
				(SELECT department_id FROM departments WHERE name = $3 AND school_id = 
					(SELECT school_id FROM schools WHERE name = $4))),
		$5, 
		$6, 
		$7, 
		$8, 
		$9);`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, alumniDetails.Name, alumniDetails.Program+" Program", alumniDetails.Department, alumniDetails.School, AdmissionYear, alumniDetails.Occupation, alumniDetails.Address, alumniDetails.Email, alumniDetails.Linkedin)
	if err != nil {
		log.Printf("Error in QueryContext: %v", err)
		return err
	}
	defer rows.Close()

	return nil
}
