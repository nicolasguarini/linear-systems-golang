package main

import (
	"fmt"
	"os"
)

func main() {
	matrices := []string{
		"./matrices/spa1.mtx",
		"./matrices/spa2.mtx",
		"./matrices/vem1.mtx",
		"./matrices/vem2.mtx",
	}
	tols := []float64{
		10e-4,
		10e-6,
		10e-8,
		10e-10,
	}
	maxIter := 20000

	//IterativeMethod(methods[0], matrices[0], tols[0], maxIter, ComputePNJacobi, UpdateJacobi)

	file, err := os.OpenFile("performances.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for i := 0; i < len(tols); i++ {
		for j := 0; j < len(matrices); j++ {
			_, iterations, execTime, relError := IterativeMethod("Jacobi", matrices[j], tols[i], maxIter, ComputePNJacobi, UpdateJacobi)
			fmt.Fprintf(file, "%s,%s,%e,%d,%f,%e\n", "Jacobi", matrices[j], tols[i], iterations, execTime, relError)
			fmt.Println("-----------------------------------------------------------")

			_, iterations, execTime, relError = IterativeMethod("Gauß-Seidel", matrices[j], tols[i], maxIter, ComputePNGaußSeidel, UpdateGaußSeidel)
			fmt.Fprintf(file, "%s,%s,%e,%d,%f,%e\n", "Gauß-Seidel", matrices[j], tols[i], iterations, execTime, relError)
			fmt.Println("-----------------------------------------------------------")

			_, iterations, execTime, relError = IterativeMethod("Gradient Descent", matrices[j], tols[i], maxIter, ComputePNGradientDescent, UpdateGradientDescent)
			fmt.Fprintf(file, "%s,%s,%e,%d,%f,%e\n", "Gradient Descent", matrices[j], tols[i], iterations, execTime, relError)
			fmt.Println("-----------------------------------------------------------")

			_, iterations, execTime, relError = IterativeMethod("Conjugated Gradient Descent", matrices[j], tols[i], maxIter, ComputePNGradientDescent, UpdateConiugatedGradient)
			fmt.Fprintf(file, "%s,%s,%e,%d,%f,%e\n", "Conjugated Gradient Descent", matrices[j], tols[i], iterations, execTime, relError)
			fmt.Println("-----------------------------------------------------------")
		}
	}
}
