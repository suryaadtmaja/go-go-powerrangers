package sortz

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func Swapz() {
	kids := []Person{
		{
			"Jake", 34,
		},
		{
			"Alez", 43,
		},
	}

	sort.Sort(ByName(kids))
	fmt.Println(kids, "Kidz")
}
