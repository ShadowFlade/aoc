package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input[:]), "\n")
	results := []rune{}
	for i := 0; i < len(lines); i++ {
        numbers := []int {};
		for _, char := range lines[i] {
			number := int(char - '0')
            numbers = append(numbers, number);
            results = append(results, rune(numbers[:1][0]) + rune(numbers[len(numbers)-1:][0]) + '\n')
		}
        numbers = []int{};
        fmt.Print(results)
	}
}
