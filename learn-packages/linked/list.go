package linked

import (
	"container/list"
	"fmt"
)

func Linkedz() {
	var x list.List

	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int), "ew?")
	}
}
