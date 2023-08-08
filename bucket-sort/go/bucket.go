package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func insertionSort(array []int) {
	for i := 0; i < len(array); i++ {
		temp := array[i]
		j := i - 1
		for ; j >= 0 && array[j] > temp; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = temp
	}
}

func bucketSort(array []int, bucketSize int) []int {
	var max, min int
	for _, n := range array {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	nBuckets := int(max-min)/bucketSize + 1
	buckets := make([][]int, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]int, 0)
	}

	for _, n := range array {
		idx := int(n-min) / bucketSize
		buckets[idx] = append(buckets[idx], n)
	}

	sorted := make([]int, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			insertionSort(bucket)
			sorted = append(sorted, bucket...)
		}
	}

	return sorted
}

func main() {
	numbersFile, err := os.ReadFile("./number-generator/numbers.txt")
	if err != nil {
		log.Fatal(err)
	}

	var numbers []int
	numbersString := strings.Split(string(numbersFile), ",")

	for _, number := range numbersString {
		n, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, n)
	}

	for _, arg := range os.Args[1:] {
		if n, err := strconv.Atoi(arg); err == nil {
			numbers = append(numbers, n)
		}
	}

	fmt.Printf("\nIniciando Execucao do Bucket Sort em Go: \n\n")
	numbers = bucketSort(numbers, 5)
	fmt.Printf("%v\n", numbers)
}
