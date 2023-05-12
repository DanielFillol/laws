package csv

import (
	"encoding/csv"
	"github.com/DanielFillol/laws/models"
	"log"
	"os"
)

// The Read function reads data from a CSV file located at the specified filePath, with the specified separator.
// It returns a slice of models.LawBySubject structs containing the data from the CSV file, excluding the header.
func Read(filePath string, separator rune, skipHeaderLine bool) ([]models.LawBySubject, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	csvR := csv.NewReader(csvFile)
	csvR.Comma = separator

	csvData, err := csvR.ReadAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var data []models.LawBySubject
	for i, line := range csvData {
		if skipHeaderLine {
			// Skip the header line
			if i != 0 {
				data = append(data, models.LawBySubject{
					Subject:        line[0],
					SubjectID:      line[1],
					LawDescription: line[2],
				})
			}
		} else {
			data = append(data, models.LawBySubject{
				Subject:        line[0],
				SubjectID:      line[1],
				LawDescription: line[2],
			})
		}
	}

	return data, nil
}
