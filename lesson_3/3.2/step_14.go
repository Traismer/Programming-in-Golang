/*
В привычных нам редакторах электронных таблиц присутствует удобное представление числа
с разделителем разрядов в виде пробела, кроме того в России целая часть
от дробной отделяется запятой. Набор таких чисел был экспортирован в формат CSV,
где в качестве разделителя используется символ ";".

На стандартный ввод вы получаете 2 таких вещественных числа,
в качестве результата требуется вывести частное от деления первого числа на второе
с точностью до четырех знаков после "запятой" (на самом деле после точки,
результат не требуется приводить к исходному формату).
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Читаем строку из стандартного ввода
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Убираем лишние пробелы и символы новой строки
	input = strings.TrimSpace(input)

	// Разделяем строку на два числа по ";"
	parts := strings.Split(input, ";")
	if len(parts) != 2 {
		fmt.Println("Invalid input")
		return
	}

	// Функция для преобразования строки в число
	convert := func(s string) (float64, error) {
		// Убираем пробелы и заменяем запятую на точку
		s = strings.ReplaceAll(s, " ", "")
		s = strings.ReplaceAll(s, ",", ".")
		return strconv.ParseFloat(s, 64)
	}

	// Преобразуем каждую часть в число
	num1, err1 := convert(parts[0])
	num2, err2 := convert(parts[1])

	if err1 != nil || err2 != nil || num2 == 0 {
		fmt.Println("Error in conversion or division by zero")
		return
	}

	// Вычисляем частное
	result := num1 / num2

	// Выводим результат с точностью до четырех знаков после запятой
	fmt.Printf("%.4f\n", result)
}
