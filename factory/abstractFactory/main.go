package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func NewEmployeeFactory(position string,
	annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func main() {
	bossFactory := NewEmployeeFactory("CEO", 100000)
	boss := bossFactory.Create("Guilherme")
	fmt.Println(boss)

	slaveFactory := NewEmployeeFactory("Cleaner", 1)
	slave := slaveFactory.Create("Jambé")
	fmt.Println(slave)
}
