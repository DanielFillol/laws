package laws

import (
	"github.com/DanielFillol/laws/data"
	"strings"

	"github.com/DanielFillol/laws/models"
)

var SubjectData []models.LawBySubject

func init() {
	SubjectData = data.CreateLaws()
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
