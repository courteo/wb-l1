package tasks

import "fmt"

func T19() {
	var s string
	fmt.Scan(&s)
	runes := []rune(s)
    for i :=0; i < len(runes)/2; i++ {												// проходимся до середины строки и меняем текущий элемент с элементом зеркальной позиции относительно середины местами
		runes[i], runes[len(runes) - 1 - i] = runes[len(runes) -1 - i], runes[i]
	}
	fmt.Println(string(runes))
}