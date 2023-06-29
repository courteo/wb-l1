package tasks

import "fmt"

func T26() {
	var s string
	fmt.Scan(&s)
	m := make(map[rune]struct{})		// массив рун и пустых структур

	for _, val := range s {
		if _, ok := m[val]; ok {		// если такой ключ есть, значит символ повтряется
			fmt.Println("false")
			return
		} else {
			m[val] = struct{}{}			// заводим ключ
		}
	}
	fmt.Println("true")
	return
}