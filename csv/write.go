package csv

import (
	"encoding/csv"
	"github.com/DanielFillol/DataJUD_API_CALLER/models"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func Write(fileName string, folderName string, responses []models.ResponseBody) error {
	err := writeLawsuits(fileName, folderName, responses)
	if err != nil {
		return err
	}
	err = writeMovements(fileName+"_movements", folderName, responses)
	if err != nil {
		return err
	}
	return nil
}

// WriteLawsuits writes a CSV file with the given file name and folder name, and the data from the responses.
func writeLawsuits(fileName string, folderName string, responses []models.ResponseBody) error {
	// Create a slice to hold all the rows for the CSV file
	var rows [][]string

	// Add the headers to the slice
	rows = append(rows, generateHeadersLawsuits())

	// Add the data rows to the slice
	for _, response := range responses {
		rows = append(rows, generateRowLawsuits(response)...)
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

	// WriteLawsuits all the rows to the CSV file
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

// generateHeadersLawsuits function returns a slice of strings containing the header values for the CSV file.
func generateHeadersLawsuits() []string {
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
		"Subjects Codes",
		"Subjects",
	}
}

// generateRowLawsuits function takes in a single models.WriteStruct argument and returns a slice of strings containing the values to be written in a row of the CSV file.
func generateRowLawsuits(response models.ResponseBody) [][]string {
	var rows [][]string

	// Create subject codes string
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

	// Create subject string
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

// WriteMovements writes a CSV file with the given file name and folder name, and the data from the responses.
func writeMovements(fileName string, folderName string, responses []models.ResponseBody) error {
	// Create a slice to hold all the rows for the CSV file
	var rows [][]string

	// Add the headers to the slice
	rows = append(rows, generateHeadersMovements())

	// Add the data rows to the slice
	for _, response := range responses {
		rows = append(rows, generateRowMovements(response)...)

	}

	// Create the CSV file
	cf, err := createFile(folderName + "/" + fileName + ".csv")
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

// generateHeadersLawsuits function returns a slice of strings containing the header values for the CSV file.
func generateHeadersMovements() []string {
	return []string{
		"LawsuitNumber",
		"Movement Code",
		"Movement Date",
		"Movement",
		"Movement Complement Code",
		"Movement Complement Name",
		"Movement Complement Description",
		"Movement Complement Value",
	}
}

// generateRowLawsuits function takes in a single models.WriteStruct argument and returns a slice of strings containing the values to be written in a row of the CSV file.
func generateRowMovements(response models.ResponseBody) [][]string {
	var rows [][]string

	for _, lawsuit := range response.Hit.Hits {
		for _, movement := range lawsuit.Source.Movements {
			for _, complement := range movement.Complement {
				var row []string
				row = append(row, lawsuit.Source.LawsuitNumber)
				row = append(row, strconv.Itoa(movement.Code))
				row = append(row, movement.DateTime.String())
				row = append(row, movement.Name)
				row = append(row, strconv.Itoa(complement.Code))
				row = append(row, complement.Name)
				row = append(row, complement.Description)
				row = append(row, strconv.Itoa(complement.Value))
				rows = append(rows, row)
			}
		}
	}

	return rows
}
