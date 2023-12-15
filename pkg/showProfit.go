package pkg

import (
	"bankCLI/pkg/models"
	"fmt"
)

func ShowProfit() {

	var balance float64
	var has bool
	// проверка на наличие клиента
	for _, val := range Accounts {
		if "profit" == val.Name {
			balance = val.Balance
			has = true
		}

	}

	if !has {
		Accounts = append(Accounts, &models.Account{
			Name:        "profit",
			Balance:     0.0,
			PhoneNumber: "544",
			City:        &Cities[0],
		})
	}

	fmt.Println(balance)
}
