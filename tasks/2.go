package tasks

import (
	"fmt"
	"sync"
)

func T2() {
	arr := []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	for _, val := range arr {		
		wg.Add(1)													// увеличиваем счетчик на 1
		go func(val int, wg *sync.WaitGroup) {						// запускаем горутину и передаем счетчик и число
			defer wg.Done()											// уменьшаем счетчик

			fmt.Printf("Val is %d, Val^2 is %d\n",val, val*val)		// выводим результат
		}(val, wg)
	}
	wg.Wait()														// ждем, чтобы все горутины завершились и выходим из функции
	fmt.Println("Done")
}