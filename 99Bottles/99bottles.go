package main

import (
	"fmt"
)

func main() {

	for i := 99; i > 2; i-- {
		j := i - 1
		s := fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.", i, i)
		t := fmt.Sprintf("Take one down and pass it around, %d bottles of beer on the wall.\n", j)

		fmt.Println(s)
		fmt.Println(t)
	}
	fmt.Println("2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n")
	fmt.Println("1 bottle of beer on the wall, 1 bottle of beer.\nTake one down and pass it around, no more bottles of beer on the wall.\n")
	fmt.Println("No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.")
}
