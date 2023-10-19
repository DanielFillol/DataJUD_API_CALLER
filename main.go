package main

import (
	"github.com/DanielFillol/DataJUD_API_CALLER/csv"
	"github.com/DanielFillol/DataJUD_API_CALLER/request"
	"log"
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
