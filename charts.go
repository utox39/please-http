package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func GenerateLineItems(repetitions int, respTimes []int64) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < repetitions; i++ {
		items = append(items, opts.LineData{Value: respTimes[i]})
	}
	return items
}

func LineShowLabel(repetitions int, requests []string, respTimes []int64) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeMacarons}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Requests response's time stats",
			Subtitle: "",
		}),
	)

	line.SetXAxis(requests).
		AddSeries("Response time (ms)", GenerateLineItems(repetitions, respTimes)).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: true,
			}),
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	return line
}

func GenCharts(repetitions int, respTimes []int64) {
	var requestNumbers []string

	for i := 0; i < repetitions; i++ {
		requestNumbers = append(requestNumbers, "response "+strconv.Itoa(i+1))
	}

	page := components.NewPage()
	page.AddCharts(
		LineShowLabel(repetitions, requestNumbers, respTimes),
	)
	f, err := os.Create("stats.html")
	if err != nil {
		panic(err)
	}
	err = page.Render(io.MultiWriter(f))
	if err != nil {
		fmt.Printf("response: chart generation error: %v\n", err)
	}
}
