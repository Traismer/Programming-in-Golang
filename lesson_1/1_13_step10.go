package main

import "fmt"

/*
Цифровой корень натурального числа — это цифра,
полученная в результате итеративного процесса суммирования цифр,
на каждой итерации которого для подсчета суммы цифр берут результат,
полученный на предыдущей итерации.

Этот процесс повторяется до тех пор, пока не будет получена одна цифра.
Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .
По данному числу определите его цифровой корень.

Входные данные
Вводится одно натуральное число n, не превышающее 10 в 7 степени

Выходные данные
Вывести цифровой корень числа n.
*/

func DigitalRoot() {
	var n int
	fmt.Scan(&n)

	for n >= 10 {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		n = sum
	}

	fmt.Println(n)
}

func main() {
	DigitalRoot()
}
