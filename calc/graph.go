package calc

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

/*
 https://github.com/gonum/plot/wiki/Example-plots
*/
func CreateGraph(f func(float64) float64) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Functions"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// A quadratic function x^2
	quad := plotter.NewFunction(f)
	quad.Color = color.RGBA{B: 255, A: 255}

	// Add the functions and their legend entries.
	p.Add(quad)
	p.Legend.Add("x^2", quad)
	//p.Legend.ThumbnailWidth = 0.5 * vg.Inch
	p.Legend.ThumbnailWidth = 1.0 * vg.Inch

	// Set the axis ranges.  Unlike other data sets,
	// functions don't set the axis ranges automatically
	// since functions don't necessarily have a
	// finite range of x and y values.
	p.X.Min = -5
	p.X.Max = 10
	p.Y.Min = 0
	p.Y.Max = 10

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "functions.png"); err != nil {
		panic(err)
	}
}
