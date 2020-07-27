package main

import (
	"fmt"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	num1 := []int{1, 0, 0, 1, 1, 1}
	num2 := []int{1, 0, 1, 0, 1, 1}
	var sum [7]int

	fmt.Print("num1:  ")
	for _, i := range num1 {
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()
	fmt.Print("num2:  ")
	for _, i := range num2 {
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()

	for j := len(sum) - 1; j > 0; j-- {
		i := j - 1
		if num1[i] != num2[i] {
			if sum[j] == 1 {
				sum[j] = 0
				sum[i] = 1
				continue
			}
			sum[j] = 1
			continue
		}
		if num1[i] == 0 {
			if sum[j] == 1 {
				sum[j] = 1
				continue
			}
			sum[j] = 0
			continue
		}
		if num1[i] == 1 {
			if sum[j] == 1 {
				sum[j] = 1
				sum[i] = 1
				continue
			}
			sum[j] = 0
			sum[j-1] = 1
		}

	}

	fmt.Print("sum: ")
	for _, i := range sum {
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()
}
