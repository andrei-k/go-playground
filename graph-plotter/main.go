package main

import (
	"math/rand"
	"os"
	"sort"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func main() {
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
	var totals []float64
	for year, total := range data {
		years = append(years, year)
		totals = append(totals, total)
	}

	// Sort the years
	sort.Ints(years)

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
	line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	// Where the magic happens
	f, _ := os.Create("bar.html")
	line.Render(f)
}
