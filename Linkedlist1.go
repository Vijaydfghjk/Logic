package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head   *node // head ={data 48, next }
	length int
}

func (l *linkedlist) prepend(n *node) {
	second := l.head // second is storing the data l.head but l.head is nil
	fmt.Println("secound is:", second)

	l.head = n

	l.head.next = second
	fmt.Println("secound is :", second)

	l.length++
}

func main() {

	mylist := linkedlist{}

	node1 := &node{data: 48}
	//node2 := &node{data: 18}
	mylist.prepend(node1)
	//mylist.prepend(node2)

	fmt.Println(mylist)
	fmt.Println(mylist.head)
	fmt.Println(mylist.head.next)
}
