package main

import "fmt"

func batchProcessingPipeline() {
	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, value := range values {
			multipliedValues[i] = value * multiplier
		}
		return multipliedValues
	}

	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, value := range values {
			addedValues[i] = value + additive
		}
		return addedValues
	}

	values := []int{10, 20, 30, 40, 50}
	for _, value := range add(multiply(values, 2), 10) {
		fmt.Println(value)
	}
}

func streamProcessingPipeline() {
	multiply := func(value, multiplier int) int {
		return value * multiplier
	}

	add := func(value, additive int) int {
		return value + additive
	}

	values := []int{10, 20, 30, 40, 50}
	for _, value := range values {
		fmt.Println(add(multiply(value, 2), 10))
	}
}

func main() {
	// batchProcessingPipeline()
	streamProcessingPipeline()
}
