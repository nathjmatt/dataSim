package main

import (
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const results_path = "./results/"

// Creates an image with the name of imageName of a plot of the samples passed.
func plotFloat64(samples []float64, imageName string) {

	p := plot.New()

	p.Title.Text = "Wave"
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Magnitude"

	line, err := plotter.NewLine(plotter.XYs{})
	if err != nil {
		panic(err)
	}

	dt := 1 / 1920.0
	for i := 0; i < len(samples); i++ {
		t := float64(i) * dt
		line.XYs = append(line.XYs, plotter.XY{X: t, Y: samples[i]})
	}
	p.Add(line)

	// Save the plot to an image file
	if err := p.Save(10*vg.Centimeter, 10*vg.Centimeter, imageName); err != nil {
		panic(err)
	}

	// Move the image file to the results folder
	if err := os.Rename(imageName, results_path+imageName); err != nil {
		panic(err)
	}

}
