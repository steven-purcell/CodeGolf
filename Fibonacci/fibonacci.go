package main

import "fmt"

func main() {
	j := 0
	for i := 1; i <= 832040; i += j {
		fmt.Println(j)
		fmt.Println(i)
		j += i
	}
	fmt.Println(j)
}
