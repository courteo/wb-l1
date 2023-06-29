package tasks

import (
	"context"
	"fmt"
	"time"
)

func ThroughChannel() {							// закрываем через канал
	ch := make(chan bool)

	go func(ch chan bool) {						// запускаем горутину, где на секунду засыпаем, а потом кидаем в канал флажок
		time.Sleep(1 * time.Second)
		ch <- true
	}(ch)

	select {									// ждем пока в канал кинут флажок и выходим из функции
	case <-ch:
		fmt.Println("ThroughChannel")
		return
	}
}

func ThroughContext() {											// закрываем через контекст
	ctx, stop := context.WithCancel(context.Background())

	go func () {
		if err := someFunc(ctx); err != nil {					// ждем ответа от функции, где бесконечный цикл
			fmt.Println("ThroughContext")
		}
	}()

	time.Sleep(1 * time.Second)
		
	stop()														// отменяем контекст
	time.Sleep(5 * time.Second)
}

func someFunc(ctx context.Context) error {
	for {														// бесконечный цикл, где проверяется контекст на ошибку
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}



func T6() {
	ThroughContext()
	ThroughChannel()
	// третий вариант закрыть main(), тогда все горутины завершатся
}