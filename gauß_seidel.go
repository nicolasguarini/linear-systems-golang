package main

import (
	"gonum.org/v1/gonum/mat"
)

func ComputePNGaußSeidel(a *mat.Dense) (*mat.Dense, *mat.Dense) {
	rows, cols := a.Dims()
	p := mat.NewDense(rows, cols, make([]float64, rows*cols))
	n := mat.NewDense(rows, cols, make([]float64, rows*cols))

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i >= j {
				p.Set(i, j, a.At(i, j))
			} else {
				n.Set(i, j, -a.At(i, j))
			}
		}
	}

	return p, n
}

func UpdateGaußSeidel(x *mat.VecDense, p *mat.Dense, a *mat.Dense, b *mat.VecDense, d *mat.VecDense) (*mat.VecDense, *mat.VecDense, *mat.VecDense) {
	var ax mat.VecDense
	ax.MulVec(a, x) // Ax = A * x
	var r mat.VecDense
	r.SubVec(b, &ax) // r = b - Ax

	y := ForwardSubstitution(p, &r) // Py = r

	x.AddVec(x, y) // x = x + y

	return x, &r, nil
}
