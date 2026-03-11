package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight = 1.7
	var userKG float64 = 78
	var imt = userKG / math.Pow(userHeight, 2)
	fmt.Print(imt)

}
