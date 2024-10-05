package main

import "fmt"

/*
Даны два числа. Найти их среднее арифметическое.

Входные данные
На вход дается два целых положительных числа a и b.

Выходные данные
Программа должна вывести среднее арифметическое чисел a и b
(ответ может быть целым числом или дробным)
*/

func ArithmeticMeanOfNumbers() {
	var a, b int
	fmt.Scan(&a, &b)

	average := (float64(a) + float64(b)) / 2.0

	if average == float64(int(average)) {
		fmt.Printf("%d\n", int(average))
	} else {
		fmt.Printf("%.1f\n", average)
	}
}

func main() {
	ArithmeticMeanOfNumbers()
}
