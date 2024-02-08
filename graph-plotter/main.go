package main

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type DataPoint struct {
	xVal int
	yVal float64
}

func main() {
	/*
		// Define the data
		data := map[int]float64{
			2017: 21828,
			2018: 30340,
			2019: 35139,
			2020: 33978,
			2021: 36744,
			2022: 41283,
			2023: 44251,
		}

		// Extract x (years) and y (usage total) values from the data map
		var years []int
		totals := make([]opts.LineData, 0)

		for year, total := range data {
			years = append(years, year)
			totals = append(totals, opts.LineData{Value: total})
		}

		// Sort the years
		sort.Ints(years)
	*/

	// Define the data points
	data := []DataPoint{
		{xVal: 2017, yVal: 21828},
		{xVal: 2018, yVal: 30340},
		{xVal: 2019, yVal: 35139},
		{xVal: 2020, yVal: 33978},
		{xVal: 2021, yVal: 36744},
		{xVal: 2022, yVal: 41283},
		{xVal: 2023, yVal: 44251},
	}

	// // Print the data points
	// for _, dp := range data {
	// 	fmt.Printf("Year: %d, Value: %.2f\n", dp.Year, dp.Value)
	// }

	// Extract x (years) and y (usage total) values from the data map
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
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	// Put data into instance
	// line.SetXAxis([]string{"2017", "2018", "2019", "2020", "2021", "2022", "2023"}).
	// 	AddSeries("Category A", generateLineItems()).
	// 	AddSeries("Category B", generateLineItems()).
	// 	SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	// Set the x-axis data
	line.SetXAxis(xVals)

	// Add the usage total data to the line chart
	line.AddSeries("Usage totals", yVals)

	line.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	// Render the output to a file
	f, _ := os.Create("plot.html")
	line.Render(f)
}
