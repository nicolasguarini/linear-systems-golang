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

	IterativeMethod("Jacobi", "./matrices/spa1.mtx", tols.AtVec(3), maxIter, ComputePNJacobi, UpdateJacobi)
	fmt.Println()
	IterativeMethod("Gauß-Seidel", "./matrices/spa1.mtx", tols.AtVec(3), maxIter, ComputePNGaußSeidel, UpdateGaußSeidel)
}
