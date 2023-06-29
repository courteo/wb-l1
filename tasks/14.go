package tasks

import (
	"fmt"
	"reflect"
)

func getType(t interface{}) {						// смотрим через свитч
	switch t.(type) {
	case int:
		fmt.Println("Got int")
	case string:
		fmt.Println("Got string")
	case bool:
		fmt.Println("Got bool")
	case chan int:
		fmt.Println("Got chan int")
	default:
		fmt.Println("Got unknown type")
	}
}

func getTypeWithReflect(t interface{}) {			// через рефлект
	fmt.Println("Got", reflect.TypeOf(t))
}

func T14() {
	var a int
	var b string
	var c bool
	var d float64
	var ch chan int
	ch1 := make(chan string)
	var m map[string]string
	
	getType(a)
	getType(b)
	getType(c)
	getType(d)
	getType(ch)
	getTypeWithReflect(ch1)
	getTypeWithReflect(m)
}