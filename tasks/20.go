package tasks

import (
	"bufio"
	"fmt"
	"os"
)

func T20() {
	reader := bufio.NewReader(os.Stdin)	
	s, _ := reader.ReadString('\n')									// считываем строку	
	
	words := make([]string, 0)
	var temp string

	for  _, val := range s {		
		if val >= 'a' && val <= 'z' || val >= 'A' && val <= 'Z' {	// проверяем буква или нет и добавляем к переменной
			temp += string(val)
		} else {													// если не буква, то
			if temp != "" {											// проверяем пуста ли переменна, если пуста, то добавляем к слайсу строк					
				words = append(words, temp)
				temp = ""
			}
			if val != '\n' {										// добавляем к слайсу другие символы кроме букв и перевода строки
				words = append(words, string(val))
			}
			
		}
	}
	
	res := ""
	for i:=len(words) - 1; i >=0; i-- {								// проходим по слайсу с конца и добавляем к результату
		res += words[i]
	}
	
	fmt.Println(res)
}