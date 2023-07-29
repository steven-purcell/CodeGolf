package main

import "fmt"

func main() {
	f := "Fizz"
	b := "Buzz"
	fb := "FizzBuzz"

	for i := 1; i <= 100; i += 1 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(fb)
		} else if i%3 == 0 {
			fmt.Println(f)
		} else if i%5 == 0 {
			fmt.Println(b)
		} else {
			fmt.Println(i)
		}

	}
}
