package main

import "fmt"

/*
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.
*/

func SumOfTwosEight() {
	var count, arr, result int
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		fmt.Scan(&arr)
		if arr >= 10 && arr <= 99 && arr%8 == 0 {
			result += arr
		}
	}
	fmt.Println(result)
}

func main() {
	SumOfTwosEight()
}
