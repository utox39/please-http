package main

import (
	"io"
	"os"
	"testing"
	"time"
)

// Generates a JSON log file with the expected format and content
func TestGenerateLogWithExpectedFormatAndContent(t *testing.T) {
	// Set up test data
	requestUrl := "https://example.com"
	requestType := "GET"
	results := Results{
		StartTime: time.Now(),
		RespTime:  100,
		Status:    "200 OK",
		StrBody:   "{\"test\": \"test\"}",
	}
	repetitions := 1
	i := 0

	// Call the function
	GenLog(requestUrl, requestType, results, repetitions, i)

	// Read the log file
	logFile, err := os.Open("log.json")
	if err != nil {
		t.Fatalf("Failed to open log file: %v", err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			panic(err)
		}
	}(logFile)

	// Read the log file content
	logContent, err := io.ReadAll(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	// Assert the log file content
	expectedLogContent := `{
  "url": "https://example.com",
  "start-time": "` + results.StartTime.Format("2006-01-02 15:04:05.999999999 -0700 MST") + `",
  "time": "100 ms",
  "request-type": "GET",
  "status-code": "200 OK",
  "response": {"test": "test"}
}`
	if string(logContent) != expectedLogContent {
		t.Errorf("Unexpected log file content.\nExpected: %s\nActual: %s", expectedLogContent, string(logContent))
	}

	if len(logContent) == 0 {
		t.Errorf("The log file is empty.")
	}
}
