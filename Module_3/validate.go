package main

import (
	"errors"
	"strings"
	"unicode"
)

func (u User) Validate() error {
	if len(u.Name) < 2 || len(u.Name) > 50 {
		return errors.New("имя не может содержать менее 2 букв или более 50 букв")
	}

	for _, r := range u.Name {
		if !unicode.IsLetter(r) {
			return errors.New("имя должно содержать только буквы")
		}
	}

	if !strings.Contains(u.Email, "@") || len(u.Email) < 6 {
		return errors.New("email должен содержать @ и иметь длину более 6 символов")
	}

	if u.Age < 0 || u.Age > 110 {
		return errors.New("возраст не может быть отрицательным или больше 110")
	}

	if u.Gender == "" {
		return errors.New("необходимо задать пол")
	}

	if len(u.Phone) < 6 || len(u.Phone) > 30 {
		return errors.New("длина номера телефона должна быть от 6 до 30 символов")
	}

	firstSymPhone := []rune(u.Phone)[0]
	if firstSymPhone != '+' {
		return errors.New("номер телефона должен начинаться с +")
	}

	return nil
}
