package main

import (
	"fmt"
)

func convertDecimalkeBinary(number int) int {
	biner := 0
	counter := 1
	remainder := 0

	for number != 0 {
		remainder = number % 2
		number = number / 2
		biner += remainder * counter
		counter *= 10

	}
	return biner
}

func main() {
	var decimal int
	fmt.Print("bilangan desimal:")
	fmt.Scanln(&decimal)
	fmt.Printf("bilangan biner %d", convertDecimalkeBinary(decimal))

}
