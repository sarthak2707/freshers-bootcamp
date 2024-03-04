package day1Exercises

import "fmt"

type Employee interface {
	SalaryCalculator() int
}

type FullTime struct {
	salary int
}

func (ft *FullTime) SalaryCalculator() int {
	return ft.salary
}

type Contractor struct {
	salary int
}

func (ct *Contractor) SalaryCalculator() int {
	return ct.salary
}

type Freelancer struct {
	salary int
}

func (fl *Freelancer) SalaryCalculator(hours int) int {
	return fl.salary * hours
}

func SalaryCalculator() {
	ft := FullTime{15000}
	ct := Contractor{3000}
	fl := Freelancer{100}

	fmt.Printf("Full-time employees get paid %d per month\n", ft.SalaryCalculator())
	fmt.Printf("The contractor gets paid %d per month\n", ct.SalaryCalculator())
	fmt.Printf("Freelancer gets paid %d (if freelancer works %d hours)\n", fl.SalaryCalculator(20), 20)
}
