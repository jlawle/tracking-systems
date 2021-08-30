package main

import "fmt"
import "gonum.org/v1/gonum/mat"
import "log"

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
	calculateVars()
}

// Calculate a given solution from matrices
func calculateVars() {
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
}

// Function for easy matrix format prints
func printMat(name string, m mat.Matrix) {
	mf := mat.Formatted(m, mat.FormatMATLAB())
	fmt.Printf("Matrix %s: \n %#v \n", name, mf)
}