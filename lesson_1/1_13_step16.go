package main

import "fmt"

/*
Из натурального числа удалить заданную цифру.

Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.

Выходные данные
Вывести число без заданных цифр.
*/

func DeleteTheSpecifiedNumber() {
	var number string
	var digit rune // это алиас int32
	fmt.Scan(&number)
	fmt.Scanf("%c", &digit)

	result := ""

	for _, ch := range number {
		if ch != digit {
			result += string(ch)
		}
	}

	fmt.Println(result)
}

func main() {
	DeleteTheSpecifiedNumber()
}
