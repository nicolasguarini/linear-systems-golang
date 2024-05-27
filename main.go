package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// Path to your .mtx file
	filename := "./matrices/small_matrix.mtx"

	// Read the .mtx file
	matrix, err := readMTX(filename)

	if err != nil {
		fmt.Println("Error reading .mtx file:", err)
		return
	}

	// Print the matrix
	fmt.Println("Matrix:")
	fmt.Printf("%v\n", mat.Formatted(matrix))
}
