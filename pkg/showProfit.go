package pkg

import "fmt"

func ShowProfit() {
	balance, ok := Database["Profit"]

	if !ok {
		Database["Profit"] = 0
	}

	fmt.Println("Прибыль банка составляет:", balance)
}
