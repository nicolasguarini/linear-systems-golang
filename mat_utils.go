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
