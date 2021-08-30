package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
)

var A *mat.Dense
var b *mat.Dense

func main() {
	A = mat.NewDense(5, 2, []float64{
		5, 1,
		6, 1,
		7, 1,
		8, 1,
		9, 1})

	b = mat.NewDense(5, 1, []float64{1, 1, 2, 3, 5})

	printMat("A", A)
	printMat("b", b)

	// FIND X
	mat := calculateVars()
	plotData(mat)
}

func plotData(m *mat.Dense) {
	p := plot.New()
	p.Add(plotter.NewGrid())

	p.Title.Text = "Plot 1"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.X.Min = 4
	p.X.Max = 12
	p.Y.Min = 4
	p.Y.Max = 12

	pts := makePoints()

	// Add in our data points
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatalf("error generating scatter plot")
	}

	// Add in 2D fitted line
	lsrl := plotter.NewFunction(func(x float64) float64 { return m.At(0, 0)*x + m.At(1, 0) })
	lsrl.Color = color.RGBA{B: 255, A: 255}

	p.Add(s, lsrl)
	p.Legend.Add("scatter", s)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		log.Fatalf("error saving plot")
	}
}

func makePoints() plotter.XYs {
	xvec := A.ColView(0)
	yvec := b.ColView(0)
	fmt.Printf("xvec len: %d \n", xvec.Len())
	if xvec.Len() != yvec.Len() {
		log.Fatalf("error vector length mismatch")
	}

	pts := make(plotter.XYs, xvec.Len())
	for i := 0; i < xvec.Len(); i++ {
		pts[i].X = xvec.AtVec(i)
		pts[i].Y = yvec.AtVec(i)

	}

	return pts
}

// Calculate a given solution from matrices
func calculateVars() *mat.Dense {
	var x mat.Dense
	x.Mul(A.T(), A)

	var inv mat.Dense
	err := inv.Inverse(&x)
	if err != nil {
		log.Fatalf("Error creating inverse: %v", err)
	}

	printMat("inv", &inv)

	var mult mat.Dense
	mult.Mul(&inv, A.T())

	printMat("mult", &mult)

	var bm mat.Dense
	bm.Mul(&mult, b)

	printMat("mult", &bm)

	return &bm

}

// Function for easy matrix format prints
func printMat(name string, m mat.Matrix) {
	mf := mat.Formatted(m, mat.FormatMATLAB())
	fmt.Printf("Matrix %s: \n %#v \n", name, mf)
}
