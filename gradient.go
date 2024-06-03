package main

import (
	"gonum.org/v1/gonum/mat"
)

func ComputePNGradientDescent(a *mat.Dense) (*mat.Dense, *mat.Dense) {
	return nil, nil
}

func UpdateGradientDescent(x *mat.VecDense, p *mat.Dense, a *mat.Dense, b *mat.VecDense, d *mat.VecDense) (*mat.VecDense, *mat.VecDense, *mat.VecDense) {
	r := ComputeR(a, b, x)

	var y mat.VecDense
	y.MulVec(a, r)

	aT := mat.Dot(r, r)
	bT := mat.Dot(r, &y)

	alpha_k := aT / bT

	x.AddScaledVec(x, alpha_k, r)

	return x, r, nil
}
