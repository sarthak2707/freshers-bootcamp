package day1

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Rows     int
	Columns  int
	Elements [][]int
}

func (m Matrix) GetRows() int {
	return m.Rows
}

func (m Matrix) GetColumns() int {
	return m.Columns
}

func (m Matrix) GetElements(row, col int) int {
	return m.Elements[row][col]
}

func (m Matrix) AddMatrix(b Matrix) (Matrix, error) {
	if (m.Rows != b.Rows) || (m.Columns != b.Columns) {
		return Matrix{}, nil
	}

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			m.Elements[i][j] += b.Elements[i][j]
		}
	}

	return m, nil
}

func (m Matrix) PrintMatrixAsJson() {
	fmt.Println(json.Marshal(m.Elements))
}
