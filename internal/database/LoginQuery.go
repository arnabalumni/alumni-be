package database

import (
	"ausAlumniServer/internal/types"
	"context"
	"log"
	"time"
)

func (s *service) LoginQuery(username string) ([]types.Admins, error) {
	query := `SELECT department_id, ishod, name, hashed_password FROM admins WHERE username = $1 LIMIT 1;`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, username)
	if err != nil {
		log.Printf("Error in QueryContext: %v", err)
		return nil, err
	}
	defer rows.Close()

	var retrievedCredential []types.Admins
	for rows.Next() {
		var admin types.Admins
		err = rows.Scan(&admin.DepartmentId, &admin.IsHod, &admin.Name, &admin.HashedPassword)
		if err != nil {
			log.Printf("Error in Scan during LoginQuery: %v", err) // Log the error, but don't exit the application
			return nil, err
		}
		retrievedCredential = append(retrievedCredential, admin)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error iterating rows in LoginQuery: %v", err)
		return nil, err
	}

	return retrievedCredential, nil
}
