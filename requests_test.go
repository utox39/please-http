package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequest(t *testing.T) {
	requestUrl := "https://httpbin.org/get"

	results, err := GetRequest(requestUrl)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.NotEmpty(t, results.StrBody)
	assert.Equal(t, http.StatusOK, results.StatusCode)
}

func TestGetRequestURLNotFound(t *testing.T) {
	requestUrl := "https://httpbin.org/notfound"

	results, err := GetRequest(requestUrl)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.Equal(t, http.StatusNotFound, results.StatusCode)
}

func TestGetRequestInvalidURL(t *testing.T) {
	requestUrl := "invalid-url"

	_, err := GetRequest(requestUrl)

	assert.Error(t, err)
}

func TestPostRequest(t *testing.T) {
	requestUrl := "https://httpbin.org/post"
	keysValues := []string{"key1=value1", "key2=value2"}

	results, err := PostRequest(requestUrl, keysValues)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.NotEmpty(t, results.StrBody)
	assert.Equal(t, http.StatusOK, results.StatusCode)
}

func TestPostRequestWithEmptyPayload(t *testing.T) {
	requestUrl := "https://httpbin.org/post"
	keysValues := []string{""}

	results, err := PostRequest(requestUrl, keysValues)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.NotEmpty(t, results.StrBody)
	assert.Equal(t, http.StatusOK, results.StatusCode)
}

func TestPostRequestURLNotFound(t *testing.T) {
	requestUrl := "https://httpbin.org/notfound"
	keysValues := []string{"key1=value1", "key2=value2"}

	results, err := PostRequest(requestUrl, keysValues)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.Equal(t, http.StatusNotFound, results.StatusCode)
}

func TestPostRequestWithInvalidURL(t *testing.T) {
	requestUrl := "invalid-url"
	keysValues := []string{"key=value"}

	results, err := PostRequest(requestUrl, keysValues)

	assert.Error(t, err)
	assert.NotEqual(t, http.StatusOK, results.StatusCode)
}

func TestPutRequestWithJsonPayload(t *testing.T) {
	requestUrl := "https://httpbin.org/put"
	keysValues := []string{"key1=value1", "key2=value2"}

	results, err := PutRequest(requestUrl, keysValues)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.NotEmpty(t, results.StrBody)
	assert.Equal(t, http.StatusOK, results.StatusCode)
}

func TestPutRequestWithEmptyPayload(t *testing.T) {
	requestUrl := "https://httpbin.org/put"
	keysValues := []string{""}

	results, err := PutRequest(requestUrl, keysValues)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.NotEmpty(t, results.StrBody)
	assert.Equal(t, http.StatusOK, results.StatusCode)
}

func TestPutRequestURLNotFound(t *testing.T) {
	requestUrl := "https://httpbin.org/notfound"
	keysValues := []string{"key1=value1", "key2=value2"}

	results, err := PutRequest(requestUrl, keysValues)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.Equal(t, http.StatusNotFound, results.StatusCode)
}

func TestPutRequestWithInvalidURL(t *testing.T) {
	requestUrl := "invalid-url"
	keysValues := []string{"key=value"}

	results, err := PutRequest(requestUrl, keysValues)

	assert.Error(t, err)
	assert.NotEqual(t, http.StatusOK, results.StatusCode)
}

func TestPatchRequestWithJsonPayload(t *testing.T) {
	requestUrl := "https://httpbin.org/patch"
	keysValues := []string{"key1=value1", "key2=value2"}

	results, err := PatchRequest(requestUrl, keysValues)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.NotEmpty(t, results.StrBody)
	assert.Equal(t, http.StatusOK, results.StatusCode)
}

func TestPatchRequestWithEmptyPayload(t *testing.T) {
	requestUrl := "https://httpbin.org/patch"
	keysValues := []string{""}

	results, err := PatchRequest(requestUrl, keysValues)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.NotEmpty(t, results.StrBody)
	assert.Equal(t, http.StatusOK, results.StatusCode)
}

func TestPatchRequestURLNotFound(t *testing.T) {
	requestUrl := "https://httpbin.org/notfound"
	keysValues := []string{"key1=value1", "key2=value2"}

	results, err := PatchRequest(requestUrl, keysValues)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.Equal(t, http.StatusNotFound, results.StatusCode)
}

func TestPatchRequestWithInvalidURL(t *testing.T) {
	requestUrl := "invalid-url"
	keysValues := []string{"key=value"}

	results, err := PatchRequest(requestUrl, keysValues)

	assert.Error(t, err)
	assert.NotEqual(t, http.StatusOK, results.StatusCode)
}

func TestDeleteRequest(t *testing.T) {
	requestUrl := "https://httpbin.org/delete"

	results, err := DeleteRequest(requestUrl)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.NotEmpty(t, results.StrBody)
	assert.Equal(t, http.StatusOK, results.StatusCode)
}

func TestDeleteRequestURLNotFound(t *testing.T) {
	requestUrl := "https://httpbin.org/notfound"

	results, err := DeleteRequest(requestUrl)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.Equal(t, http.StatusNotFound, results.StatusCode)
}

func TestDeleteRequestWithInvalidURL(t *testing.T) {
	requestUrl := "invalid-url"

	_, err := DeleteRequest(requestUrl)

	assert.Error(t, err)
}

func TestHeadRequest(t *testing.T) {
	requestUrl := "https://httpbin.org/"

	results, err := HeadRequest(requestUrl)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.Empty(t, results.StrBody)
	assert.Equal(t, http.StatusOK, results.StatusCode)
}

func TestHeadRequestURLNotFound(t *testing.T) {
	requestUrl := "https://httpbin.org/notfound"

	results, err := HeadRequest(requestUrl)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.Equal(t, http.StatusNotFound, results.StatusCode)
}

func TestHeadRequestWithInvalidURL(t *testing.T) {
	requestUrl := "invalid-url"

	_, err := HeadRequest(requestUrl)

	assert.Error(t, err)
}

func TestOptionsRequest(t *testing.T) {
	requestUrl := "https://httpbin.org/"

	results, err := OptionsRequest(requestUrl)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.Empty(t, results.StrBody)
	assert.Equal(t, http.StatusOK, results.StatusCode)
}

func TestOptionsRequestURLNotFound(t *testing.T) {
	requestUrl := "https://httpbin.org/notfound"

	results, err := OptionsRequest(requestUrl)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.NotNil(t, results.Headers)
	assert.NotNil(t, results.Protocol)
	assert.NotZero(t, results.RespTime)
	assert.Equal(t, http.StatusNotFound, results.StatusCode)
}

func TestOptionsRequestWithInvalidURL(t *testing.T) {
	requestUrl := "invalid-url"

	_, err := OptionsRequest(requestUrl)

	assert.Error(t, err)
}
