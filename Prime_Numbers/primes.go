package main

import "fmt"

func main() {

	for i := 2; i <= 10000; i++ {
		counter := 0
		for j := i - 1; j > 1; j-- {

			if i%j == 0 {
				counter++
			}
		}
		if counter == 0 {
			fmt.Println(i)
		}
	}
}
