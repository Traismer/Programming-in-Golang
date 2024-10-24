package main

import (
	"fmt"
	"unicode"
)

/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
Длина пароля должна быть не менее 5 символов, он может содержать только арабские цифры и
буквы латинского алфавита. На вход подается строка-пароль.
Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
*/

func main() {
	var password string
	fmt.Scan(&password)

	if isValidPassword(password) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}

func isValidPassword(password string) bool {
	if len(password) < 5 {
		return false
	}

	for _, char := range password {
		if !isLatinLetterOrDigit(char) {
			return false
		}
	}

	return true
}

func isLatinLetterOrDigit(char rune) bool {
	return ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') || unicode.IsDigit(char)
}
