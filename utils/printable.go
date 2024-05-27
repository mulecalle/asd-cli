package utils

import (
	"fmt"
)

func PrintSliceString(header string, input []string) {
	fmt.Printf("%s\n", header)
	for index := range input {
		fmt.Printf("%s\n", input[index])
	}
}
