package tasks

import "fmt"

func T11() {
	arr1 := []int{4, 1, 7, 18, 12, 2}
	arr2 := []int{18, 2, 13, -1, 4}

	res := make([]int, 0) 					// результирующий массив
	m := make(map[int]bool)					// мапка, где будем хранить ключи
	for _, val := range arr1 {
		m[val] = true						// при первом проходе кидать в мапу все значение
	}
	for _, val := range arr2 {				
		if _, ok := m[val]; ok {			// при втором проходе смотрим существует ли ключ в мапе, если да, то добавляем ключ в массив
			res = append(res, val)
		}
	}
	fmt.Println(res)
}