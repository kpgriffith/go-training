package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

type ByFirst []person

func (a ByFirst) Len() int {
	return len(a)
}

func (a ByFirst) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByFirst) Less(i, j int) bool {
	return a[i].first < a[j].first
}

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	people = []person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(ByFirst(people))
	fmt.Println(people)

}
