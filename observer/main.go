// The objective of the Observer is to observe the modifications of an object and notify
// In order to achieve this we have the Observer and the Observable

package main

import (
	"container/list"
	"fmt"
)

type Observable struct {
	subs *list.List
}

// Subscribing to Observable
func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

// Unsubscribing to Observable
func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

// Fire notification
func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

type Person struct {
	Observable
	Name   string
	status string
}

func NewPerson(name, status string) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		Name:       name,
		status:     status,
	}
}

func (p *Person) CatchACold() {
	p.status = "Sick"
	p.Fire(PropertyChanged{"Sick", p.status}) // notify
}

func (p *Person) Cured() {
	p.status = "Cured"
	p.Fire(PropertyChanged{"Cured", p.status}) // notify
}

type DoctorService struct{}

type PropertyChanged struct {
	name  string
	Value interface{}
}

func (d *DoctorService) Notify(data interface{}) {
	fmt.Println("The pacient is", data.(PropertyChanged).Value)
}

func main() {
	p := NewPerson("Guilherme", "Cured")
	ds := &DoctorService{}

	p.Subscribe(ds)
	p.CatchACold()

	p.Unsubscribe(ds)
	p.Cured()
}
