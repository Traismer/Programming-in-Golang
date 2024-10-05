package main

import "fmt"

/*
По данным числам, определите количество чисел, которые равны нулю.

Входные данные
Вводится натуральное число N, а затем N чисел.

Выходные данные
Выведите количество чисел, которые равны нулю.
*/

func CounterNumbers() {
	var n, count int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		if num == 0 {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	CounterNumbers()
}
