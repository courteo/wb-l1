package tasks

import "fmt"

func T13() {
	a := 1
	b := 3
	fmt.Println(a, b)
	a = a + b			// a= first + second			
	b = a - b			// здесь a=first + second значит b = first
	a = a - b			// до a = first + second, после a = second
	fmt.Println(a, b)
}