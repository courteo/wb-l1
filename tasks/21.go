package tasks

import "fmt"

type someInterface interface {					// создали интерфейс с функцией print
	Print()
}

type f1 struct {
}

func (f *f1) Print() {							// есть некоторая структура, которая имеют функцию print и соответствует интерфейса
	fmt.Println("f1")
}

type f2 struct {
}

func (f *f2) PrintF2() {						// есть другая структура, которая имеют другую функцию и не соотвествует интерфейсу
	fmt.Println("f2")
}

type f2Adapter struct {							// создали адаптер для f2, чтобы он мог соответствовать интефейсу
	adapter  *f2
}

func (f *f2Adapter) Print() {					// функция print, чтобы адаптер соответствовал интерфейсу
	f.adapter.PrintF2()
}

type printer struct {
}

func (p *printer) WantToPrint(s someInterface) {	// функция принимающая наш интерфейс для проверки
	s.Print()
}

func T21() {
	printer := &printer{}
	f1 := &f1{}
	printer.WantToPrint(f1)							// f1 соответствует интерфейсу

	f2 := &f2{}
	f2Adapter := &f2Adapter{f2}
	printer.WantToPrint(f2Adapter)					// f2Adapter соответсвует интерфейсу, а f2 нет, поэтому используем адаптер
}