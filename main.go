package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// Path to your .mtx file
	filename := "./matrices/small_matrix.mtx"

	// Read the .mtx file
	a, err := readMTX(filename)

	if err != nil {
		fmt.Println("Error reading .mtx file:", err)
		return
	}

	// Print the matrix
	fmt.Println("Matrix A:")
	fmt.Printf("%v\n", mat.Formatted(a))

	p, n := computePN(a)
	fmt.Println("Matrix P:")
	fmt.Println(mat.Formatted(p))

	fmt.Println("Matrix N:")
	fmt.Println(mat.Formatted(n))
}
