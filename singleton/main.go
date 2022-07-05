// Singleton is a creational design pattern with the objective of ensuring that a class has only one instance,
// while providing a global access point to this instance.

package main

import (
	"fmt"
	"sync"
)

type single struct {
}

func (s *single) call() {
	fmt.Println("There can be only one")
}

var lock = &sync.Mutex{}
var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()         // locks the block of code to other go routines
		defer lock.Unlock() // unlocks the block of code until it fully executes
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
			singleInstance.call()
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func main() {

	for i := 0; i < 30; i++ {
		go getInstance()
	}

	// Here we use fmt.Scanln() to ensure the main function needs to wait for inputs,
	//giving time for the goroutine to run
	fmt.Scanln()
}
