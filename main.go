package main

import (
	"fmt"
	"math"
)

func main() {
	const imtPower = 2
	userHeight := 1.7 //если тип будет на вывод то короткая запись
	var userKG float64
	userKG = 78 //если мы не используем сразу переменную
	imt := userKG / math.Pow(userHeight, imtPower)
	fmt.Print(imt)

}
