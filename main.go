package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Калькулятор индекса массы тела")
	userKG, userHeight := getUSerInput()
	imt := calculateIMT(userKG, userHeight)
	isLean := imt < 16
	if isLean {
		fmt.Println("у вас недостаток веса")
	}
	outputResult(imt)

}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Print(result)
}

func calculateIMT(userKG float64, userHeight float64) float64 {
	const imtPower = 2
	imt := userKG / math.Pow(userHeight/100, imtPower)
	return imt
}

func getUSerInput() (float64, float64) {
	var userHeight float64
	var userKG float64
	fmt.Print("Введите свой рост: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKG)

	return userKG, userHeight
}
