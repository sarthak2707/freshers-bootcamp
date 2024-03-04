package day1

import "errors"

type Matrix struct {
	Rows    int
	Columns int
	Data    [][]int
}

func (m *Matrix) GetRow() int {
	return m.Rows
}

func (m *Matrix) GetColumn() int {
	return m.Columns
}

func (m *Matrix) GetData(row, column int) int {
	return m.Data[row][column]
}

func (m *Matrix) SetDataAtGivenRowAndColumn(row, column, value int) {
	m.Data[row][column] = value
}

func (m *Matrix) AddTwoMatrix(m2 *Matrix) (*Matrix, error) {
	if m.Rows != m2.Rows || m.Columns != m2.Columns {
		return nil, errors.New("matrix are not of same size")
	}
	result := &Matrix{
		Rows:    m.Rows,
		Columns: m.Columns,
		Data:    make([][]int, m.Rows),
	}
	for i := 0; i < m.Rows; i++ {
		result.Data[i] = make([]int, m.Columns)
		for j := 0; j < m.Columns; j++ {
			result.Data[i][j] = m.Data[i][j] + m2.Data[i][j]
		}
	}
	return result, nil
}

func (m *Matrix) PrintMatrix() {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			print(m.Data[i][j], " ")
		}
		println()
	}
}
