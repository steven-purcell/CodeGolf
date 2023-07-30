package main

import "fmt"

func main() {
	for i := 1800; i <= 2400; i++ {
		if i%400 == 0 || (i%4 == 0 && !(i%100 == 0)) {
			fmt.Println(i)
		}
	}
}
