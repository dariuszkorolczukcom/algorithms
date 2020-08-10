package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	var ints []int
	// err := json.Unmarshal([]byte("[8,3,6,5,4,7,2,1,10,11]"), &ints)
	err = json.Unmarshal(data, &ints)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Before: ")
	for _, i := range ints {
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()
	fmt.Println()

	ints = bubble_sort(ints)

	fmt.Print("After: ")
	for _, i := range ints {
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()
}

func bubble_sort(A []int) []int {
	k := 0
	for i := 0; i < len(A); i++ {
		for j := len(A) - 1; j > i; j-- {
			if A[j] < A[j-1] {
				A[j] = A[j] + A[j-1]
				A[j-1] = A[j] - A[j-1]
				A[j] = A[j] - A[j-1]
				k++
			}
		}
	}
	fmt.Printf("Number of swaps: %d \n", k)
	return A
}
