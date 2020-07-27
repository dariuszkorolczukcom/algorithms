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

	for j := 1; j < len(ints); j++ {
		key := ints[j]
		i := j - 1
		for i >= 0 && ints[i] > key {
			ints[i+1] = ints[i]
			i--
			ints[i+1] = key
		}
	}

	fmt.Print("After: ")
	for _, i := range ints {
		fmt.Print(strconv.Itoa(i) + " ")
	}
	fmt.Println()
}
