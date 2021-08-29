package main

import "fmt"
import "gonum.org/v1/gonum/mat"

func main() {
	bvals := []float64{1, 1, 2, 3, 5}
	a := []float64{5, 1, 6, 1, 7, 1, 8, 1, 9, 1}
	A := mat.NewDense(5, 2, a)
	b := mat.NewDense(5, 1, bvals)

	af := mat.Formatted(A, mat.Prefix(" "))
	bf := mat.Formatted(b, mat.Prefix(" "))

	fmt.Printf("Matrix A: \n %v \n", af)
	fmt.Printf("Matrix b: \n %v \n", bf)

	// FIND X

	var x mat.Dense
	var c mat.Dense
	x.Mul(A.T(), A)

	var inv mat.Dense
	err := inv.Inverse(x)
	if err != nil {
		log.Fatalf("Error creating inverse: %v", err)
	}

	c.Mul(inv, A.T())

	cf := mat.Formatted(c, mat.Prefix(" "))
	fmt.Printf("Matrix c: \n %v \n", cf)
}
