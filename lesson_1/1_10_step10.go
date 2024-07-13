package main

import "fmt"

/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются.
Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения в первом числе.
*/

func OccurrenceOfNumber() {
	var first, last string
	fmt.Scan(&first, &last)

	for i := 0; i < len(first); i++ {
		for j := 0; j < len(last); j++ {
			if first[i] == last[j] {
				fmt.Printf("%c ", first[i])
				break
			}
		}
	}
}

func main() {
	OccurrenceOfNumber()
}
