// The objective is to reduce common data
// HOW:
// Store common data externally
// Specify an index or a pointer into the external data store

package main

import (
	"fmt"
	"strings"
)

// BAD WAY
/*
type User struct {
	FullName string
}

func NewUser(fullName string) *User {
	return &User{FullName: fullName}
}
*/

// BETTER WAY
var allNames []string

type User struct {
	names []uint8
}

func NewUser(fullName string) *User {
	// Verifies if the value exists on allNames, if so return the value, if not store it to allNames
	getOrAdd := func(s string) uint8 {
		for i := range allNames {
			if allNames[i] == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}

	result := User{}
	parts := strings.Split(fullName, " ")
	for _, p := range parts { // For each name do verification/storing in getOrAdd()
		result.names = append(result.names, getOrAdd(p))
	}
	return &result
}

func (u *User) FullName() string {
	var parts []string
	for _, id := range u.names {
		parts = append(parts, allNames[id])
	}
	return strings.Join(parts, " ")
}

func main() {
	/*
		john := NewUser("John Doe")
		jane := NewUser("Jane Doe")
		alsoJane := NewUser("Jane Smith")
		fmt.Println(john.FullName)
		fmt.Println(jane.FullName)
		fmt.Println("Memory taken by users:",
			len([]byte(john.FullName))+
				len([]byte(alsoJane.FullName))+
				len([]byte(jane.FullName)))
	*/

	rui := NewUser("Rui Monteiro")
	filipe := NewUser("Filipe Monteiro")
	fmt.Println(rui.FullName())
	fmt.Println(filipe.FullName())
	totalMem := 0
	for _, a := range allNames {
		totalMem += len([]byte(a))
	}
}
