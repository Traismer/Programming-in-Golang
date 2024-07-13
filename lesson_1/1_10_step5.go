package main

import "fmt"

/*
Последовательность состоит из натуральных чисел и завершается числом 0.
Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.
*/

func CountMaxElements() {
	var arr, count, max int
	for fmt.Scan(&arr); arr != 0; fmt.Scan(&arr) {
		if arr > max {
			max = arr
			count = 1
		} else if arr == max {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	CountMaxElements()
}
