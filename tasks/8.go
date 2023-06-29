package tasks

import (
	"fmt"
	"strconv"
)

func T8() {
	var num int64 = 24

	bits := strconv.FormatInt(num, 2)			// переводим число в строку битов для проверки ( в решении не участвует)
	fmt.Printf("Old bits %s\n",bits)
	
	var i int
	var neededBit,res int64
	fmt.Scan(&neededBit, &i)
	
	
	if neededBit == 1 {
		res = num | (1 << i)					// Возвращает 1, если хотя бы один из соответствующих разрядов обоих чисел равен 1
	} else {
		res = num &^ (1 << i)					//  каждый бит res равен 0, если соответствующий бит (1 << i) равен 1. Если бит в (1 << i) равен 0, то берется значение соответствующего бита из num.
	}
	
	fmt.Println(res, strconv.FormatInt(res, 2))
}