package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func max(values []int) int {
	//assign the first element equal to min
	max := values[0] //assign the first element equal to max
	for _, number := range values {
		if number > max {
			max = number
		}
	}
	return max
}

func coutingSort(array []int) []int {
	size := len(array)
	bigger := max(array) + 1
	outputPosition := 0

	var output []int = make([]int, size)
	var count []int = make([]int, bigger)

	for i := 0; i < size; i++ {
		count[array[i]] += 1
	}

	for i := 0; i < len(count); i++ {
		if count[i] != 0 {
			for count[i] > 0 {
				output[outputPosition] = i
				outputPosition += 1
				count[i] -= 1
			}
		}

	}

	return output
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

	fmt.Printf("\nIniciando Execucao do Counting Sort em Go: \n\n")
	numbers = coutingSort(numbers)
	/* fmt.Printf("%v\n", numbers) */

}
