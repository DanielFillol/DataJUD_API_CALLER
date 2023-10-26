package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/DanielFillol/DataJUD_API_CALLER/models"
	"io/ioutil"
	"log"
	"strconv"
)

const SIZE = 100

// APIRequestCode makes an API request to the specified URL using the specified HTTP method and authentication header.
// It returns a models.ResponseBodyNextPage struct containing the API response body and an error (if any).
func APIRequestCode(url, method string, auth string, request models.ReadCsvCode) (models.ResponseBodyNextPage, error) {
	// Create responseReturn new BodyRequestLawsuit struct with the document ID and pagination settings for the initial API call.
	class, err := strconv.Atoi(request.ClassCode)
	if err != nil {
		return models.ResponseBodyNextPage{}, errors.New("failed to parse the class code to int: " + request.ClassCode)
	}

	req := models.BodyRequestCode{
		Size: SIZE,
		Query: models.Query{
			Bool: models.Bool{
				Must: []models.Must{
					{Match: models.Match{ClasseCodigo: class}},
				},
			},
		},
		Sort: []models.Sort{{models.Timestamp{Order: "asc"}}},
	}

	// Serialize the BodyRequestLawsuit struct to JSON.
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return models.ResponseBodyNextPage{}, errors.New(err.Error())
	}

	// Create responseReturn new buffer with the JSON-encoded request body.
	reqBody := bytes.NewBuffer(jsonReq)

	// Make the API call and get the response.
	res, err := call(url, method, auth, reqBody)
	if err != nil {
		return models.ResponseBodyNextPage{}, errors.New(err.Error())
	}

	// Read the response body.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.ResponseBodyNextPage{}, errors.New(err.Error())
	}

	// Unmarshal the response body into responseReturn ResponseBodyLawsuit struct.
	var response models.ResponseBodyNextPage
	err = json.Unmarshal(body, &response)
	if err != nil {
		return models.ResponseBodyNextPage{}, errors.New(err.Error())
	}

	responseReturn := models.ResponseBodyNextPage{
		Took:     response.Took,
		TimedOut: response.TimedOut,
		Shards:   response.Shards,
		Hit:      response.Hit,
	}

	// If the API response has more pages of data, make additional API calls and append the results to the response.
	if len(response.Hit.Hits) == SIZE {
		nextPage := strconv.Itoa(int(response.Hit.Hits[SIZE-1].Sort[0]))
		// Make the API call and get the response.
		lawsuits, err := callNextPage(url, method, auth, req, nextPage)
		if err != nil {
			return models.ResponseBodyNextPage{}, errors.New(err.Error())
		}

		responseReturn.Hit.Hits = append(responseReturn.Hit.Hits, lawsuits...)
		return responseReturn, nil
	}

	return responseReturn, nil
}

// callNextPage returns a slice of models.Lawsuit structs containing the data from all pages of the API response.
func callNextPage(url string, method string, auth string, req models.BodyRequestCode, cursor string) ([]models.Hit2NextPage, error) {
	var lawsuits []models.Hit2NextPage
	c, err := strconv.Atoi(cursor)
	if err != nil {
		return lawsuits, errors.New("could not covert cursor into int: " + cursor)
	}

	for {
		log.Println("success getting next page:" + strconv.Itoa(c))
		// Create a new BodyRequest struct with the document ID and updated pagination settings for the next API call.
		request := models.BodyRequestCodeNextPage{
			Size: SIZE,
			Query: models.Query{
				Bool: models.Bool{
					Must: []models.Must{
						{Match: models.Match{ClasseCodigo: req.Query.Bool.Must[0].Match.ClasseCodigo}},
					},
				},
			},
			Sort:        []models.Sort{{models.Timestamp{Order: "asc"}}},
			SearchAfter: []int64{int64(c)},
		}

		// Serialize the BodyRequest struct to JSON.
		jsonReq, err := json.Marshal(request)
		if err != nil {
			log.Println(err)
			return lawsuits, err
		}

		// Create a new buffer with the JSON-encoded request body.
		reqBody := bytes.NewBuffer(jsonReq)

		// Call the API using the provided url, method, authorization, and request body.
		res, err := call(url, method, auth, reqBody)
		if err != nil {
			log.Println(err)
			return lawsuits, errors.New(err.Error())
		}

		// Read the response body.
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return lawsuits, err
		}

		// Unmarshal the response body into a models.ResponseBody struct.
		var response models.ResponseBodyNextPage
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Println(err)
			return lawsuits, err
		}

		// Append the current response to the lawsuits slice.
		lawsuits = append(lawsuits, response.Hit.Hits...)

		// If the API response indicates there are no more pages, break out of the loop.
		if len(response.Hit.Hits) < SIZE {
			break
		}

		// Update the cursor for the next API call.
		c = int(response.Hit.Hits[SIZE-1].Sort[0])
	}

	return lawsuits, nil
}
