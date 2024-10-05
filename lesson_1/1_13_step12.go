package main

import "fmt"

/*
По данному числу n закончите фразу "На лугу пасется..."
одним из возможных продолжений: "n коров", "n корова", "n коровы",
правильно склоняя слово "корова".

Входные данные
Дано число n (0<n<100).

Выходные данные
Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy,
например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.
*/

func DeclensionOfTheWordCow() {
	var n int
	fmt.Scan(&n)

	if n%10 == 1 && n%100 != 11 {
		fmt.Printf("%d korova\n", n)
	} else if (n%10 >= 2 && n%10 <= 4) && !(n%100 >= 12 && n%100 <= 14) {
		fmt.Printf("%d korovy\n", n)
	} else {
		fmt.Printf("%d korov\n", n)
	}
}

func main() {
	DeclensionOfTheWordCow()
}
