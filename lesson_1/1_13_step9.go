package main

import "fmt"

/*
Найдите количество минимальных элементов в последовательности.

Входные данные
Вводится натуральное число N, а затем N целых чисел последовательности.

Выходные данные
Выведите количество минимальных элементов последовательности.
*/
func NumberOfMinimalElemants() {
	var N int
	fmt.Scan(&N)

	elements := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&elements[i])
	}

	min := elements[0]
	countMin := 1

	// Ищем минимальное значение и его частоту
	for i := 1; i < N; i++ {
		if elements[i] < min {
			min = elements[i]
			countMin = 1
		} else if elements[i] == min {
			countMin++
		}
	}

	fmt.Println(countMin)
}

func main() {
	NumberOfMinimalElemants()
}
