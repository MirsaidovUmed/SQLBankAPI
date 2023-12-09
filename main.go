package main

import (
	"fmt"
	"time"
)

var database = make(map[string]int)

func createCli() {
	var (
		name    string
		balance int
	)
	fmt.Scan(&name, &balance)
	database[name] = balance

	fmt.Println("________________")
	fmt.Println("Готово")
	fmt.Println("________________")
}

func topUpClientBalance(name string) {
	var refill int
	fmt.Scan(&refill)
	balance, exists := database[name]
	if !exists {
		fmt.Println("Клиент не найден")
		return
	}
	database[name] = balance + refill
	fmt.Println("________________")
	fmt.Println("Готово")
	fmt.Println("________________")
}

func showClientBalance() {
	for key, val := range database {
		fmt.Println(key, val)
	}
	fmt.Println("______________")
}

func withdrawMoneyFromBalance(name string) {
	var takeOff int
	fmt.Scan(&takeOff)
	balance, exists := database[name]
	if !exists {
		fmt.Println("Клиент не найден")
		return
	}
	if balance >= takeOff {
		database[name] = balance - takeOff
	} else {
		fmt.Println("Извините недостаточно средств")
	}
	fmt.Println("________________")
	fmt.Println("Готово")
	fmt.Println("________________")
}

func main() {
	var name string
	for {
		var choice int
		fmt.Println("1. Создать клиента")
		fmt.Println("2. Пополнить баланс клиента")
		fmt.Println("3. Показать всех клиентов")
		fmt.Println("4. Снять деньги с баланса клиента")
		fmt.Println("5. Выйти")

		fmt.Scan(&choice)

		if choice == 1 {
			createCli()
		} else if choice == 2 {
			fmt.Scan(&name)
			topUpClientBalance(name)
		} else if choice == 3 {
			showClientBalance()
		} else if choice == 4 {
			fmt.Scan(&name)
			withdrawMoneyFromBalance(name)
		} else if choice == 5 {
			break
		}

		time.Sleep(1 * time.Second)

	}
}
