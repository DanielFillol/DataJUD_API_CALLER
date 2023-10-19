package csv

import (
	"encoding/csv"
	"github.com/DanielFillol/DataJUD_API_CALLER/models"
	"log"
	"os"
)

// The Read function reads data from a CSV file located at the specified filePath, with the specified separator.
// It returns a slice of models.ReadCsv structs containing the data from the CSV file, excluding the header.
func Read(filePath string, separator rune, skipHeaderLine bool) ([]models.ReadCsv, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer csvFile.Close()

	csvR := csv.NewReader(csvFile)
	csvR.Comma = separator

	csvData, err := csvR.ReadAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var data []models.ReadCsv
	for i, line := range csvData {
		if skipHeaderLine {
			// Skip the header line
			if i != 0 {
				document := line[0]
				data = append(data, models.ReadCsv{
					CNJNumber: document,
				})
			}
		} else {
			document := line[0]
			data = append(data, models.ReadCsv{
				CNJNumber: document,
			})
		}
	}

	return data, nil
}
