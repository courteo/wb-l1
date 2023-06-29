package tasks

import (
	"fmt"
	"time"
)

func T5() {
	var n int
	fmt.Scan(&n)
	ch := make(chan int)
	ch1 := make(chan bool)

	go func(ch1 chan bool) {							// запускаем горутину, которая засыпает n секунд
		time.Sleep(time.Duration(n) * time.Second)		// а потом передает каналу флажок
		ch1 <- true
	}(ch1)

	go func(ch chan int) {								// запускаем горутину, где в канал кидаем числа
		
		for i:=0; i< 10000; i++ {
			ch <- i*i*i
		}
		
		close(ch)										// закрываем канал
	}(ch)

	for val := range ch {								// считываем данные из канала
		select {										// проверяем кинули ли в во второй канал флажок
		default:										// если нет, то выводим число
			fmt.Println(val)
		case <-ch1:										// если да, то выводим таймаут и завершаем программу
			fmt.Println("Timeout")
			return
		}
		
	}
}