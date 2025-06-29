package main

// Программа выводит шахматную доску указанным размером (по умолчанию 8x8), размер доски не проверяется (допустимо 0x0 и 1x1)

import (
	"flag"
	"fmt"
	"strings"
)

func desk(x int) {
	// Размер доски нам уже известен, можем сформировать четные и нечетные строки
	// even - четная строка, odd - нечетная. Инициализируем неявно дефолтным значением типа.
	var even, odd string
	// Для формаирования строки используем string.Builder - объявляем две переменные типа string.Builder (even_sb - формируем четные строки, odd_sb - нечетные)
	// Наверное правильно было бы на этом этапе использовать конкатенацию, но я хотел попробовать использовать string.Builder,
	// тем более, что он уже упоминаллся на занятии "Работа со строками и рунами utf-8"
	var even_sb, odd_sb strings.Builder

	// Формируем строку. Строка размером x символов, чередование # и _ зависит от четности строки
	for i := range x {
		if i%2 == 0 {
			even_sb.WriteString("#")
			odd_sb.WriteString(" ")
		} else {
			even_sb.WriteString(" ")
			odd_sb.WriteString("#")
		}
	}

	// Возвращаем строку и присваиваем ранее объявленным переменным even и odd
	even = even_sb.String()
	odd = odd_sb.String()

	// Выводим x строк, нечетная - odd, четная - even
	//for i := 0; i < x; i++ {
	// IDE предложил заменить на modernized rangeint
	for i := range x {
		if i%2 == 0 {
			fmt.Println(odd)
		} else {
			fmt.Println(even)
		}
	}
}

func main() {
	// Используем flag package для реализации флага -s, flag.Int возвращает указатель на переменную int, в которой хранится значение флага
	var size = flag.Int("s", 8, "size of chess desk")
	flag.Parse()

	// Вызываем функцию печати доски, size - указатель типа int, передаем в функцию разыменованный указатель
	desk(*size)
}
