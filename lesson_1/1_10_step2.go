package main

import . "fmt"

/*
Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10. Квадрат каждого числа должен выводится в новой строке.
*/

func SquaresOfNaturalNumbers() {
	for number := 1; number <= 10; number++ {
		Println(number * number)
	}
}

func main() {
	SquaresOfNaturalNumbers()
}
