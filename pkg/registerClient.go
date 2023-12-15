package pkg

import (
	"bankCLI/pkg/models"
	"fmt"
)

func RegisterClient() {
	var name string

	fmt.Println("Введите имя клиента")

	fmt.Scan(&name)

	var phone string

	fmt.Println("Введите номер телефона")

	fmt.Scan(&phone)

	var cityName string

	fmt.Println("Введите название города")

	fmt.Scan(&cityName)

	var city models.City
	var has bool
	for _, v := range Cities {
		if v.Name == cityName {
			city = v
			has = true
		}
	}

	if !has {
		fmt.Println("Город не найден в нашей бд, обратитесь по номеру 544")
		return
	}

	Accounts = append(Accounts, &models.Account{
		Name:        name,
		Balance:     0.0,
		PhoneNumber: phone,
		City:        &city,
	})
}
