package main

import (
	"fmt"
	"time"

	"gonum.org/v1/gonum/mat"
)

// Generic function type for computing P and N matrices
type ComputePNFunc func(a *mat.Dense) (*mat.Dense, *mat.Dense)

// Generic function type for updating the iterative solution
type UpdateFunc func(x *mat.VecDense, pInv *mat.Dense, a *mat.Dense, b *mat.VecDense) (*mat.VecDense, *mat.VecDense)

func CheckStop(r *mat.VecDense, b *mat.VecDense, tol float64) bool {
	normR := r.Norm(1)
	normB := b.Norm(1)
	ratio := normR / normB

	return ratio >= tol
}

func IterativeMethod(methodName string, filename string, tol float64, maxIter int, computePN ComputePNFunc, update UpdateFunc) *mat.VecDense {
	fmt.Println(methodName, " - Matrix:", filename, " Tolerance:", tol, " Max Iterations:", maxIter)
	a, err := ReadMTX(filename)

	if err != nil {
		fmt.Println("Error reading .mtx file:", err)
		return nil
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

	var pInv mat.Dense
	pInv.Inverse(p)

	x := mat.NewVecDense(b.Len(), make([]float64, b.Len()))

	var ax mat.VecDense
	ax.MulVec(a, x)

	var r mat.VecDense
	r.SubVec(&ax, &b)

	k := 0
	for CheckStop(&r, &b, tol) {
		k += 1
		newX, newR := update(x, &pInv, a, &b)

		x = newX
		r = *newR

		if k > maxIter {
			fmt.Println("The solution does not converge.")
			break
		}
	}

	if k <= maxIter {
		fmt.Println("The solution converges.")
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	fmt.Println("Number of iterations: ", k)
	fmt.Println("Execution time: ", executionTime)

	return x
}
