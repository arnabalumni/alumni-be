package database

import (
	"ausAlumniServer/internal/types"
	"context"
	"log"
	"time"
)

func (s *service) UpdateAlumni(alumniDetails types.StudentDetails) error {

	query := `UPDATE graduated_students
              SET name = $2, occupation = $3, current_address = $4, email = $5, linkedin = $6
              WHERE student_id = $1;`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query,
		alumniDetails.StudentId,  // $1 - id of the student to update
		alumniDetails.Name,       // $2 - new name
		alumniDetails.Occupation, // $3 - new occupation
		alumniDetails.Address,    // $4 - new address (assuming this matches the 'current_address' column)
		alumniDetails.Email,      // $5 - new email
		alumniDetails.Linkedin,   // $6 - new linkedin
	)

	if err != nil {
		log.Printf("Error updating student details: %v", err)
		return err
	}

	return nil
}
