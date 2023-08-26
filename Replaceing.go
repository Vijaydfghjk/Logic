package main

import (
	"fmt"
)

func main() {
	var word string = "papa"
	var b, a rune = 'p', 'm'
	fmt.Println("Please input word and letters to convert. ex) papa p m\n> ")
	fmt.Scanf("%s %c %c\n", &word, &b, &a)

	for _, w := range word {
		if w == b {
			w = a
		}
		fmt.Printf("%c", w)
	}

}



// in for loop -> first w will send p that is == b variable mean w - p ==p -b yes it is matching
// w=a instead p  m will come set bcz this condition w = a.