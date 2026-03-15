package main

import (
	"fmt"
	"math"
)

func main() {
	const imtPower = 2
	var userHeight float64
	var userKG float64
	fmt.Print("Калькулятор индекса массы тела\n")
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKG)
	imt := userKG / math.Pow(userHeight, imtPower)
	fmt.Print("Ваш индекс массы тела: ", imt)

}
