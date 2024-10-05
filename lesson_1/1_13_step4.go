package main

import "fmt"

/*
Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток.
Например, если k=13257=3*3600+40*60+57, то h=3 и m=40.

Входные данные
На вход программе подается целое число k (0 < k < 86399).

Выходные данные
Выведите на экран фразу:
It is ... hours ... minutes.
Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.
*/

func DefinitionOfWholeHours() {
	var k int
	fmt.Scan(&k)

	hours := k / 3600
	minutes := (k % 3600) / 60

	fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)
}

func main() {
	DefinitionOfWholeHours()
}
