package main

import "fmt"

func main() {
	for i := 1; i <= 1000; i++ {
		counter := 0
		// fmt.Print(i, " ")
		if i == 1 {
			counter = 0
		}
		for j := i; j >= 1; {
			if j%2 == 0 {
				j /= 2
				// fmt.Print(j, " ")
				counter++
			}
			if j == 1 {
				break
			}
			if j%2 != 0 {
				j *= 3
				j++
				counter++
				// fmt.Print(j, " ")
			}
		}
		fmt.Print(counter)
		fmt.Println()
	}
}
