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
	strNumbers := map[string]int64{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "zero": 0}

	for i := 0; i < len(lines); i++ {
		numbers := []int64{}

		var prevVal interface{}

		var strTemp string
		for _, char := range lines[i] {
			num, err := strconv.ParseInt(string(char), 10, 0) // base 10, auto-detect bit size
			_, okStr := prevVal.(string)
			_, okInt := prevVal.(int64)
			if err == nil {
				numbers = append(numbers, num)
				prevVal = num
			} else if okStr || prevVal == nil {
				strTemp += string(char) //str char

				if value, ok := strNumbers[strTemp]; ok {
					numbers = append(numbers, value)
					strTemp = ""
				}

				prevVal = string(char)
			} else if okInt {
				strTemp = string(char)
				prevVal = string(char)
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
		fmt.Println(newNumberInt, " new numbers int")

		if err != nil {
			panic("are you that dumb?")
		}

		total += int64(newNumberInt)

		numbers = []int64{}
	}
	fmt.Println(total, " total number")
}
