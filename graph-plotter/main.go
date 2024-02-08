package main

import (
	"fmt"
	"log"
	"sort"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

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

	// Extract x (years) and y (totals) values from the data map
	var years []int
	var totals []float64
	for year, total := range data {
		years = append(years, year)
		totals = append(totals, total)
	}

	// Sort the years
	sort.Ints(years)

	// Create a new plot
	p := plot.New()

	// Create a line plotter
	line, err := plotter.NewLine(plotter.XYs{})
	if err != nil {
		log.Fatalf("Could not create line plotter: %v", err)
	}

	// Add the data to the line plotter
	for i, year := range years {
		line.XYs = append(line.XYs, struct{ X, Y float64 }{float64(year), totals[i]})
	}

	// Add the line plotter to the plot
	p.Add(line)

	// Set the title and labels
	p.Title.Text = "Totals over years"
	p.X.Label.Text = "Year"
	p.Y.Label.Text = "Totals"

	// Save the plot to a file
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "plot.png"); err != nil {
		log.Fatalf("Could not save plot: %v", err)
	}
	fmt.Println("Plot saved to plot.png")
}
