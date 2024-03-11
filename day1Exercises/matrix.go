package day1Exercises

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	rows    int
	columns int
	matrix  [][]int
}

func (m *Matrix) getRowCount() int {
	return m.rows
}

func (m *Matrix) getColumnCount() int {
	return m.columns
}

func (m *Matrix) setCellValue(i, j, val int) {
	m.matrix[i][j] = val
}

func (m *Matrix) addMatrix(m2 Matrix) Matrix {
	var result Matrix
	rows := m.getRowCount()
	columns := m.getColumnCount()
	result.rows = rows
	result.columns = columns
	result.matrix = make([][]int, rows)
	for i := 0; i < rows; i++ {
		result.matrix[i] = make([]int, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			result.matrix[i][j] = m.matrix[i][j] + m2.matrix[i][j]
		}
	}
	return result
}

func (m *Matrix) multiplyMatrix(m2 Matrix) Matrix {
	var result Matrix
	rows := m.getRowCount()
	columns := m2.getColumnCount()
	innerDim := m.getColumnCount()
	result.rows = rows
	result.columns = columns
	result.matrix = make([][]int, rows)
	for i := 0; i < rows; i++ {
		result.matrix[i] = make([]int, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			val := 0
			for k := 0; k < innerDim; k++ {
				val += m.matrix[i][k] * m2.matrix[k][j]
			}
			result.matrix[i][j] = val
		}
	}
	return result
}

func (m *Matrix) printAsJson() {
	val, err := json.MarshalIndent(m.matrix, "", " ")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(val))
	}
}

func MatrixExec() {
	var matrix Matrix

	matrix.rows = 3
	matrix.columns = 4
	matrix.matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 7}, {9, 9, 11, 12}}

	fmt.Println("row count: ", matrix.getRowCount())
	fmt.Println("column count: ", matrix.getColumnCount())
	fmt.Println("the original matrix: ")
	matrix.printAsJson()

	matrix.setCellValue(1, 3, 8)
	matrix.setCellValue(2, 1, 10)
	fmt.Println("the matrix after cell updates: ")
	matrix.printAsJson()

	adder := Matrix{3, 4, [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}}
	additionResult := matrix.addMatrix(adder)
	fmt.Println("addition result")
	additionResult.printAsJson()

	multiplier := Matrix{4, 3, [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}}
	multiplicationResult := matrix.multiplyMatrix(multiplier)
	fmt.Println("multiplication result")
	multiplicationResult.printAsJson()
}
