package main

import (
	"fmt"
	"shubham_solanki/day1"
)

func main() {

	// Question 1
	fmt.Println("Question 1 : Matrix")
	m1 := day1.Matrix{
		Rows:    2,
		Columns: 2,
		Data: [][]int{
			{1, 2},
			{3, 4},
		},
	}
	m2 := &day1.Matrix{
		Rows:    2,
		Columns: 2,
		Data: [][]int{
			{2, 4},
			{8, 6},
		},
	}
	m1.GetRow()
	m1.GetColumn()
	m1.SetDataAtGivenRowAndColumn(0, 0, 5)
	matrix, _ := m1.AddTwoMatrix(m2)
	matrix.PrintMatrix()

	//Question 2
	fmt.Println("Question 2 : Tree")
	root := &day1.Node{Data: "+"}
	root.InsertLeftNode("a")
	root.InsertRightNode("-")
	root.Left.InsertLeftNode("b")
	root.Left.InsertRightNode("c")

	fmt.Print("PreOrder Traversal: ")
	root.PrintPreOrderTraversal()
	fmt.Println()
	fmt.Print("PostOrder Traversal: ")
	root.PrintPostOrderTraversal()
	fmt.Println()

	//Question 3
	fmt.Println("Question 3 : Employee")
	f := day1.FullTimeEmployee{
		SalaryCalculator: day1.SalaryCalculator{
			PaymentPerDayOrHours: 500,
			WorkedDaysOrHours:    30,
		},
	}
	fmt.Println(f.GetSalary())

	c := day1.ContractEmployee{
		SalaryCalculator: day1.SalaryCalculator{
			PaymentPerDayOrHours: 100,
			WorkedDaysOrHours:    30,
		},
	}
	fmt.Println(c.GetSalary())

	fr := day1.FreelanaceEmployee{
		SalaryCalculator: day1.SalaryCalculator{
			PaymentPerDayOrHours: 100,
			WorkedDaysOrHours:    20,
		},
	}
	fmt.Println(fr.GetSalary())
}
