package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	ContentType  = "application/json"
	GenericError = "please: error: %v\n"
	JsonError    = "please: json error: %v\n"
	RequestError = "please: %v error: %v\n"
)

func GetRequest(requestUrl string) Results {
	var results Results

	results.StartTime = time.Now()
	resp, err := http.Get(requestUrl)
	results.RespTime = time.Since(results.StartTime).Milliseconds()
	if err != nil {
		log.Fatalf(RequestError, GET, err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf(GenericError, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf(GenericError, err)
		}
	}(resp.Body)

	results.StrBody = string(body)
	results.StatusCode = resp.StatusCode
	results.Status = resp.Status
	results.Headers = resp.Header
	results.Protocol = resp.Proto

	return results
}

func PostRequest(requestUrl string, keysValues []string) Results {
	var results Results
	jsonMap := make(map[string]string)

	// TODO: Add syntax check
	var splitValue []string
	for _, value := range keysValues {
		if value != "" {
			splitValue = strings.Split(value, "=")
			jsonMap[splitValue[0]] = splitValue[1]
		}
	}

	jsonBody, err := json.Marshal(jsonMap)
	if err != nil {
		log.Printf(JsonError, err)
	}
	payload := bytes.NewBuffer(jsonBody)

	results.StartTime = time.Now()
	resp, err := http.Post(requestUrl, ContentType, payload)
	results.RespTime = time.Since(results.StartTime).Milliseconds()
	if err != nil {
		log.Fatalf(RequestError, POST, err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf(GenericError, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf(GenericError, err)
		}
	}(resp.Body)

	results.StrBody = string(body)
	results.StatusCode = resp.StatusCode
	results.Status = resp.Status
	results.Headers = resp.Header
	results.Protocol = resp.Proto

	return results
}

func PutRequest(requestUrl string, keysValues []string) Results {
	var results Results
	jsonMap := make(map[string]string)

	var splitValue []string
	for _, value := range keysValues {
		if value != "" {
			splitValue = strings.Split(value, "=")
			jsonMap[splitValue[0]] = splitValue[1]
		}
	}

	jsonBody, err := json.Marshal(jsonMap)
	if err != nil {
		log.Printf(JsonError, err)
	}
	payload := bytes.NewBuffer(jsonBody)

	// Define the PUT request and set any additional headers
	req, err := http.NewRequest(http.MethodPut, requestUrl, payload)
	req.Header.Set("Content-Type", ContentType)
	if err != nil {
		log.Fatalf(RequestError, PUT, err)
	}

	// Instantiate the http Client
	client := &http.Client{}

	results.StartTime = time.Now()
	// Perform the PUT request
	resp, err := client.Do(req)
	results.RespTime = time.Since(results.StartTime).Milliseconds()
	if err != nil {
		log.Fatalf(RequestError, PUT, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf(GenericError, err)
		}
	}(resp.Body)

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf(GenericError, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf(GenericError, err)
		}
	}(resp.Body)

	results.StrBody = string(body)
	results.StatusCode = resp.StatusCode
	results.Status = resp.Status
	results.Headers = resp.Header
	results.Protocol = resp.Proto

	return results
}

func PatchRequest(requestUrl string, keysValues []string) Results {
	var results Results
	jsonMap := make(map[string]string)

	var splitValue []string
	for _, value := range keysValues {
		if value != "" {
			splitValue = strings.Split(value, "=")
			jsonMap[splitValue[0]] = splitValue[1]
		}
	}

	jsonBody, err := json.Marshal(jsonMap)
	if err != nil {
		log.Printf(JsonError, err)
	}

	payload := bytes.NewBuffer(jsonBody)

	// Define the PATCH request and set any additional headers
	req, err := http.NewRequest(http.MethodPatch, requestUrl, payload)
	req.Header.Set("Content-Type", ContentType)
	if err != nil {
		log.Fatalf(RequestError, PATCH, err)
	}

	// Instantiate the http Client
	client := &http.Client{}

	results.StartTime = time.Now()
	// Perform the PATCH request
	resp, err := client.Do(req)
	results.RespTime = time.Since(results.StartTime).Milliseconds()
	if err != nil {
		log.Fatalf(RequestError, PATCH, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf(GenericError, err)
		}
	}(resp.Body)

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf(GenericError, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf(GenericError, err)
		}
	}(resp.Body)

	results.StrBody = string(body)
	results.StatusCode = resp.StatusCode
	results.Status = resp.Status
	results.Headers = resp.Header
	results.Protocol = resp.Proto

	return results
}

func DeleteRequest(requestUrl string) Results {
	var results Results

	// Define the DELETE request and set any additional headers
	req, err := http.NewRequest(http.MethodDelete, requestUrl, nil)
	if err != nil {
		log.Fatalf(RequestError, DELETE, err)
	}

	// Instantiate the http Client
	client := &http.Client{}

	results.StartTime = time.Now()
	// Perform the DELETE request
	resp, err := client.Do(req)
	results.RespTime = time.Since(results.StartTime).Milliseconds()
	if err != nil {
		log.Fatalf(RequestError, DELETE, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf(GenericError, err)
		}
	}(resp.Body)

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf(GenericError, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf(GenericError, err)
		}
	}(resp.Body)

	results.StrBody = string(body)
	results.StatusCode = resp.StatusCode
	results.Status = resp.Status
	results.Headers = resp.Header
	results.Protocol = resp.Proto

	return results
}

func HeadRequest(requestUrl string) Results {
	var results Results

	// Define the HEAD request and set any additional headers
	req, err := http.NewRequest(http.MethodHead, requestUrl, nil)
	if err != nil {
		log.Fatalf(RequestError, HEAD, err)
	}
	// Instantiate the http Client
	client := &http.Client{}

	results.StartTime = time.Now()
	// Perform the HEAD request
	resp, err := client.Do(req)
	results.RespTime = time.Since(results.StartTime).Milliseconds()
	if err != nil {
		log.Fatalf(RequestError, HEAD, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf(GenericError, err)
		}
	}(resp.Body)

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf(GenericError, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf(GenericError, err)
		}
	}(resp.Body)

	results.StrBody = string(body)
	results.StatusCode = resp.StatusCode
	results.Status = resp.Status
	results.Headers = resp.Header
	results.Protocol = resp.Proto

	return results
}

func OptionsRequest(requestUrl string) Results {
	var results Results

	// Define the OPTIONS request and set any additional headers
	req, err := http.NewRequest(http.MethodOptions, requestUrl, nil)
	if err != nil {
		log.Fatalf(RequestError, OPTIONS, err)
	}
	// Instantiate the http Client
	client := &http.Client{}

	results.StartTime = time.Now()
	// Perform the OPTIONS request
	resp, err := client.Do(req)
	results.RespTime = time.Since(results.StartTime).Milliseconds()
	if err != nil {
		log.Fatalf(RequestError, OPTIONS, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf(GenericError, err)
		}
	}(resp.Body)

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf(GenericError, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf(GenericError, err)
		}
	}(resp.Body)

	results.StrBody = string(body)
	results.StatusCode = resp.StatusCode
	results.Status = resp.Status
	results.Headers = resp.Header
	results.Protocol = resp.Proto

	return results
}
