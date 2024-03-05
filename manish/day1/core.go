package day1

import "fmt"

func Question1() {
	m := Matrix{
		Rows:    2,
		Columns: 2,
		Elements: [][]int{
			{1, 2},
			{3, 4},
		},
	}

	fmt.Println("No. of rows of matrix : ", m.GetRows())
	fmt.Println("No. of columns of matrix : ", m.GetColumns())
	fmt.Println("Element at Index (1,1) : ", m.GetElements(1, 1))

	matrix, err := m.AddMatrix(Matrix{
		Rows:    2,
		Columns: 2,
		Elements: [][]int{
			{1, 2},
			{3, 4},
		},
	})

	if err != nil {
		fmt.Println("matrix after addition: ", matrix)
	}

	fmt.Println("matrix : ", m.Elements)
	m.PrintMatrixAsJson()
}

func Question2() {
	root := Node{Val: "+", Left: nil, Right: nil}
	root.InsertLeft("a")
	root.InsertRight("-")
	root.Right.InsertLeft("b")
	root.Right.InsertRight("c")

	fmt.Println("PreOrder Traversal : ")
	root.PrintPreOrderTraversal()
	fmt.Println("\nPostOrder Traversal : ")
	root.PrintPostOrderTraversal()
}

func Question3() {
	fullTimeEmployee := FullTimeEmployee{DaysWorked: 30, DailyRate: 500}
	contractEmployee := ContractEmployee{DaysWorked: 30, DailyRate: 100}
	freelanceEmployee := FreelanceEmployee{HoursWorked: 20, HourlyRate: 100}

	fmt.Println("Full Time Employee Salary : ", fullTimeEmployee.GetSalary())
	fmt.Println("Contract Employee Salary : ", contractEmployee.GetSalary())
	fmt.Println("Freelance Employee Salary : ", freelanceEmployee.GetSalary())
}
