package main

import (
	"fmt"
)

func isArmstrong(n int) int {
	index := 0
	result := 0
	array := make([]int, 25)

	if n <= 0 {
		return -1
	}

	for n != 0 {
		array[index] = n % 10
		result += array[index] * array[index] * array[index]
		n = n / 10
		index = index + 1

		if index >= 25 || result >= 1000 {
			return -1
		}
	}

	return 1
}

func main() {
	var i, limit int

	fmt.Print("Entrez une limite inférieure à 1000 : ")
	fmt.Scan(&limit)

	for limit >= 1000 || limit < 1 {
		fmt.Print("Entrez une limite comprise entre 1 et 1000 : ")
		fmt.Scan(&limit)
	}

	fmt.Println()

	for i = 1; i <= limit; i++ {
		if isArmstrong(i) != -1 {
			fmt.Println(i)
		}
	}
}
