package tasks

import (
	"fmt"
	"sync"
)

func T3() {
	array := []int{2, 4, 6, 8, 10, 12, 14}
	ch := make(chan int)									// создаем канал куда закидываем квадраты
	wg := &sync.WaitGroup{}									

	for _, val := range array { 							// проходимся по массиву 
		wg.Add(1)
		go func(val int, ch chan int, wg *sync.WaitGroup) { // запускаем горутину, где закидываем в канал все квадраты 
			defer wg.Done()									// и уменьшаем счетчик
			ch <- val*val
		}(val, ch, wg)
		
	}
	res := 0

	go func(wg *sync.WaitGroup, ch chan int) { 				// запускаем горутину, где ждем пока запишем все квадраты в канал
		wg.Wait()											
		close(ch)											// и закрываем канал, чтобы не было дэдлока
	}(wg, ch)

	for val := range ch {									// считываем с канала квадраты и суммируем
		res += val
	}
	
	fmt.Println(res)
}