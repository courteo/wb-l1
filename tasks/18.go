package tasks

import (
	"fmt"
	"sync"
)

type myAtomic struct {					// создаем структуру, которая содержит счетчик и мьютекс
	counter int
	mutex   sync.Mutex
}

func (m *myAtomic) GetCounter() int {	// лочим мьютекс и присваиваем счетчик переменной и анлочим мьютекс, чтобы другие горутины могли его получить
	m.mutex.Lock()
	res := m.counter
	m.mutex.Unlock()
	return res
}

func (m *myAtomic) Increase() {			// лочим мьютекс и увеличиваем счетчик и анлочим мьютекс, чтобы другие горутины могли его увеличить
	m.mutex.Lock()
	m.counter++
	m.mutex.Unlock()
}

func T18() {
	atom := &myAtomic{}
	wg := &sync.WaitGroup{}

	for i:=0; i <1000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {				// 1000 раз увеличиваем счетчик и используем вэйтгруппы для того, чтобы посмотреть какой счетчик после 1000 итераций
			defer wg.Done()
			atom.Increase()
		}(wg)
	}
	fmt.Println(atom.GetCounter())					// смотрим сколько насчитал счетчик, когда еще не завершились горутины
	wg.Wait()
	fmt.Println(atom.GetCounter())					// смотрим сколько насчитал счетчик в конце
}