package main

import "fmt"

type generalFunc func(int) bool

func main() {
	nums := []int{1, 1, 2, 3, 5, 8, 13}
	smaller := filter(nums, smallerThan10)
	fmt.Printf("%v", smaller)
}

func filter(nums []int, condition generalFunc) []int {
	out := []int{}
	for _, num := range nums {
		if condition(num) {
			out = append(out, num)
		}
	}
	return out
}

func smallerThan10(i int) bool {
	return i < 10
}
