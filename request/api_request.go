package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/DanielFillol/DataJUD_API_CALLER/models"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// APIRequest makes an API request to the specified URL using the specified HTTP method and authentication header.
// It returns a models.ResponseBody struct containing the API response body and an error (if any).
func APIRequest(url, method string, auth string, request models.ReadCsv) (models.ResponseBody, error) {
	cnj, err := modifyCNJ(request.CNJNumber)
	if err != nil {
		return models.ResponseBody{}, err
	}

	// Create a new BodyRequest struct with the document ID and pagination settings for the initial API call.
	req := models.BodyRequest{
		Query: models.Query{Match: models.Match{CNJNumber: cnj}},
	}

	// Serialize the BodyRequest struct to JSON.
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return models.ResponseBody{}, err
	}

	// Create a new buffer with the JSON-encoded request body.
	reqBody := bytes.NewBuffer(jsonReq)

	// Make the API call and get the response.
	res, err := call(url, method, auth, reqBody)
	if err != nil {
		return models.ResponseBody{}, errors.New(err.Error() + "  " + req.Query.Match.CNJNumber)
	}

	// Read the response body.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.ResponseBody{}, err
	}

	// Unmarshal the response body into a ResponseBody struct.
	var response models.ResponseBody
	err = json.Unmarshal(body, &response)
	if err != nil {
		return models.ResponseBody{}, err
	}

	return models.ResponseBody{
		Took:     response.Took,
		TimedOut: response.TimedOut,
		Shards:   response.Shards,
		Hit:      response.Hit,
	}, nil
}

// call sends an HTTP request to the specified URL using the specified method and request body, with the specified authorization header.
// It returns the HTTP response or an error if the request fails.
func call(url, method string, AUTH string, body io.Reader) (*http.Response, error) {
	// Create an HTTP client with a 10-second timeout.
	client := &http.Client{Timeout: time.Second * 10}

	// Create a new HTTP request with the specified method, URL, and request body.
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// Set the Content-Type and Authorization headers for the request.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", AUTH)

	// Send the request and get the response.
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// If the response status code is not OK, return an error with the status code.
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(strconv.Itoa(response.StatusCode))
	}

	return response, nil
}
