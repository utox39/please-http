package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
)

type Results struct {
	Headers    http.Header
	Protocol   string
	RespTime   int64
	StrBody    string
	StartTime  time.Time
	StatusCode int
	Status     string
}

const (
	GET     string = "GET"
	POST    string = "POST"
	PUT     string = "PUT"
	DELETE  string = "DELETE"
	PATCH   string = "PATCH"
	HEAD    string = "HEAD"
	OPTIONS string = "OPTIONS"

	GenericError = "please: error: "
)

func PrintResults(results Results) {
	pterm.Println("\n- Start time: " + pterm.LightBlue(results.StartTime))
	pterm.Println("- Protocol: " + pterm.Blue(results.Protocol))

	outputStringStatus := "- Status: "
	if results.StatusCode >= 100 && results.StatusCode <= 199 {
		pterm.Println(outputStringStatus + pterm.Yellow(results.Status))
	} else if results.StatusCode >= 200 && results.StatusCode <= 299 {
		pterm.Println(outputStringStatus + pterm.Green(results.Status))
	} else if results.StatusCode >= 300 && results.StatusCode <= 399 {
		pterm.Println(outputStringStatus + pterm.Magenta(results.Status))
	} else if results.StatusCode >= 400 && results.StatusCode <= 499 {
		pterm.Println(outputStringStatus + pterm.Red(results.Status))
	} else if results.StatusCode >= 500 && results.StatusCode <= 599 {
		pterm.Println(outputStringStatus + pterm.Yellow(results.Status))
	}

	pterm.Println("- Time: " + pterm.Green(results.RespTime) + pterm.Green(" ms"))
	pterm.Println("\n- Headers: ")

	for name, values := range results.Headers {
		for _, value := range values {
			pterm.Println("  " + pterm.White(name) + ": " + pterm.LightBlue(value))
		}
	}

	if results.StrBody != "" {
		fmt.Printf("\n- Response:\n%v\n", results.StrBody)
	}
}

func Request(requestType string, requestUrl string, createLog bool, genChart bool, repetitions int, keysValues []string) {
	var respTimes []int64
	var results Results
	var err error

	logFileSuccessfully := "- Log file generated successfully."

	for i := 1; i <= repetitions; i++ {
		switch requestType {
		case GET:
			results, err = GetRequest(requestUrl)
		case POST:
			results, err = PostRequest(requestUrl, keysValues)
		case PUT:
			results, err = PutRequest(requestUrl, keysValues)
		case PATCH:
			results, err = PatchRequest(requestUrl, keysValues)
		case DELETE:
			results, err = DeleteRequest(requestUrl)
		case HEAD:
			results, err = HeadRequest(requestUrl)
		case OPTIONS:
			results, err = OptionsRequest(requestUrl)
		}

		if err != nil {
			var fatalErr PleaseError
			fatalErr.Err = err
			fatalErr.ExitCode = 1
			FatalError(fatalErr)
		}

		respTimes = append(respTimes, results.RespTime)
		PrintResults(results)

		if createLog {
			byteQuantity := GenLog(requestUrl, requestType, results, repetitions, i)
			if byteQuantity > 0 {
				pterm.Println(logFileSuccessfully + pterm.Green(results.Status))
			}
		}
	}

	if genChart && repetitions >= 2 {
		GenCharts(repetitions, respTimes)
		pterm.Println("- Chart generated successfully." + pterm.Green(results.Status))
	} else if genChart && repetitions < 2 {
		fmt.Println("\nplease: chart generation error: there must be at least 2 repetitions.")
	}
}

func main() {
	var createLog bool
	var genChart bool
	var repetitions int

	app := &cli.App{
		Name:  "please",
		Usage: "Http client",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "log",
				Aliases:     []string{"l"},
				Value:       false,
				Usage:       "create a log file of the http request response",
				Destination: &createLog,
			},
			&cli.BoolFlag{
				Name:        "gen-chart",
				Aliases:     []string{"c"},
				Value:       false,
				Usage:       "generate a response time chart and must be called with the --repeat flag (--repeat=n, n>= 2)",
				Destination: &genChart,
			},
			&cli.IntFlag{
				Name:        "repeat",
				Aliases:     []string{"r"},
				Value:       1,
				Usage:       "repeat a request n times",
				Destination: &repetitions,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "get",
				Usage: "Make a GET request.\tE.g: please get https://httpbin.org/get",
				Action: func(cCtx *cli.Context) error {
					requestUrl := cCtx.Args().Get(0)
					Request(GET, requestUrl, createLog, genChart, repetitions, nil)
					return nil
				},
			},
			{
				Name:  "post",
				Usage: "Make a POST request.\tE.g: please post https://httpbin.org/post foo=bar",
				Action: func(cCtx *cli.Context) error {
					requestUrl := cCtx.Args().Get(0)

					if cCtx.Args().Len() < 2 {
						var fatalErr PleaseError
						fatalErr.Err = fewArgsErrMsg
						fatalErr.ExitCode = 1
						FatalError(fatalErr)
					}

					var keysValues []string
					for i := 1; i <= cCtx.Args().Len(); i++ {
						keysValues = append(keysValues, cCtx.Args().Get(i))
					}

					Request(POST, requestUrl, createLog, genChart, repetitions, keysValues)
					return nil
				},
			},
			{
				Name:  "put",
				Usage: "Make a PUT request.\tE.g: please put https://httpbin.org/put foo=bar",
				Action: func(cCtx *cli.Context) error {
					requestUrl := cCtx.Args().Get(0)

					if cCtx.Args().Len() < 2 {
						var fatalErr PleaseError
						fatalErr.Err = fewArgsErrMsg
						fatalErr.ExitCode = 1
						FatalError(fatalErr)
					}

					var keysValues []string
					for i := 1; i <= cCtx.Args().Len(); i++ {
						keysValues = append(keysValues, cCtx.Args().Get(i))
					}

					Request(PUT, requestUrl, createLog, genChart, repetitions, keysValues)
					return nil
				},
			},
			{
				Name:  "patch",
				Usage: "Make a PATCH request.\tE.g: please patch https://httpbin.org/patch foo=bar",
				Action: func(cCtx *cli.Context) error {
					requestUrl := cCtx.Args().Get(0)

					if cCtx.Args().Len() < 2 {
						var fatalErr PleaseError
						fatalErr.Err = fewArgsErrMsg
						fatalErr.ExitCode = 1
						FatalError(fatalErr)
					}

					var keysValues []string
					for i := 1; i <= cCtx.Args().Len(); i++ {
						keysValues = append(keysValues, cCtx.Args().Get(i))
					}

					Request(PATCH, requestUrl, createLog, genChart, repetitions, keysValues)
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "Make a DELETE request.\tE.g: please delete https://httpbin.org/delete",
				Action: func(cCtx *cli.Context) error {
					requestUrl := cCtx.Args().Get(0)
					Request(DELETE, requestUrl, createLog, genChart, repetitions, nil)
					return nil
				},
			},
			{
				Name:  "head",
				Usage: "Make a HEAD request.\tE.g: please head https://httpbin.org/",
				Action: func(cCtx *cli.Context) error {
					requestUrl := cCtx.Args().Get(0)
					Request(HEAD, requestUrl, createLog, genChart, repetitions, nil)
					return nil
				},
			},
			{
				Name:  "options",
				Usage: "Make a OPTIONS request.\tE.g: please options https://httpbin.org/",
				Action: func(cCtx *cli.Context) error {
					requestUrl := cCtx.Args().Get(0)
					Request(OPTIONS, requestUrl, createLog, genChart, repetitions, nil)
					return nil
				},
			},
		},
		Action:  nil,
		Version: "0.3",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
