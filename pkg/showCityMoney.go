package pkg

import (
	"fmt"
)

func ShowCityMoney() {
	var city string
	var result float64
	fmt.Scan(&city)

	for _, val := range Accounts {
		if val.City.Name == city {
			result += val.Balance
		}
	}

	fmt.Println(result)
}
