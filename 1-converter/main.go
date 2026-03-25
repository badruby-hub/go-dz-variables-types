package main

import "fmt"

func main() {
	const usdEur = 1 * 0.9412
	const usdRub = 1 * 92
	sum := enterSum()
	fmt.Print(sum)
}

func enterSum() float64 {
	var sum float64
	fmt.Print("Введите сумму для обмена: ")
	fmt.Scan(&sum)
	return sum
}

func outputResult(sum float64, usdRub, usdEur float64) float64 {
	return sum
}
