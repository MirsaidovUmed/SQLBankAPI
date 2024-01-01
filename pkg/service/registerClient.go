package service

import (
	"fmt"
)

func (s *Service) RegisterClient() {
	var name string

	fmt.Println("Введите имя клиента")

	fmt.Scan(&name)

	var phone string

	fmt.Println("Введите номер телефона")

	fmt.Scan(&phone)

	var cityName string

	fmt.Println("Введите название города")

	fmt.Scan(&cityName)

	// get city
	city, err := s.Repository.GetCity(cityName)

	if err != nil {
		fmt.Println("Ошибка, данный город не найден, звоните по номеру 544")
		return
	}

	s.Repository.AddAccount(name, 0.0, phone, city)

	fmt.Println("Успешно зарегистрирован клиент!")
}
