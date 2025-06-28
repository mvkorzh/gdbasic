package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

func desk(x int) error {
	// Выполняем проверку введенного значения кол-ва клеток на сторону
	// Мы Не принимаем нечет и 0
	if x%2 != 0 || x == 0 {
		//fmt.Println("Мы печатаем шахматные доски только с четным кол-вом клеток")
		return errors.New("Мы печатаем шахматные доски только с четным кол-вом клеток")
	}

	// Размер доски нам уже известен, можем сформировать четные и нечетные строки
	even := strings.Repeat(" #", x/2)
	odd := strings.Repeat("# ", x/2)

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

	return nil
}

func main() {
	//var size = 1

	// Используем flag package для реализации флага -s
	var size = flag.Int("s", 8, "size of chess desk")
	flag.Parse()

	// Вызываем функцию печати доски и обрабатываем ошибку
	if err := desk(*size); err != nil {
		fmt.Println("Ошибка: ", err)
	}
}
