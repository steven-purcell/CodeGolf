package main

import "fmt"

func main() {
	fmt.Println("It's Go time.")

	for i := 1; i <= 10; i++ {
		arr := make([]int, i, i)
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				// fmt.Println(j)
				arr = append(arr, j)
			}
		}
		// printSlice(arr, i)
		fmt.Println(arr)
	}
}

// func printSlice(arr []int, index int) {
// 	fmt.Print(arr[index])
// 	printSlice(arr[0:], index)
// }
