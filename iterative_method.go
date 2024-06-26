package main

import (
	"fmt"
	"time"

	"gonum.org/v1/gonum/mat"
)

// Generic function type for computing P and N matrices
type ComputePNFunc func(a *mat.Dense) (*mat.Dense, *mat.Dense)

// Generic function type for updating the iterative solution
type UpdateFunc func(x *mat.VecDense, p *mat.Dense, a *mat.Dense, b *mat.VecDense, d *mat.VecDense) (*mat.VecDense, *mat.VecDense, *mat.VecDense)

func CheckStop(r *mat.VecDense, b *mat.VecDense, tol float64) bool {
	normR := r.Norm(2)
	normB := b.Norm(2)
	ratio := normR / normB

	return ratio >= tol
}

func IterativeMethod(methodName string, filename string, tol float64, maxIter int, computePN ComputePNFunc, update UpdateFunc) (*mat.VecDense, int, float64, float64) {
	fmt.Println(methodName, "- Matrix:", filename, " Tolerance:", tol, " Max Iterations:", maxIter)
	a, err := ReadMTX(filename)

	if err != nil {
		fmt.Println("Error reading .mtx file:", err)
		return nil, 0, 0, 0
	}

	startTime := time.Now()

	rows, _ := a.Dims()

	tmp := make([]float64, rows)
	for i := range tmp {
		tmp[i] = 1
	}

	xEs := mat.NewVecDense(rows, tmp)

	var b mat.VecDense
	b.MulVec(a, xEs)

	p, _ := computePN(a)

	x := mat.NewVecDense(b.Len(), make([]float64, b.Len()))

	var ax mat.VecDense
	ax.MulVec(a, x)

	r := ComputeR(a, &b, x)

	d := r // Direction vector for coniugate gradient method
	k := 0
	for CheckStop(r, &b, tol) {
		k += 1
		x, r, d = update(x, p, a, &b, d)

		if k > maxIter {
			fmt.Println("The solution does not converge.")
			break
		}
	}

	if k <= maxIter {
		fmt.Println("The solution converges.")
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime).Seconds()

	var s mat.VecDense
	s.SubVec(x, xEs)

	relativeError := s.Norm(2) / xEs.Norm(2)

	fmt.Println("Number of iterations: ", k)
	fmt.Println("Execution time: ", executionTime)
	fmt.Println("Relative Error: ", relativeError)

	return x, k, executionTime, relativeError
}
