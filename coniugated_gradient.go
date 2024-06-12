package main

import "gonum.org/v1/gonum/mat"

func UpdateConiugatedGradient(x *mat.VecDense, p *mat.Dense, a *mat.Dense, b *mat.VecDense, d *mat.VecDense) (*mat.VecDense, *mat.VecDense, *mat.VecDense) {
	r := ComputeR(a, b, x)

	var y mat.VecDense
	y.MulVec(a, d)

	var z mat.VecDense
	z.MulVec(a, r)

	alpha_k := mat.Dot(d, r) / mat.Dot(d, &y)

	var alpha_d mat.VecDense
	alpha_d.ScaleVec(alpha_k, d)
	x.AddVec(x, &alpha_d)

	r = ComputeR(a, b, x)

	var w mat.VecDense
	w.MulVec(a, r)

	beta_k := mat.Dot(d, &w) / mat.Dot(d, &y)

	var beta_d mat.VecDense
	beta_d.ScaleVec(beta_k, d)
	d.SubVec(r, &beta_d)

	return x, r, d
}
