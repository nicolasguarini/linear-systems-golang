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

func ComputePN(a *mat.Dense) (*mat.Dense, *mat.Dense) {
	rows, cols := a.Dims()
	p := mat.NewDense(rows, cols, make([]float64, rows*cols))
	n := mat.NewDense(rows, cols, make([]float64, rows*cols))

	for i := 0; i < rows; i++ {
		p.Set(i, i, a.At(i, i))
		for j := 0; j < cols; j++ {
			if i != j {
				n.Set(i, j, -a.At(i, j))
			}
		}
	}

	return p, n
}

func CheckStop(r *mat.VecDense, b *mat.VecDense, tol float64) bool {
	normR := r.Norm(1)
	normB := b.Norm(1)

	ratio := normR / normB
	//fmt.Println("Norm ratio: ", ratio)
	return ratio >= tol
}

func UpdateJacobi(x *mat.VecDense, pInv *mat.Dense, a *mat.Dense, b *mat.VecDense) (*mat.VecDense, *mat.VecDense) {
	var ax mat.VecDense
	ax.MulVec(a, x)

	var r mat.VecDense
	r.SubVec(b, &ax)

	var pr mat.VecDense
	pr.MulVec(pInv, &r)

	x.AddVec(x, &pr)

	return x, &r
}
