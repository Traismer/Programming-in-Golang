package main

import "fmt"

/*
Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается. Каждый год сумма вклада становится больше.
Определите, через сколько лет вклад составит не менее y рублей.
*/

func BankDeposit() {
	var x, p, y, result int
	fmt.Scan(&x, &p, &y)
	for year := 1; x < y; year++ {
		x = int(x + (x*p)/100)
		result = year
	}
	fmt.Println(result)
}

func main() {
	BankDeposit()
}
