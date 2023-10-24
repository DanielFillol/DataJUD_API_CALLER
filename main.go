package main

import (
	"github.com/DanielFillol/DataJUD_API_CALLER/csv"
	"github.com/DanielFillol/DataJUD_API_CALLER/request"
	"io"
	"log"
	"os"
	"time"
)

const (
	BASE             = "https://api-publica.datajud.cnj.jus.br/api_publica_"
	AUTH             = "APIKey cDZHYzlZa0JadVREZDJCendQbXY6SkJlTzNjLV9TRENyQk1RdnFKZGRQdw=="
	METHOD           = "POST"
	WORKERS          = 10
	FILENAME         = "response"
	FOLDER           = "data"
	FILEPATH_LAWSUIT = "data/requestsLawsuits.csv"
	FILEPATH_CODE    = "data/requestsCode.csv"
	SEPARATOR        = ','
	HEADER           = true
	IS_LAWSUIT       = true
)

func main() {
	// Setup Log file
	logFile, err := os.Create("output.log.txt")
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	// Create a multi-writer that writes to both the file and os.Stdout (terminal)
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(multiWriter)

	if IS_LAWSUIT {
		// Load data to be requested from CSV file
		requests, err := csv.ReadLawsuit(FILEPATH_LAWSUIT, SEPARATOR, HEADER)
		if err != nil {
			log.Fatal("Error loading requests from CSV: ", err)
		}

		// Make API requests asynchronously
		start := time.Now()
		log.Println("Starting API calls...")

		results, err := request.AsyncAPIRequestLawsuit(requests, WORKERS, BASE, METHOD, AUTH)
		if err != nil {
			log.Println("Error making API requests: ", err)
		}
		log.Println("Finished API calls in ", time.Since(start))

		// WriteLawsuits API response to CSV file
		err = csv.WriteLawsuits(FILENAME, FOLDER, results)
		if err != nil {
			log.Fatal("Error writing API response to CSV: ", err)
		}
	} else {
		// Load data to be requested from CSV file
		requests, err := csv.ReadCode(FILEPATH_CODE, SEPARATOR, HEADER)
		if err != nil {
			log.Fatal("Error loading requests from CSV: ", err)
		}

		// Make API requests asynchronously
		start := time.Now()
		log.Println("Starting API calls...")

		results, err := request.AsyncAPIRequestCode(requests, WORKERS, BASE, METHOD, AUTH)
		if err != nil {
			log.Println("Error making API requests: ", err)
		}
		log.Println("Finished API calls in ", time.Since(start))

		// WriteLawsuits API response to CSV file
		s := time.Now()
		log.Println("Start parsing to .csv ")
		err = csv.WriteCode(FILENAME, FOLDER, results)
		if err != nil {
			log.Fatal("Error writing API response to CSV: ", err)
		}
		log.Println("Finished parsing to .csv in", time.Since(s))
	}

}
