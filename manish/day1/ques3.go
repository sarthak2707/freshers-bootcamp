package day1

type Employee interface {
	GetSalary() int
}

type FullTimeEmployee struct {
	DaysWorked int
	DailyRate  int
}

type ContractEmployee struct {
	DaysWorked int
	DailyRate  int
}

type FreelanceEmployee struct {
	HoursWorked int
	HourlyRate  int
}

func (f FullTimeEmployee) GetSalary() int {
	return f.DaysWorked * f.DailyRate
}

func (c ContractEmployee) GetSalary() int {
	return c.DaysWorked * c.DailyRate
}

func (f FreelanceEmployee) GetSalary() int {
	return f.HoursWorked * f.HourlyRate
}
