package main

import (
	"fmt"
)

/*
Дается строка. Нужно удалить все символы, которые встречаются более одного раза и
вывести получившуюся строку
*/

func main() {
	var text string
	fmt.Scan(&text)

	counts := make(map[rune]int)

	for _, char := range text {
		counts[char]++
	}

	result := ""
	for _, char := range text {
		if counts[char] == 1 {
			result += string(char)
		}
	}

	fmt.Println(result)
}
