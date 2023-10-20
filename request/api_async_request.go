package request

import (
	"github.com/DanielFillol/DataJUD_API_CALLER/models"
	"log"
	"sync"
)

const API = "/_search"

// AsyncAPIRequest makes API requests asynchronously
func AsyncAPIRequest(users []models.ReadCsv, numberOfWorkers int, url string, method string, auth string) ([]models.ResponseBody, error) {
	// Create a channel to signal when the goroutines are done processing inputs
	done := make(chan struct{})
	defer close(done)
	// Create a channel to receive inputs from
	inputCh := StreamInputs(done, users)

	// Create a wait group to wait for the worker goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)

	// Create a channel to receive results from
	resultCh := make(chan models.ResponseBody)

	// Spawn worker goroutines to process inputs
	var errorOnApiRequests error
	for i := 0; i < numberOfWorkers; i++ {
		go func() {
			// Each worker goroutine consumes inputs from the shared input channel
			for input := range inputCh {
				// Make the API request and send the response to the result channel
				tj, err := defineTJ(input.CNJNumber)
				bodyStr, err := APIRequest(url+tj+API, method, auth, input)
				resultCh <- bodyStr
				if err != nil {
					// If there is an error making the API request, print the error
					log.Println("error send request: " + err.Error())
					errorOnApiRequests = err
					//break
				}
				log.Println("success send request: " + "200 " + input.CNJNumber)
			}
			// When the worker goroutine is done processing inputs, signal the wait group
			wg.Done()
		}()
	}

	// Wait for all worker goroutines to finish processing inputs
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Return early on error in any given call on API requests
	if errorOnApiRequests != nil {
		return nil, errorOnApiRequests
	}

	// Collect results from the result channel and return them as a slice
	var results []models.ResponseBody
	for result := range resultCh {
		results = append(results, result)
	}

	return results, nil
}

// StreamInputs sends inputs from a slice to a channel
func StreamInputs(done <-chan struct{}, inputs []models.ReadCsv) <-chan models.ReadCsv {
	// Create a channel to send inputs to
	inputCh := make(chan models.ReadCsv)
	go func() {
		defer close(inputCh)
		for _, input := range inputs {
			select {
			case inputCh <- input:
			case <-done:
				// If the done channel is closed prematurely, finish the loop (closing the input channel)
				break
			}
		}
	}()
	return inputCh
}
