# Datajud API Caller
This API is a Go project that automates the retrieval of legal data from a remote API. It's designed to work with legal cases and uses CNJ (Conselho Nacional de Justiça) numbers for data retrieval.
You can find all API documentation on their [wiki page](https://datajud-wiki.cnj.jus.br)

## Project Overview
The project's main goal is to make API requests to a legal database, retrieve data for specific CNJ numbers, and save the API responses in CSV files. It is organized into several components:
1. API Requests: The project includes functions for making API requests using HTTP POST requests to the specified API endpoint. It handles authentication and constructs the request body.
2. Data Models: Data models (structs) are provided for representing API request bodies (BodyRequest) and API responses (ResponseBody). These models are used for marshaling and unmarshaling JSON data.
3. CSV Handling: The project offers functionality for reading data from CSV files, making it easy to provide input data for the API requests. It also includes functions for writing API responses to new CSV files, making it simple to save the results in a structured format.
4. Main Application: The main.go file serves as the entry point of the application. It orchestrates the entire process, including reading data from CSV files, making API requests, and writing the responses back to CSV files. 

## Project Components
Here's a brief overview of the key components and code files in the project:
- csv/read.go: Handles reading data from CSV files.
- csv/write.go: Manages writing API responses to CSV files.
- data/requests.csv: Has a few CNJ examples to search on the API
- models/models_csv.go: Defines a data model for representing CNJ numbers.
- models/models_request.go: Contains data models for API request bodie.
- models/models_response.go: Defines data models for API response bodie.
- api_define_tj.go: Defines the "TJ" (Tribunal de Justiça) related to CNJ numbers.
- api_modifyCNJ.go: Modifies CNJ numbers by decomposing and reassembling them.
- api_request.go: Contains functions for making API requests and handling HTTP communication.
- api_async_request.go: Implements asynchronous API requests using goroutines, distributing work among multiple workers.
- main.go: The main entry point of the application, where all components come together to perform the workflow.

## Usage
To use the project, follow these steps:
- Load data into a CSV file. Ensure that the CSV file follows the expected format for input data. Only one column with the CNJ numbers.
- Set the API endpoint, authentication header, and other configuration constants in the main.go file as needed. The API KEY is public and can be found [here](https://datajud-wiki.cnj.jus.br/api-publica/acesso)
- Run the project using go run main.go. The program will read the input data from the CSV file, make asynchronous API requests, and save the API responses in CSV files.

## Getting Started
To get started with this project, you'll need to have Go installed on your system. You can download and install Go from the [official website](https://golang.org/dl/)

Clone the project repository:
```bash
git clone github.com/DanielFillol/DataJUD_API_CALLER

``` 
Navigate to the project directory:
```bash
cd DataJUD_API_CALLER
``` 
You could also just open the folder using your system software

Install project dependencies:
```go
go get -u
```

Run the project:
```go
go run main.go
```

The project will start making API requests and saving the responses in CSV files.

## Contributing
Contributions to this project are welcome. Feel free to submit bug reports, feature requests, or create pull requests to improve the project.
