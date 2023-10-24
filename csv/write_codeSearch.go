package csv

import (
	"encoding/csv"
	"github.com/DanielFillol/DataJUD_API_CALLER/models"
	"log"
	"strconv"
)

// WriteCode writes two CSV files with the given file name and folder name, and the data from the responses. One file for the Lawsuits and another for the lawsuit movements
func WriteCode(fileName string, folderName string, responses []models.ResponseBodyNextPage) error {
	err := writeCode(fileName, folderName, responses)
	if err != nil {
		return err
	}
	return nil
}

// writeCode writes a CSV file with the given file name and folder name, and the data from the responses.
func writeCode(fileName string, folderName string, responses []models.ResponseBodyNextPage) error {
	// Create a slice to hold all the rows for the CSV file
	var rows [][]string

	// Add the headers to the slice
	rows = append(rows, generateHeadersCode())

	// Add the data rows to the slice
	for _, response := range responses {
		rows = append(rows, generateRowCode(response)...)
	}

	// Create the CSV file
	cf, err := createFile(folderName + "/" + fileName + "requestsCode.csv")
	if err != nil {
		log.Println(err)
		return err
	}

	// Close the file when the function completes
	defer cf.Close()

	// Create a new CSV writer
	w := csv.NewWriter(cf)

	// WriteLawsuits all the rows to the CSV file
	err = w.WriteAll(rows)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// generateHeadersCode function returns a slice of strings containing the header values for the CSV file.
func generateHeadersCode() []string {
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
		"Index",
		"Type",
		"ID",
		"Score",
		"Class Code",
		"Class",
		"LawsuitNumber",
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
		"Subjects Codes",
		"Subjects",
	}
}

// generateRowCode function takes in a single models.WriteStruct argument and returns a slice of strings containing the values to be written in a row of the CSV file.
func generateRowCode(response models.ResponseBodyNextPage) [][]string {
	var rows [][]string

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
		row = append(row, strconv.Itoa(lawsuit.Source.Class.Code))
		row = append(row, lawsuit.Source.Class.Name)
		row = append(row, lawsuit.Source.LawsuitNumber)
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
		var subjectsCodes string
		var subjects string
		for j, s := range lawsuit.Source.Subjects {
			if j != 0 {
				subjects += " | " + s.Name
				subjectsCodes += " | " + strconv.Itoa(s.Code)
			} else {
				subjects += s.Name
				subjectsCodes += strconv.Itoa(s.Code)
			}
		}
		row = append(row, subjectsCodes)
		row = append(row, subjects)

		rows = append(rows, row)
	}

	return rows
}
