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

	IterativeMethod("Jacobi", "./matrices/spa2.mtx", tols.AtVec(3), maxIter, ComputePNJacobi, UpdateJacobi)
	fmt.Println()
	IterativeMethod("Gauß-Seidel", "./matrices/spa2.mtx", tols.AtVec(3), maxIter, ComputePNGaußSeidel, UpdateGaußSeidel)
	fmt.Println()
	IterativeMethod("Gradient Descent", "./matrices/spa2.mtx", tols.AtVec(3), maxIter, ComputePNGradientDescent, UpdateGradientDescent)
	fmt.Println()
	IterativeMethod("Coniugated Gradient Descent", "./matrices/spa2.mtx", tols.AtVec(3), maxIter, ComputePNGradientDescent, UpdateConiugatedGradient)
}
