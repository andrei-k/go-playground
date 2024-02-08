package main

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Define the struct representing a data point on the chart
type DataPoint struct {
	xVal int
	yVal float64
}

func main() {
	// Define the data points (unique users)
	data := []DataPoint{
		{xVal: 2017, yVal: 21828},
		{xVal: 2018, yVal: 30340},
		{xVal: 2019, yVal: 36545},
		{xVal: 2020, yVal: 35337},
		{xVal: 2021, yVal: 36744},
		{xVal: 2022, yVal: 41283},
		{xVal: 2023, yVal: 44251},
	}

	/*
		// Define the data points (sessions)
		data := []DataPoint{
			{xVal: 2021, yVal: 1096848},
			{xVal: 2022, yVal: 1560869},
			{xVal: 2023, yVal: 1746767},
		}
	*/

	/*
		// Define the data points (total hours on app)
		data := []DataPoint{
			{xVal: 2021, yVal: 27614},
			{xVal: 2022, yVal: 70828},
			{xVal: 2023, yVal: 81917},
		}
	*/

	// Extract x and y values from the data points
	var xVals []int
	yVals := make([]opts.LineData, 0)

	for _, point := range data {
		xVals = append(xVals, point.xVal)
		yVals = append(yVals, opts.LineData{Value: point.yVal})
	}

	// Create a new line instance
	line := charts.NewLine()

	// Set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Unique users per year",
			// Subtitle: "Unique users per year",
		}))

	// Put data into instance
	line.SetXAxis(xVals)
	line.AddSeries("Unique users", yVals)
	line.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	// Render the output to a file
	f, _ := os.Create("plot.html")
	line.Render(f)

	println("The plot was created in the file plot.html")
}
