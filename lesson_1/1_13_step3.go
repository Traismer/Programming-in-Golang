package main

import "fmt"

/*
Дано трехзначное число. Переверните его, а затем выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных
Выведите перевернутое число.
*/

func InvertedNumber() {
	var number int
	fmt.Scan(&number)
	fmt.Printf("%d%d%d\n", number%10, (number%100)/10, number/100)
}

func main() {
	InvertedNumber()
}
