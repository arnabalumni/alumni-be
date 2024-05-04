package database

import (
	"log"
	"strconv"
	"strings"
)

type ProgramData struct {
	Programs map[string][]int
}

type DepartmentData struct {
	Departments map[string]ProgramData
}

type SchoolData map[string]DepartmentData

func (s *service) BuildDepartmentsData() (SchoolData, error) {
	query := `SELECT
                sc.name AS school_name,
                dp.name AS department_name,
                pr.name AS program_name,
                array_agg(DISTINCT gs.admission_year ORDER BY gs.admission_year) AS years
              FROM schools sc
              JOIN departments dp ON sc.school_id = dp.school_id
              JOIN programs pr ON dp.department_id = pr.department_id
              JOIN graduated_students gs ON pr.program_id = gs.program_id
              GROUP BY sc.name, dp.name, pr.name
              ORDER BY sc.name, dp.name, pr.name;`

	rows, err := s.db.Query(query)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return nil, err
	}
	defer rows.Close()

	result := make(SchoolData)

	for rows.Next() {
		var schoolName, deptName, progName string
		var yearsStr string
		if err := rows.Scan(&schoolName, &deptName, &progName, &yearsStr); err != nil {
			log.Printf("Error scanning rows: %v", err)
			return nil, err
		}
		log.Printf("Scanned years string: %v", yearsStr)
		// Convert yearsStr to []int
		yearsStr = strings.Trim(yearsStr, "{}") // Assuming yearsStr looks like "{2001,2002,2003}"
		yearItems := strings.Split(yearsStr, ",")

		years := make([]int, len(yearItems))
		for i, yearStr := range yearItems {
			year, err := strconv.Atoi(strings.TrimSpace(yearStr))
			if err != nil {
				log.Printf("Error converting year from string to int: %v", err)
				return nil, err
			}
			years[i] = year
		}

		if _, ok := result[schoolName]; !ok {
			result[schoolName] = DepartmentData{Departments: make(map[string]ProgramData)}
		}
		if _, ok := result[schoolName].Departments[deptName]; !ok {
			result[schoolName].Departments[deptName] = ProgramData{Programs: make(map[string][]int)}
		}
		result[schoolName].Departments[deptName].Programs[progName] = years
	}

	return result, nil
}
