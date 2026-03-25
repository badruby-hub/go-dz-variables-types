package main

import "fmt"

func main() {
	const usdt = 1
	const rub = 92
	sum := enterSum()
	fmt.Print(sum)
}

func enterSum() float64 {
	var sum float64
	fmt.Print("Введите сумму для обмена: ")
	fmt.Scan(&sum)
	return sum
}

func outputResult(sum float64, fromRate, toRate float64) float64 {
	return sum
}
