package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// data, err := ioutil.ReadFile("input.txt")
	// if err != nil {
	// 	fmt.Println("File reading error", err)
	// 	return
	// }
	var ints []int
	err := json.Unmarshal([]byte("[2,4,5,7,1,2,3,6]"), &ints)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Before: ")
	for _, i := range ints {
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()
	fmt.Println()

	merge_sort(ints, 0, len(ints))

	fmt.Print("After: ")
	for _, i := range ints {
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()
}

func merge_sort(A []int, p int, r int) []int {
	if p < r {
		q := (p + r) / 2
		merge_sort(A, p, q)
		merge_sort(A, q+1, r)
		merge(A, p, q, r)
	}
	return A
}

func merge(A []int, p int, q int, r int) []int {
	n1 := q - p + 1
	n2 := r - q

	var L []int
	var R []int

	for i := 1; i < n1; i++ {
		L = append(L, A[p+i-1])
	}

	fmt.Print("L: ")
	for _, varr := range L {
		fmt.Print(strconv.Itoa(varr) + " ")
	}
	fmt.Println()

	for j := 0; j < n2; j++ {
		R = append(R, A[q+j])
	}

	fmt.Print("R: ")
	for _, varr := range R {
		fmt.Print(strconv.Itoa(varr) + " ")
	}
	fmt.Println()
	L = append(L, 100)
	R = append(R, 100)
	i := 0
	j := 0

	for k := p; k < r; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}

	}

	return A
}
