package main

import (
	"fmt"
	"strings"
)

/*
Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X.
Если подстроки S нет в строке X - вывести -1
*/

func FindSubstringFirstOccurrence(X, S string) int {
	index := strings.Index(X, S)
	if index == -1 {
		return -1
	}
	return index
}

func main() {
	var X, S string
	fmt.Scan(&X, &S)
	fmt.Println(FindSubstringFirstOccurrence(X, S))
}
