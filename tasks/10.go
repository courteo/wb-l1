package tasks

import "fmt"

func T10() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	res := make(map[int][]float64)											// мапа, где ключ это шаг, значение слайс температур
	for i := 0; i < len(arr); i++ {
		temp := int(arr[i])													// переводим температуру в инт
		temp = temp - temp%10												// отнимаем от темпы ее остаток от деление на 10, то есть разряд единиц, чтобы получить шаг
		if _, ok := res[temp]; !ok {										// если такого ключа нет, то создаем слайс
			res[temp] = make([]float64, 0)
		}
		res[temp] = append(res[temp], arr[i])								// добавляем к слайсу темпу
	}
	fmt.Println(res)
}