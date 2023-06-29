package tasks

import "fmt"

type HumanStats struct { // вес и рост человека
	weight int
	heght  int
}

type Human struct { // Структура человек с именем и физическими данными
	name  string
	stats HumanStats
}

func (h Human) GetName() string { // получение имени человека
	return h.name
}

func (h Human) PrintStats() { // вывод данных человека
	fmt.Printf("Weight: %d, Height: %d\n", h.stats.weight, h.stats.heght)
}

type Action struct { // структура содержит только человека
	Human
}

func T1() {
	stats := HumanStats{
		weight: 60,
		heght:  170,
	}

	human := Human{
		name:  "John",
		stats: stats,
	}

	action := Action{
		Human: human,
	}
	fmt.Println(action.GetName())
	action.PrintStats()
}