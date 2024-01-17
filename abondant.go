package main

import "fmt"

func isAbondant(n int) int {
	sum := 0
	for i := 2; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	if sum <= n {
		return 0 // Le nombre n n'est pas abondant
	} else {
		return 1 // Le nombre n est abondant
	}
}

func main() {
	var i, limit int

	fmt.Print("Entrez une limite inférieure à 100 : ")
	fmt.Scan(&limit)

	for limit >= 100 || limit < 1 {
		fmt.Print("Entrez une limite comprise entre 1 et 100 : ")
		fmt.Scan(&limit)
	}

	fmt.Println()

	for i = 1; i <= limit; i++ {
		if isAbondant(i) == 1 {
			fmt.Println(i)
		}
	}
}
