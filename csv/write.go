package csv

import (
	"encoding/csv"
	"github.com/DanielFillol/DataJUD_API_CALLER/models"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// Write writes a CSV file with the given file name and folder name, and the data from the responses.
func Write(fileName string, folderName string, responses []models.ResponseBody) error {
	// Create a slice to hold all the rows for the CSV file
	var rows [][]string

	// Add the headers to the slice
	rows = append(rows, generateHeaders())

	// Add the data rows to the slice
	for _, response := range responses {
		rows = append(rows, generateRow(response)...)
	}

	// Create the CSV file
	cf, err := createFile(folderName + "/" + fileName + "requests.csv")
	if err != nil {
		log.Println(err)
		return err
	}

	// Close the file when the function completes
	defer cf.Close()

	// Create a new CSV writer
	w := csv.NewWriter(cf)

	// Write all the rows to the CSV file
	err = w.WriteAll(rows)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// createFile function takes in a file path and creates a file in the specified directory. It returns a pointer to the created file and an error if there is any.
func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		log.Println(err)
		return nil, err
	}
	return os.Create(p)
}

// generateHeaders function returns a slice of strings containing the header values for the CSV file.
//
//	I ignore here all the movements, i will to this extraction latter
func generateHeaders() []string {
	return []string{
		"Took",
		"Time Out",
		"Shards Total",
		"Shards Successful",
		"Shards Skipped",
		"Shards failed",
		"Hits Total Value",
		"Hits Total Relation",
		"Hits Max Score",
		// Here its an [] return
		"Index",
		"Type",
		"ID",
		"Score",
		"LawsuitNumber",
		"Class Code",
		"Class",
		"System Code",
		"System",
		"Format Code",
		"Format",
		"Court",
		"Last Update",
		"Degree",
		"Update Document",
		"Distribution Date",
		"Id",
		"Secrecy Level",
		"County Code IBGE",
		"County Code",
		"County",
		// Here its an [] return
		"Subjects Codes",
		"Subjects",
	}
}

// generateRow function takes in a single models.WriteStruct argument and returns a slice of strings containing the values to be written in a row of the CSV file.
func generateRow(response models.ResponseBody) [][]string {
	var rows [][]string

	// Append subjects Codes
	var subjectsCodes string
	for _, hit := range response.Hit.Hits {
		for i, s := range hit.Source.Subjects {
			if i != 0 {
				subjectsCodes += " | " + strconv.Itoa(s.Code)
			} else {
				subjectsCodes += strconv.Itoa(s.Code)
			}
		}
	}

	// Append subjects

	var subjects string
	for _, hit := range response.Hit.Hits {
		for i, s := range hit.Source.Subjects {
			if i != 0 {
				subjects += " | " + s.Name
			} else {
				subjects += s.Name
			}
		}
	}

	for _, lawsuit := range response.Hit.Hits {
		row := []string{
			// All those that repeat
			strconv.Itoa(response.Took),
			strconv.FormatBool(response.TimedOut),
			strconv.Itoa(response.Shards.Total),
			strconv.Itoa(response.Shards.Successful),
			strconv.Itoa(response.Shards.Skipped),
			strconv.Itoa(response.Shards.Failed),
			strconv.Itoa(response.Hit.Total.Value),
			response.Hit.Total.Relation,
			strconv.Itoa(int(response.Hit.MaxScore)),
		}

		// All those are unique
		row = append(row, lawsuit.Index)
		row = append(row, lawsuit.Type)
		row = append(row, lawsuit.Id)
		row = append(row, strconv.Itoa(int(lawsuit.Score)))
		row = append(row, lawsuit.Source.LawsuitNumber)
		row = append(row, strconv.Itoa(lawsuit.Source.Class.Code))
		row = append(row, lawsuit.Source.Class.Name)
		row = append(row, strconv.Itoa(lawsuit.Source.System.Code))
		row = append(row, lawsuit.Source.System.Name)
		row = append(row, strconv.Itoa(lawsuit.Source.Format.Code))
		row = append(row, lawsuit.Source.Format.Name)
		row = append(row, lawsuit.Source.Court)
		row = append(row, lawsuit.Source.DateLastUpdate.String())
		row = append(row, lawsuit.Source.Degree)
		row = append(row, lawsuit.Source.Timestamp.String())
		row = append(row, lawsuit.Source.DistributionDate.String())
		row = append(row, lawsuit.Source.Id)
		row = append(row, strconv.Itoa(lawsuit.Source.SecrecyLevel))
		row = append(row, strconv.Itoa(lawsuit.Source.CourtInstance.CountyCodeIBGE))
		row = append(row, strconv.Itoa(lawsuit.Source.CourtInstance.Code))
		row = append(row, lawsuit.Source.CourtInstance.Name)
		row = append(row, subjectsCodes)
		row = append(row, subjects)
		rows = append(rows, row)
	}

	return rows
}
