package main

import "fmt"

/*
На ввод подаются пять целых чисел, которые записываются в массив. Однако эта часть программы уже написана.
Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.
*/

func FiveIntTheMaxNumber() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	max := array[0]
	for i := 1; i < 5; i++ {
		if array[i] > max {
			max = array[i]
		}
	}

	fmt.Println(max)
}

func main() {
	FiveIntTheMaxNumber()
}
