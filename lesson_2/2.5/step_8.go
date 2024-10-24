package main

import (
	"fmt"
	"strings"
)

/*
На вход подается строка. Нужно определить, является ли она палиндромом.
Если строка является палиндромом - вывести Палиндром иначе - вывести Нет.
Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях
(например, "топот", "око", "заказ").
*/

func IsPalindrome() {
	var input string
	fmt.Scan(&input)

	input = strings.ToLower(input) // Переводим в нижний регистр

	reversed := []rune(input)
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	if string(reversed) == input {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}

func main() {
	IsPalindrome()
}
