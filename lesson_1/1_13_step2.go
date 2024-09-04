package main

import "fmt"

/*
Дано трехзначное число. Найдите сумму его цифр.

Формат входных данных
На вход дается трехзначное число.

Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.
*/
func SumAllNumbers() {
	var number int
	fmt.Scan(&number)
	// получаем первое число + второе число + третье число
	sum := number/100 + (number%100)/10 + number%10
	fmt.Println(sum)
}

func main() {
	SumAllNumbers()
}
