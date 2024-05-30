package main

import (
	"fmt"
	"time"

	"gonum.org/v1/gonum/mat"
)

func main() {
	tols := mat.NewVecDense(4, []float64{
		0.0001,
		0.000001,
		0.00000001,
		0.0000000001,
	})
	maxIter := 1000000

	filename := "./matrices/spa1.mtx"
	a, err := ReadMTX(filename)

	if err != nil {
		fmt.Println("Error reading .mtx file:", err)
		return
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

	p, _ := ComputePN(a)

	var pInv mat.Dense
	pInv.Inverse(p)

	x := mat.NewVecDense(b.Len(), make([]float64, b.Len()))

	var ax mat.VecDense
	ax.MulVec(a, x)

	var r mat.VecDense
	r.SubVec(&ax, &b)

	k := 0
	for CheckStop(&r, &b, tols.At(3, 0)) {
		k += 1
		newX, newR := UpdateJacobi(x, &pInv, a, &b)

		x = newX
		r = *newR

		if k > maxIter {
			fmt.Println("Non converge!")
			break
		}
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)

	fmt.Println("Soluzione: ")
	fmt.Println(mat.Formatted(x))

	fmt.Println("Numero iterazioni: ", k)
	fmt.Println("Tempo di esecuzione: ", executionTime)
}
