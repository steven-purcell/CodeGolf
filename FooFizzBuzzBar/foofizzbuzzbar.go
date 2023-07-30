package main

import "fmt"

func main() {
	for i := 1; i <= 1000; i++ {
		// fizz := "Fizz"
		// buzz := "Buzz"
		// foo := "Foo"
		// bar := "Bar"
		array := []string{"Fizz", "Buzz", "Foo", "Bar"}
		arr := []int{2, 3, 5, 7}

		for i := 1; i <= 1000; i++ {
			for j := 0; j < len(arr); j++ {
				if i%arr[j] == 0 {
					fmt.Print(array[j])
				}
			}
		}
	}
}
