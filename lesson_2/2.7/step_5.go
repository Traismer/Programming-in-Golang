/*
На вход подается целое число.
Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.

Например, у нас есть число 9119.
Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1.
В итоге получаем 811181
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)

	var result strings.Builder

	for _, char := range input {
		digit := int(char - '0')
		squared := digit * digit
		result.WriteString(fmt.Sprintf("%d", squared))
	}

	fmt.Println(result.String())
}
