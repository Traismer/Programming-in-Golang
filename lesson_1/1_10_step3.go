package main

import "fmt"

/*
Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B).
Вывести сумму всех чисел от A до B  включительно.
*/

func sumOfNaturalNumbers() {
	var A, B int
	fmt.Scan(&A, &B)
	fmt.Println((A + B) * (B - A + 1) / 2)
}

func main() {
	sumOfNaturalNumbers()
}
