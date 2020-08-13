package main

import (
	"encoding/json"
	"fmt"
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
	// data, err := ioutil.ReadFile("input.txt")
	// if err != nil {
	// 	fmt.Println("File reading error", err)
	// 	return
	// }
	var ints []int
	err := json.Unmarshal([]byte("[13,-3,-25,20,-3,-16,-23,18,20,-7,12,-5,-22,15,-4,7]"), &ints)
	// err = json.Unmarshal(data, &ints)
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
	low, high, sum := findMaximumSubarray(ints, 0, len(ints)-1)
	fmt.Printf("sum: %d, low: %d, high: %d ###", sum, low, high)
	fmt.Println()
}

func findMaximumSubarray(A []int, low int, high int) (int, int, int) {
	// fmt.Printf("findMaximumSubarray called with: low: %d, high: %d", low, high)
	// fmt.Println()
	if high == low {
		return low, high, A[low]
	}
	mid := (low + high) / 2
	leftLow, leftHigh, leftSum := findMaximumSubarray(A, low, mid)
	rightLow, rightHigh, rightSum := findMaximumSubarray(A, mid+1, high)
	crossLow, crossHigh, crossSum := findMaxCrossingSubarray(A, low, mid, high)
	if leftSum >= rightSum && leftSum >= crossSum {
		fmt.Printf("leftLow: %d, leftHigh %d, leftSum: %d", leftLow, leftHigh, leftSum)
		fmt.Println()
		return leftLow, leftHigh, leftSum
	}
	if rightSum >= leftSum && rightSum >= crossSum {
		fmt.Printf("rightLow: %d, rightHigh %d, rightSum: %d", rightLow, rightHigh, rightSum)
		fmt.Println()
		return rightLow, rightHigh, rightSum
	}
	fmt.Printf("crossLow: %d, crossHigh %d, crossSum: %d", crossLow, crossHigh, crossSum)
	fmt.Println()
	return crossLow, crossHigh, crossSum
}

func findMaxCrossingSubarray(A []int, low int, mid int, high int) (maxLeft int, maxRight int, leftRightSum int) {
	//fmt.Printf("findMaxCrossingSubarray called with: low: %d, mid %d, high: %d", low, mid, high)
	//fmt.Println()

	leftSum := -100000
	sum := 0

	for i := mid; i >= low; i-- {
		sum += A[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}

	rightSum := -100000
	sum = 0

	for j := mid + 1; j <= high; j++ {
		sum += A[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}
	leftRightSum = leftSum + rightSum
	//fmt.Printf("leftRightSum is: %d, maxLeft: %d, maxRight: %d", leftRightSum, maxLeft, maxRight)
	//fmt.Println()
	return
}
