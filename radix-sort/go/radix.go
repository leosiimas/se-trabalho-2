package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func radixSortLSD(arr []int) []int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	exp := 1
	for max/exp > 0 {
		count := make([]int, 10)
		for i := 0; i < len(arr); i++ {
			count[(arr[i]/exp)%10]++
		}
		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}
		Output := make([]int, len(arr))
		for i := len(arr) - 1; i >= 0; i-- {
			Output[count[(arr[i]/exp)%10]-1] = arr[i]
			count[(arr[i]/exp)%10]--
		}
		for i := 0; i < len(arr); i++ {
			arr[i] = Output[i]
		}
		exp *= 10
	}
	return arr
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

	fmt.Printf("\nIniciando Execucao do Radix Sort em Go: \n\n")
	result := radixSortLSD(numbers)
	fmt.Printf("%v\n", result)
}
