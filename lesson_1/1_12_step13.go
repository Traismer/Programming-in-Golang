package main

import "fmt"

/*
Напишите программу, принимающая на вход число N(N≥4), а затем N целых чисел-элементов среза.
На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
*/

func NSliceIndex() {
	var n int
	fmt.Scan(&n)

	var slice []int
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		slice = append(slice, num)
	}

	fmt.Println(slice[3]) // 3-ий элемент с индексом 0
}

func main() {
	NSliceIndex()
}
