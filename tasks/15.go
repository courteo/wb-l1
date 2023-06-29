package tasks

/*
	Убрали глобальную переменную, т.к. глобальная переменная может быть изменена в любой точке программы, 
	что может повлиять на работу других частей программы.
	Ошибки вызванные использованием глобальных переменных сложно отловить
*/


func createHugeString(val int) string {
	res := "1"
	
	for i:=0; i< val; i++ {
		res += "1"
	}

	return res
}


func someFunc1() string {
	return createHugeString(1<<10)[:100]
}

func T15() {
	someFunc1()
}
