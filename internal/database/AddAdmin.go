package database

import (
	"ausAlumniServer/internal/types"
	"context"
	"log"
	"time"
)

func (s *service) AddAdmin(hodDetails types.Hod, username string, hashedPassword string) error {
	query := `INSERT INTO admins (department_id, ishod, name, username, hashed_password)
			VALUES (
				(SELECT department_id FROM departments WHERE school_id = (SELECT school_id FROM schools WHERE name = $1) AND name = $2),
				true,
				$3,
				$4,
				$5
			);`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, hodDetails.School, hodDetails.Department, hodDetails.Name, username, hashedPassword)
	if err != nil {
		log.Printf("Error in QueryContext: %v", err)
		return err
	}
	defer rows.Close()

	return nil
}
