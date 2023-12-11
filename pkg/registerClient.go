package pkg

import "fmt"

func RegisterClient() {
	var (
		name string
	)
	fmt.Println("Введите имя клиента")
	fmt.Scan(&name)
	Database[name] = 0
}
