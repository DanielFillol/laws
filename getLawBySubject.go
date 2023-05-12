package laws

import (
	"log"
	"strings"

	"github.com/DanielFillol/laws/csv"
	"github.com/DanielFillol/laws/models"
)

const (
	FILEPATH   = "data/laws_by_subject.csv"
	SEPARATOR  = ','
	SKIPHeader = true
)

var SubjectData []models.LawBySubject

func init() {
	// Initialize the lawsSubject variable by reading the CSV file.
	csvLaw, err := csv.Read(FILEPATH, SEPARATOR, SKIPHeader)
	if err != nil {
		log.Fatalf("error reading source file: %v", err)
	}
	SubjectData = csvLaw
}

// GetLawsBySubject returns legal typification's related to a given subject.
func GetLawsBySubject(subject string) []models.Typification {
	// Create an empty slice to store the legal typification's.
	var legalTypification []models.Typification

	// Convert the subject to lowercase for case-insensitive matching.
	s := strings.ToLower(subject)

	// Loop through the lawsSubject variable to find matching laws.
	for _, laws := range SubjectData {
		if strings.Contains(s, strings.ToLower(laws.Subject)) {
			// Check whether the typification already exists in the slice.
			isOk := true
			for _, lt := range legalTypification {
				if lt.TypeOfOccurrence == laws.Subject {
					isOk = false
				}
			}

			// If the typification does not already exist, add it to the slice.
			if isOk {
				legalTypification = append(legalTypification, models.Typification{
					Legislation:      []models.Legislation{{Law: laws.LawDescription}},
					TypeOfOccurrence: laws.Subject,
				})
			}
		}
	}

	// Return the legal typification's
	return legalTypification
}
