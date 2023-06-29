package tasks

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
)

func insertData(ch chan int) {												// бесконечно записывать в канал рандомное число
	for {
		ch <- rand.Intn(1000)
	}
}

func worker(ch chan int, number int) {										// бесконечно(т.к. канал не закрыт) считывает из канала число
	for val := range ch{
		fmt.Printf("Worker %d has got %d\n",number, val)
	}
}

func T4() {
	var n int
	fmt.Scan(&n)
	ch := make(chan int)

	for i:=0; i < n; i++ {													// заводим n воркеров
		go worker(ch, i + 1)
	}

	go insertData(ch)														// запускаем поток данных

	c := make(chan os.Signal, 1)											// создаем канал который принимает систем колы                                       
	signal.Notify(c, os.Interrupt)  										// функция ловит систем колы и записывает в канал
	for sig := range c {                                             		// если в канале систем кол то завершаем программу
		log.Printf("captured %v, stopping profiler and exiting..", sig)	
		os.Exit(1)                                                     
	} 
}