package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func intArrayToBytes(arr []int) []byte {
	size := len(arr) * int(unsafe.Sizeof(arr[0]))

	byteSlice := make([]byte, size)

	arrPtr := (*reflect.SliceHeader)(unsafe.Pointer(&arr))

	copy(byteSlice, *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: arrPtr.Data,
		Len:  size,
		Cap:  size,
	})))

	return byteSlice
}

func main() {
	rand.Seed(time.Now().UnixNano())

	totalNumbers := 90000000
	fmt.Printf("\nGerando %d numeros\n\n", totalNumbers)
	var numbers []string

	for i := 0; i <= totalNumbers; i++ {
		number := strconv.Itoa(rand.Intn(100000))
		numbers = append(numbers, number)
	}

	numbersFile, err := os.OpenFile("./number-generator/numbers.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	numbersFile.WriteString(strings.Join(numbers, ","))

	numbersFile.Close()
}
