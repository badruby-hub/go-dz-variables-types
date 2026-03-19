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
	outputResult(imt)

}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Print(result)
}
