package main

import (
	"gonum.org/v1/gonum/mat"
)

func ComputePNJacobi(a *mat.Dense) (*mat.Dense, *mat.Dense) {
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

	p.Inverse(p)
	return p, n
}

func UpdateJacobi(x *mat.VecDense, pInv *mat.Dense, a *mat.Dense, b *mat.VecDense, d *mat.VecDense) (*mat.VecDense, *mat.VecDense, *mat.VecDense) {
	var ax mat.VecDense
	ax.MulVec(a, x)

	var r mat.VecDense
	r.SubVec(b, &ax)

	var pr mat.VecDense
	pr.MulVec(pInv, &r)

	x.AddVec(x, &pr)

	return x, &r, nil
}
