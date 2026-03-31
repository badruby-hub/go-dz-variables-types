package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Калькулятор индекса массы тела")
	userKG, userHeight := getUSerInput()
	imt := calculateIMT(userKG, userHeight)
	outputResult(imt)
	if imt < 16 {
		fmt.Println("у вас сильный дифецит массы тела")
	} else if imt < 18.5 {
		fmt.Println("у вас дифецит массы тела")
	} else if imt < 25 {
		fmt.Println("у вас нормальная  масса тела")
	} else if imt < 30 {
		fmt.Println("у вас избыточный вес")
	} else {
		fmt.Println("у вас степень ожирения")
	}

}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Println(result)
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
