package main

import (
	"fmt"
	"math"
)

func main() {
	const imtPower = 2
	var userHeight float64
	var userKG float64
	fmt.Println("Калькулятор индекса массы тела")
	fmt.Print("Введите свой рост: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKG)
	imt := userKG / math.Pow(userHeight/100, imtPower)
	// fmt.Printf("Ваш индекс массы тела: %v", imt)
	// fmt.Print("Ваш индекс массы тела:", imt)
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)

}
