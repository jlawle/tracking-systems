package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
	"log"
)

var A *mat.Dense
var b *mat.Dense

func main() {
	A = mat.NewDense(6, 2, []float64{
		5, 1,
		6, 1,
		7, 1,
		8, 1,
		9, 1,
		8, 1})
	b = mat.NewDense(6, 1, []float64{1, 1, 2, 3, 5, 14})

	printMat("A", A)
	printMat("b", b)

	// FIND X
	mat := calculateVars()
	plotData(mat)
}


func getFileData() error {
	file, err := os.Open("data.txt")
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), " ")
		


	}



}

func plotData(m *mat.Dense) {
	p := plot.New()
	p.Add(plotter.NewGrid())

	p.Title.Text = "Plot 2"
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
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.Shape = draw.CrossGlyph{}

	// Add in 2D fitted line
	lsrl := plotter.NewFunction(func(x float64) float64 { return m.At(0, 0)*x + m.At(1, 0) })
	lsrl.Color = color.RGBA{B: 255, A: 255}

	p.Add(s, lsrl)
	p.Legend.Add("scatter", s)
	p.Legend.Add("lsrl", lsrl)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "plot2.png"); err != nil {
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
