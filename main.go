package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Калькулятор индекса массы тела")

	for {
		userKG, userHeight := getUSerInput()
		imt, err := calculateIMT(userKG, userHeight)
		if err != nil {
			// fmt.Println("no param")
			// continue
			panic("no param")

		}
		outputResult(imt)
		isRepeat := checkRepeatUser()

		if !isRepeat {
			break
		}

	}

}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Println(result)
	switch {
	case IMT < 16:
		fmt.Println("у вас сильный дифецит массы тела")
	case IMT < 18.5:
		fmt.Println("у вас дифецит массы тела")
	case IMT < 25:
		fmt.Println("у вас нормальная  масса тела")
	case IMT < 30:
		fmt.Println("у вас избыточный вес")
	default:
		fmt.Println("у вас степень ожирения")
	}
}

func calculateIMT(userKG float64, userHeight float64) (float64, error) {
	if userKG <= 0 || userHeight <= 0 {
		return 0, errors.New("NO_PARAM")
	}
	const imtPower = 2
	imt := userKG / math.Pow(userHeight/100, imtPower)
	return imt, nil
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

func checkRepeatUser() bool {
	var userChoise string
	fmt.Print("Хотите повторить? ")
	fmt.Scan(&userChoise)

	if userChoise == "yes" || userChoise == "Yes" {
		return true
	}
	return false
}
