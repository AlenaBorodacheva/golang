package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	user := User{
		Name:   "Alice",
		Email:  "alice@example.com",
		Age:    30,
		Phone:  "+7(999)-999-99-99",
		Gender: Female}

	if err := user.Validate(); err != nil {
		fmt.Println("Ошибка валидации:", err)
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}
	fmt.Println("JSON:", string(data))

	var parsedUser User
	err = json.Unmarshal(data, &parsedUser)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}

	fmt.Printf("Parsed User: %+v\n", parsedUser)
}
