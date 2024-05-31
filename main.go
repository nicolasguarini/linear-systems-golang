package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	tols := mat.NewVecDense(4, []float64{
		0.0001,
		0.000001,
		0.00000001,
		0.0000000001,
	})
	maxIter := 20000

	x := IterativeMethod("Jacobi", "./matrices/spa1.mtx", tols.At(0, 0), maxIter, ComputePNJacobi, UpdateJacobi)
	fmt.Println("Soluton length:", x.Len())
}
