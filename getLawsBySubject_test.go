package laws

import (
	"github.com/DanielFillol/laws/models"
	"reflect"
	"testing"
)

func TestGetLawsBySubject(t *testing.T) {
	// Initialize some test cases.
	testCases := []struct {
		subject                 string
		expectedTypification    []models.Typification
		expectedErr             error
		expectedTypificationLen int
	}{
		// Crimes contra o Patrimônio / Estelionato
		{
			subject: "Crimes contra o Patrimônio / Estelionato",
			expectedTypification: []models.Typification{
				{
					Legislation: []models.Legislation{
						{Law: "Artigo 171 do Decreto Lei nº 2.848 de 07 de Dezembro de 1940"},
					},
					TypeOfOccurrence: "Estelionato"},
			},
			expectedErr:             nil,
			expectedTypificationLen: 1,
		},
		// Crimes contra o Patrimônio / Furto Privilegiado
		{
			subject: "Crimes contra o Patrimônio / Furto Privilegiado",
			expectedTypification: []models.Typification{
				{
					Legislation: []models.Legislation{
						{Law: "Artigo 155 do Decreto Lei nº 2.848 de 07 de Dezembro de 1940"},
					},
					TypeOfOccurrence: "Furto"},
				{
					Legislation: []models.Legislation{
						{Law: "Parágrafo 2 Artigo 155 do Decreto Lei nº 2.848 de 07 de Dezembro de 1940"},
					},
					TypeOfOccurrence: "Furto Privilegiado"},
			},
			expectedErr:             nil,
			expectedTypificationLen: 2,
		},
		// Crimes contra a vida / Homicídio Simples
		{
			subject: "Crimes contra a vida / Homicídio Simples",
			expectedTypification: []models.Typification{
				{
					Legislation: []models.Legislation{
						{Law: "Artigo 121 do Decreto Lei nº 2.848 de 07 de Dezembro de 1940"},
					},
					TypeOfOccurrence: "Homicídio"},
				{
					Legislation: []models.Legislation{
						{Law: "Parágrafo 1 Artigo 205 do Decreto Lei nº 1.001 de 21 de Outubro de 1969"},
					},
					TypeOfOccurrence: "Homicídio simples"},
			},
			expectedErr:             nil,
			expectedTypificationLen: 2,
		},
		// not in the list
		{
			subject:                 "not in the list",
			expectedTypification:    nil,
			expectedErr:             nil,
			expectedTypificationLen: 0,
		},
	}

	for _, tc := range testCases {
		typification := GetLawsBySubject(tc.subject)

		if !reflect.DeepEqual(typification, tc.expectedTypification) {
			t.Errorf("GetLawsBySubject(%q) = %v, expected %v", tc.subject, typification, tc.expectedTypification)
		}

		if len(typification) != tc.expectedTypificationLen {
			t.Errorf("GetLawsBySubject(%q) returned %d typification, expected %d", tc.subject, len(typification), tc.expectedTypificationLen)
		}

		if tc.expectedErr != nil {
			// Check that an error was returned
			t.Errorf("GetLawsBySubject(%q) returned an unexpected error: %v", tc.subject, tc.expectedErr)
		}
	}
}
