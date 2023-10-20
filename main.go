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
	BASE      = "https://api-publica.datajud.cnj.jus.br/api_publica_"
	AUTH      = "APIKey cDZHYzlZa0JadVREZDJCendQbXY6SkJlTzNjLV9TRENyQk1RdnFKZGRQdw=="
	METHOD    = "POST"
	WORKERS   = 1
	FILENAME  = "response"
	FOLDER    = "data"
	FILEPATH  = "data/requests.csv"
	SEPARATOR = ','
	HEADER    = true
)

func main() {
	// Load data to be requested from CSV file
	requests, err := csv.Read(FILEPATH, SEPARATOR, HEADER)
	if err != nil {
		log.Fatal("Error loading requests from CSV: ", err)
	}

	// Setup Log file
	logFile, err := os.Create("output.log.txt")
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	// Create a multi-writer that writes to both the file and os.Stdout (terminal)
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(multiWriter)

	// Make API requests asynchronously
	start := time.Now()
	log.Println("Starting API calls...")

	results, err := request.AsyncAPIRequest(requests, WORKERS, BASE, METHOD, AUTH)
	if err != nil {
		log.Println("Error making API requests: ", err)
	}
	log.Println("Finished API calls in ", time.Since(start))

	// Write API response to CSV file
	err = csv.Write(FILENAME, FOLDER, results)
	if err != nil {
		log.Fatal("Error writing API response to CSV: ", err)
	}
}
