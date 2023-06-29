package tasks

import "fmt"

func T23() {
	slice := []int{2, 4, 8, 10}
	var i int
	fmt.Scan(&i)
	if i >= 0 && i < len(slice) {					// проверяем есть ли такой индекс
		slice = append(slice[:i], slice[i+1:]...)
	}
 	
	fmt.Println(slice)
}