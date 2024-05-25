// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt and errors packages

import (
	"errors"
	"fmt"
)

// First Func method
func FirstFunc(v interface{}) (interface{}, error) {
	var ok bool // no value so false.

	if !ok {
		return nil, errors.New("false errors for log")
	}
	return v, nil
}

// SecondFunc method
func SecondFunc() {
	defer func() {
		//var da interface{}
		if da := recover(); da != nil {
			fmt.Println("recovering error ", da)
		}
	}()
	var v interface{}
	v = struct{}{}
	var err error
	if _, err = FirstFunc(v); err != nil {
		panic(err)
	}

	fmt.Println("The error never happen")
}

// main method
func main() {
	SecondFunc()
	fmt.Println("The execution ended")
}

/*
    var da interface{}
   purpose of interface.

    When you use panic(), you can provide any value as an argument.This value can be of any type, even custom types you've defined.
    Since you can't predict the type of value that might be passed to panic(), Go uses the interface{} type to accommodate values of any type.




*/
