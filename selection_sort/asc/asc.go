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

	for j := 0; j < len(ints)-1; j++ {

		key := ints[j]
		swap := 0

		for i := j + 1; i < len(ints); i++ {
			if ints[i] < key {
				key = ints[i]
				swap = i
			}
		}

		if swap > 0 {
			ints[swap] = ints[j]
			ints[j] = key
		}
	}

	fmt.Print("After: ")
	for _, i := range ints {
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()
}
