package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println("Let's Go.")
	a, b := get_args()

	common_values := get_factors(a, b)

	gcf := get_greatest_common_factor(a, b, common_values)
	fmt.Println(gcf)
	if gcf != 0 {
		ares := a / gcf
		num := strconv.Itoa(ares)
		bres := b / gcf
		denom := strconv.Itoa(bres)

		fmt.Println(num + "/" + denom)
	}

}

func get_args() (int, int) {
	args := os.Args[1:2]
	// fmt.Println(args[0])
	res := args[0]
	stringSlice := strings.Split(res, "/")
	// fmt.Println(args)
	a := stringSlice[0]
	b := stringSlice[1]

	x, _ := strconv.Atoi(a)
	y, _ := strconv.Atoi(b)

	return x, y
}

func get_factors(a, b int) []int {

	var a_array_length int = a - 2
	var b_array_length int = b - 2
	af := make([]int, a_array_length, a)
	bf := make([]int, b_array_length, b)

	for i := a; i > 1; i-- {
		af = append(af, i)
	}
	for i := b; i > 1; i-- {
		bf = append(bf, i)
	}

	common_values := intersect(af, bf)

	return common_values
}

func get_greatest_common_factor(a int, b int, arr []int) int {
	for i := 0; i <= len(arr); i++ {
		if arr[i] != 0 && a%arr[i] == 0 && b%arr[i] == 0 {
			return arr[i]
		}
	}
	return 0
}

func intersect(slice1, slice2 []int) []int {
	var intersect []int
	for _, element1 := range slice1 {
		for _, element2 := range slice2 {
			if element1 == element2 {
				intersect = append(intersect, element1)
			}
		}
	}
	return intersect //return slice after intersection
}
