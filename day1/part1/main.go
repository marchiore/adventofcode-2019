package main

import (
	"advent-of-code-19/pkg/utils"
	"fmt"
	"math"
	"strconv"
)

func main() {

	lines, err := utils.ReadLines("day1/part1/input.txt")

	if err != nil {
		fmt.Printf("Erro ao ler o arquivo")
	}

	total := 0

	for _, line := range lines {

		number, err := strconv.Atoi(line)

		if err != nil {
			fmt.Printf("Erro ao converter a linha para número: %s\n", line)
		}

		total += int(math.Floor(float64(number/3)) - 2)
	}

	fmt.Println(total)
}
