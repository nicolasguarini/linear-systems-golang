package main

import "gonum.org/v1/gonum/mat"

func UpdateConiugatedGradient(x *mat.VecDense, p *mat.Dense, a *mat.Dense, b *mat.VecDense, d *mat.VecDense) (*mat.VecDense, *mat.VecDense, *mat.VecDense) {
	r := ComputeR(a, b, x)

	// y^(k) = A * d^(k)
	var y mat.VecDense
	y.MulVec(a, d)

	// z^(k) = A * r^(k)
	var z mat.VecDense
	z.MulVec(a, r)

	// α_k = (d^(k) * r^(k)) / (d^(k) * y^(k))
	alpha_k := mat.Dot(d, r) / mat.Dot(d, &y)

	// x^(k+1) = x^(k) + α_k * d^(k)
	var alpha_d mat.VecDense
	alpha_d.ScaleVec(alpha_k, d)
	x.AddVec(x, &alpha_d)

	// r^(k+1) = b - A * x^(k+1)
	r = ComputeR(a, b, x)

	// w^(k+1) = A * r^(k+1)
	var w mat.VecDense
	w.MulVec(a, r)

	// β_k = (d^(k) * w^(k+1)) / (d^(k) * y^(k))
	beta_k := mat.Dot(d, &w) / mat.Dot(d, &y)

	// d^(k+1) = r^(k+1) - β_k * d^(k)
	var beta_d mat.VecDense
	beta_d.ScaleVec(beta_k, d)
	d.SubVec(r, &beta_d)

	return x, r, d
}
