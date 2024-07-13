package main

import "fmt"

/*
Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.
*/

func FirstNumber() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)
	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
	}
}

func main() {
	FirstNumber()
}
