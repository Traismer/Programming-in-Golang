package main

import "fmt"

/*
Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10. Квадрат каждого числа должен выводится в новой строке.
*/

func SquaresOfNaturalNumbers(number int) {
	for i := 1; i <= number; i++ {
		fmt.Println(i * i)
	}
}

func main() {
	SquaresOfNaturalNumbers(10)
}
