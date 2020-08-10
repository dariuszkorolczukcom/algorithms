package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

var w int

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
	w = 0
	merge_sort(ints, 1, len(ints))
	fmt.Printf("Number of swaps: %d \n", w)
	fmt.Print("After: ")
	for _, i := range ints {
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()
}

func merge_sort(A []int, p int, r int) []int {
	fmt.Printf("merge_sort called with: A: %v, p: %d, r: %d ###", A, p, r)
	fmt.Println()
	if p < r {
		q := (p + r) / 2
		merge_sort(A, p, q)
		merge_sort(A, q+1, r)
		merge(A, p-1, q, r)
	}
	return A
}

func merge(A []int, p int, q int, r int) []int {
	fmt.Printf("merge called with: A: %v, p: %d, q %d, r: %d", A, p, q, r)
	fmt.Println()
	//5,5,6
	n1 := q - p + 1 //1
	n2 := r - q     //1

	var L []int
	var R []int

	for i := 1; i < n1; i++ {
		L = append(L, A[p+i-1])
	}
	fmt.Printf("L: %v ", L)

	for j := 0; j < n2; j++ {
		R = append(R, A[q+j])
	}
	fmt.Printf("R: %v\n", R)

	L = append(L, 100)
	R = append(R, 100)
	i := 0
	j := 0

	for k := p; k < r; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
			w++
		} else {
			A[k] = R[j]
			j++
			w++
		}

	}
	//fmt.Printf("merge finished with: A: %v", A)
	//fmt.Println()

	return A
}
