package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")
	var total int64
	for i := 0; i < len(lines); i++ {
		numbers := []int64{}

		for _, char := range lines[i] {
			num, err := strconv.ParseInt(string(char), 10, 0) // base 10, auto-detect bit size
			if err == nil {
				numbers = append(numbers, num)
			}
		}

		if len(numbers) <= 0 {
			continue
		}

		var newNumber string
		for _, valNum := range numbers {
			newNumber += strconv.FormatInt(valNum, 10)
		}

		newNumber = newNumber[0:1] + newNumber[len(newNumber)-1:]
		newNumberInt, err := strconv.Atoi(newNumber)

		if err != nil {
			panic("are you that dumb?")
		}

		total += int64(newNumberInt)

		numbers = []int64{}
	}
	fmt.Println(total, " total number")
}
