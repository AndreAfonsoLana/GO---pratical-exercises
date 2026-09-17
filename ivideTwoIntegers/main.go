package main

import "fmt"

func main() {
	fmt.Println("Welcome ")

	result := divide(10, -1)

	fmt.Println("Result %d", result)
}

func divide(dividend int, divisor int) int {
	total := 0

	positivo := false
	if divisor > 0 {
		positivo = true
	} else {
		positivo = false
	}

	if dividend <= divisor {
		total = 0
		return total
	} else if positivo {
		for loop := 0; loop < dividend; {
			total = total + 1
			loop += divisor
		}
		return total
	}
	fmt.Println("%d %d %s", dividend, divisor, positivo)
	if divisor < 0 && positivo {
		fmt.Println("entrou")
		for loop := 0; loop < dividend; {
			fmt.Println(" loop: %d  dividend: %d", loop, dividend)
			total = total - 1
			loop -= divisor
		}
		return total
	}

	return 0
}
