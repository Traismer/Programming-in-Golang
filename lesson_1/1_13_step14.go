package main

import "fmt"

/*
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является,
то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.

Входные данные
Вводится натуральное число.

Выходные данные
Выведите ответ на задачу.
*/

func Fibonacci() {
	var A int
	fmt.Scan(&A)

	if A <= 1 {
		fmt.Println(-1)
		return
	}

	f1, f2 := 1, 1
	n := 2

	for f2 < A {
		f1, f2 = f2, f1+f2
		n++
	}

	if f2 == A {
		fmt.Println(n)
	} else {
		fmt.Println(-1)
	}
}

func main() {
	Fibonacci()
}
