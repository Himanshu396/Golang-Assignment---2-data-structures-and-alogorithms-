package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("Go Linked Lists Tutorial")

	mylist := list.New()
	mylist.PushBack(1)
	mylist.PushFront(2)
	mylist.PushFront(6)
	mylist.PushFront(9)

	for element := mylist.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		fmt.Println(element.Value)
	}

}
