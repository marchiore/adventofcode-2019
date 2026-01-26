package main

import (
	"advent-of-code-19/pkg/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, err := utils.ReadLines("day2/part1/input.txt")

	if err != nil {
		fmt.Printf("Erro ao ler o arquivo")
	}

	numbers := getNumbersFromInput(lines[0])

	for i := 0; i < len(numbers); i += 4 {
		if numbers[i] == 99 {
			fmt.Println("halt program")
			break
		}

		var operation = numbers[i]
		var firstNumber = numbers[i+1]
		var secondNumber = numbers[i+2]
		var output = numbers[i+3]
		var result = 0

		switch operation {
		case 1:
			// sum
			result = numbers[firstNumber] + numbers[secondNumber]
		case 2:
			// multiply
			result = numbers[firstNumber] * numbers[secondNumber]
		}

		numbers[output] = result
	}

	fmt.Println(numbers)
}

func getNumbersFromInput(input string) []int {
	strNumbers := strings.Split(input, ",")

	var numbers []int

	for _, value := range strNumbers {
		n, err := strconv.Atoi(value)

		if err != nil {
			fmt.Println("error converting number")
		}

		numbers = append(numbers, n)
	}
	return numbers
}

func process(arr []int) {

}
