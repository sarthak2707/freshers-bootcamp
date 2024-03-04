package day1

type CalculateSalary interface {
	GetSalary() int
}

type SalaryCalculator struct {
	PaymentPerDayOrHours int
	WorkedDaysOrHours    int
}

type FullTimeEmployee struct {
	SalaryCalculator
}

type ContractEmployee struct {
	SalaryCalculator
}

type FreelanaceEmployee struct {
	SalaryCalculator
}

func (f *FullTimeEmployee) GetSalary() int {
	return f.PaymentPerDayOrHours * f.WorkedDaysOrHours
}

func (c *ContractEmployee) GetSalary() int {
	return c.PaymentPerDayOrHours * c.WorkedDaysOrHours
}

func (f *FreelanaceEmployee) GetSalary() int {
	return f.PaymentPerDayOrHours * f.WorkedDaysOrHours
}
