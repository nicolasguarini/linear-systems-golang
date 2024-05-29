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

	filename := "./matrices/small_matrix.mtx"
	a, err := ReadMTX(filename)

	if err != nil {
		fmt.Println("Error reading .mtx file:", err)
		return
	}

	fmt.Println("Matrix A:")
	fmt.Printf("%v\n", mat.Formatted(a))
	rows, _ := a.Dims()

	tmp := make([]float64, rows)
	for i := range tmp {
		tmp[i] = 1
	}

	xEs := mat.NewVecDense(rows, tmp)

	var b mat.VecDense
	b.MulVec(a, xEs)

	fmt.Println(b)

	p, n := ComputePN(a)
	fmt.Println("Matrix P:")
	fmt.Println(mat.Formatted(p))

	fmt.Println("Matrix N:")
	fmt.Println(mat.Formatted(n))

	x0 := mat.NewVecDense(b.Len(), make([]float64, b.Len()))

	var ax mat.VecDense
	ax.MulVec(a, x0)

	axMinusB := SubVectors(&ax, &b)

	fmt.Println("ax:", ax)
	fmt.Println("b:", b)
	fmt.Println("ax - b:", axMinusB)

	fmt.Println("Norm of ax-b:", axMinusB.Norm(2))
	fmt.Println("Norm of b:", b.Norm(2))

	if axMinusB.Norm(2)/b.Norm(2) < tols.At(0, 0) {
		fmt.Println("Finished: stop criterium satisfied!")
	} else {
		fmt.Println("Keep going...")
	}

}
