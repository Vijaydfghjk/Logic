package main

import (
	"fmt"
)

type node1 struct {
	data int
	next *node1
}

type linkedlist1 struct {
	head   *node1 // head ={data 48, next }
	length int
}

func (l *linkedlist1) prepend1(n *node1) {

	second := l.head // now 48 will store second

	l.head = n           // 18
	l.head.next = second //48
	fmt.Println("second is ", second)
	l.length++
}

func (l linkedlist1) printlistData1() {
	toPrint := l.head
	fmt.Println("lengtth is :", l.length)
	for l.length != 0 {

		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--

		fmt.Printf("\n")
	}
}

func main() {

	mylist1 := linkedlist1{}

	nodee1 := &node1{data: 48}
	nodee2 := &node1{data: 18}
	//nodee3 := &node1{data: 16}

	mylist1.prepend1(nodee1)
	mylist1.prepend1(nodee2)
	//mylist1.prepend1(nodee3)
	//fmt.Println(mylist1)
	//fmt.Println(mylist.head.next)
	//fmt.Println(mylist.head)
	mylist1.printlistData1()
}
