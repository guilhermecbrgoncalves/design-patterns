// The objective is to create a sort of security middleware for structs or interfaces in order to not expose them without
// validation or to aggregate structs and/or interfaces that are not on the same place

package main

import "fmt"

type Driven interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car    Car
	driver *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 16 {
		c.car.Drive()
	} else {
		fmt.Println("Driver too young")
	}
}

/*
func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{Car{}, driver}
}
*/

func main() {
	car := &CarProxy{Car{}, &Driver{20}}
	car.Drive()
}
