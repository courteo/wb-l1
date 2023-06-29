package tasks

import "fmt"

func first(ch chan int, arr []int) {					// получает массив и канал и записывает в канал данные из массива
	for _, val := range arr {
		ch <- val
	}
	close(ch)											// закрываем канал
}

func second(chFirst, chSecond chan int) {				// получаем из первого канала данные и записываем во второй измененнные данные
	for val := range chFirst {
		chSecond <- val * 2
	}
	close(chSecond)										// закрываем второй канал
}

func T9() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chFirst := make(chan int)							// инициализируем каналы
	chSecond := make(chan int)
	
	go first(chFirst, arr)								// запускаем в горутине функцию и передаем туда первый канал и массив
	go second(chFirst, chSecond)						// запускаем в горутине функцию и передаем туда два канала

	for val := range chSecond {							// считываем из второго канала данные
		fmt.Println(val)
	}

}