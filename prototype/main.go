// The Prototype pattern instructs the copying objects themselves to make copies.
// It introduces a common interface for all objects that support cloning.
// This allows objects to be copied without being bound to their specific classes.

package main

import "fmt"

type prototype interface {
	show(string)
	clone() prototype
}

// Directory type for node
type directory struct {
	name     string
	parents  []directory
	children []prototype
}

// Implement show function
func (f *directory) show(i string) {
	fmt.Println(i + f.name)
	for _, v := range f.children {
		v.show(i + i)
	}
}

// Implement clone function
func (f *directory) clone() prototype {
	cloneFolder := &directory{name: f.name + "_clone"}
	var tempChildren []prototype
	for _, i := range f.children {
		c := i.clone()
		tempChildren = append(tempChildren, c)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

type subdirectory struct {
	name string
}

// Implement show function
func (f *subdirectory) show(i string) {
	fmt.Println(i + f.name)
}

// Implement clone function
func (f *subdirectory) clone() prototype {
	return &subdirectory{name: f.name + "_clone"}
}

func main() {
	// Create new categories
	s1 := &subdirectory{"subdirectory 1"}
	s2 := &subdirectory{"subdirectory 2"}
	s3 := &subdirectory{"subdirectory 3"}

	d1 := &directory{
		children: []prototype{s1},
		name:     "Directory 1",
	}

	d2 := &directory{
		children: []prototype{d1, s2, s3},
		name:     "Directory 2",
	}

	fmt.Println("\nOpen directory 2")
	d2.show("  ")

	cloneFolder := d2.clone()
	fmt.Println("\nClone and open")
	cloneFolder.show("  ")

}
