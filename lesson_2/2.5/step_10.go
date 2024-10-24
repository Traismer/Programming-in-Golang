package main

import (
	"fmt"
	"strings"
)

/*
На вход дается строка, из нее нужно сделать другую строку,
оставив только нечетные символы (считая с нуля)
*/

func main() {
	var input string
	fmt.Scan(&input)

	var result strings.Builder
	for i := 1; i < len(input); i += 2 {
		result.WriteByte(input[i])
	}

	fmt.Println(result.String())
}
