package main

import "fmt"

/*
Заданы три числа - a,b,c(a<b<c) - длины сторон треугольника.
Нужно проверить, является ли треугольник прямоугольным.
Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"

Входные данные
6 8 10

Выходные данные
Прямоугольный
*/

func DetermineWhichTriangle() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a*a+b*b == c*c || a*a+c*c == b*b || b*b+c*c == a*a {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}

func main() {
	DetermineWhichTriangle()
}
