package database

import (
	"context"
	"log/slog"
	"time"
)

func (s *service) DeleteAlumniQuery(alumniId int) error {
	query := `DELETE FROM public.graduated_students WHERE student_id = $1;`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.db.QueryContext(ctx, query, alumniId)
	if err != nil {
		slog.Info("Error in QueryContext: %v", err)
		return err
	}

	return nil
}
