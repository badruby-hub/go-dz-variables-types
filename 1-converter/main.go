package main

import "fmt"

func main() {
	const usdEur = 1 * 0.9412
	const usdRub = 1 * 92
	const eurRub = usdEur * usdRub
	fmt.Println("Курс 1$ USD ⇄ 0.9412€ EUR =", usdEur, "EUR")
	fmt.Println("Курс 1$ USD ⇄ 92₽ RUB =", usdRub, "RUB")
	fmt.Println("Курс 0.9412€ EUR ⇄ 92₽ RUB =", eurRub, "RUB")
}
