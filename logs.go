package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	
	"github.com/tidwall/sjson"
)

func GenLog(requestUrl string, requestType string, results Results, repetitions int, i int) int {
	const jsonTemplate = `
{
  "url": "",
  "start-time": "",
  "time": "",
  "request-type": "",
  "status-code": "",
`
	value, _ := sjson.Set(jsonTemplate, "url", requestUrl)
	value, _ = sjson.Set(value, "start-time", results.StartTime)
	value, _ = sjson.Set(value, "time", strconv.FormatInt(results.RespTime, 10)+" ms")
	value, _ = sjson.Set(value, "request-type", requestType)
	value, _ = sjson.Set(value, "status-code", results.Status)

	// Create the log file/s
	var filePath string
	if repetitions == 1 {
		filePath = "log.json"
	} else {
		dirName := "logs/"
		if _, err := os.Stat(dirName); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(dirName, os.ModePerm)
			if err != nil {
				fmt.Printf("please: can't create logs dir: %v\n", err)
				return 0
			}
		}

		filePath = filepath.Join(dirName, "log"+strconv.Itoa(i)+".json")
	}

	logFile, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("please: log file's error: %v\n", err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Fatalf("please: error closing the file %v: %v", logFile, err)
		}
	}(logFile)

	// Write to the log.json file
	var logBody string
	if results.StrBody != "" {
		logBody = value + "  \"response\": " + results.StrBody + "\n}"
	} else {
		logBody = value + "  \"response\": \"\"\n}"
	}
	byteQuantity, err := logFile.WriteString(strings.TrimSpace(logBody))
	if err != nil {
		log.Fatalf("please: error writing the log file: %v\n", err)
	}

	return byteQuantity
}
