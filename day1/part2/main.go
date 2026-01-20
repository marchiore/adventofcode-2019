package main

import (
	"advent-of-code-19/pkg/utils"
	"fmt"
	"math"
	"strconv"
)

func main() {

	lines, err := utils.ReadLines("day1/part2/input.txt")

	if err != nil {
		fmt.Printf("Erro ao ler o arquivo")
	}

	total := 0

	for _, line := range lines {

		number, err := strconv.Atoi(line)

		if err != nil {
			fmt.Printf("Erro ao converter a linha para número: %s\n", line)
		}

		fmt.Printf("total of fuel for %d is %d\n", number, calculateRequiredFuel(number))
		total += calculateRequiredFuel(number)
	}

	fmt.Println(total)
}

func calculateRequiredFuel(mass int) int {

	// fmt.Println(mass)
	necessaryFuel := int(math.Floor(float64(mass/3)) - 2)

	if necessaryFuel < 0 {
		return 0
	}

	return necessaryFuel + calculateRequiredFuel(necessaryFuel)
}
