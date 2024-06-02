package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func ReadMTX(filename string) (*mat.Dense, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rows, cols, nnz int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "%") {
			continue
		}
		fmt.Sscanf(line, "%d %d %d", &rows, &cols, &nnz)
		break
	}

	a := mat.NewDense(rows, cols, make([]float64, rows*cols))

	for scanner.Scan() {
		var row, col int
		var value float64
		line := scanner.Text()
		fmt.Sscanf(line, "%d %d %f", &row, &col, &value)

		a.Set(row-1, col-1, value)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return a, nil
}

func ForwardSubstitution(L *mat.Dense, b *mat.VecDense) *mat.VecDense {
	rows, cols := L.Dims()
	if rows != cols {
		panic("L must be a square matrix")
	}

	x := mat.NewVecDense(rows, nil)

	for i := 0; i < rows; i++ {
		sum := b.AtVec(i)
		for j := 0; j < i; j++ {
			sum -= L.At(i, j) * x.AtVec(j)
		}
		x.SetVec(i, sum/L.At(i, i))
	}

	return x
}

func ComputeR(a *mat.Dense, b *mat.VecDense, x *mat.VecDense) *mat.VecDense {
	var ax mat.VecDense
	ax.MulVec(a, x)

	var r mat.VecDense
	r.SubVec(b, &ax)

	return &r
}
