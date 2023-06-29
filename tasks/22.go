package tasks

import (
	"fmt"
	"sort"
	"strconv"
)

type BigInt struct {									// большое число это массив цифр с конца 
	number []int
}

func (b BigInt) GetLen() int {							// получить длину числа
	return len(b.number)
}

func (b BigInt) GetIndex(i int) (int, bool) {			// получить i цифру
	if i < 0 || i > b.GetLen() - 1 {
		return 0, false
	}
	return b.number[i], true
}

func (b *BigInt) SetDigit(pos, newDigit int) bool {		// назначить новую цифру
	if pos < 0 || pos > b.GetLen() - 1 {
		return false
	}
	b.number[pos] = newDigit
	return true
}

func (b *BigInt) IncreaseLen(count int) {				// увеличить длину числа 
	for i:=0; i < count; i++ {
		b.number = append(b.number, 0)
	}
}

func (b *BigInt) Reverse() {							// перевернуть число 
	sort.SliceStable(b.number, func(i, j int) bool {
        return true
    })
}

func (b *BigInt) DeleteLast() {							// удалить последнюю цифру
	b.number = b.number[:b.GetLen() - 1]
}

func CreateBigInt(a string) (BigInt, error) {			// создать большое число из строки
	res := BigInt{
		number: make([]int, 0),
	}
	for i:=len(a) - 1; i >= 0; i-- {
		temp, err := strconv.Atoi(string(a[i]))
		if err != nil {
			return BigInt{}, fmt.Errorf("Not a Number\n")
		}
		res.number = append(res.number, temp)
	}
	return res, nil
}

func (b BigInt) plus(operand BigInt) BigInt {
	
	newLen := 0
	if b.GetLen() > operand.GetLen() {					// назначаем новую длину числу
		newLen = b.GetLen() + 1
	} else {
		newLen = operand.GetLen() + 1
	}

	c := b
	c.IncreaseLen(newLen - c.GetLen())

	for i:=0; i < newLen; i++ {
		first, _ := c.GetIndex(i)
		second, _ := operand.GetIndex(i)
		nextDigit, _ := c.GetIndex(i + 1)
		c.SetDigit(i, first + second )			// меняем i позицию на сумму цифр

		newDigit, _ := c.GetIndex(i)
		c.SetDigit(i + 1, newDigit / 10 + nextDigit)	// проверяем нужно ли добавить на след разряд цифру
		c.SetDigit(i, newDigit % 10)					// если сумма цифр больше 9 то убираем разряд десятков
	}
	if digit, _ := c.GetIndex(newLen - 1); digit == 0{ 	// если новый разряд не использовался, то убираем его
		c.DeleteLast()
	}
	return c
}

func compare(first, second BigInt) int { // 1 - первое больше, -1 - второе больше, 0 - равны
	if first.GetLen() > second.GetLen() { // первое по длине больше
		return 1
	}

	if second.GetLen() > first.GetLen() { // второе по длине больше
		return 0
	}

	for i:=first.GetLen() - 1; i >= 0; i-- { // поразрядно проверяем какое больше
		f, _ := first.GetIndex(i)
		s, _ := second.GetIndex(i)
		if f > s {
			return 1
		}
		if s > f {
			return -1
		}
	}

	return 0								// равны
}

func subIfFirstBigger(first, second BigInt) BigInt {	// вычитает числа с учетом, что первое всегда!!! больше
	res := BigInt{
		number: make([]int, first.GetLen()),
	}
	flag := 0
	for i:=0; i < second.GetLen(); i++ {				// поразрядно отнимаем
		f, _ := first.GetIndex(i)
		s, _ := second.GetIndex(i)
		if f - s - flag< 0 {							// если нужно будет взять единицу и след разряда
			res.SetDigit(i, 10 + f - s - flag)
			flag = 1
		} else {
			res.SetDigit(i, f - s - flag)
			flag = 0
		}
	}

	for i:=second.GetLen(); i < first.GetLen(); i++{	//
		f, _ := first.GetIndex(i)
		if f - flag < 0 {
			res.SetDigit(i, 10 + f - flag)
			flag = 1
		} else {
			res.SetDigit(i, f - flag)
			flag =0
		}
	}

	newLen := res.GetLen()
	for {												// убираем ненужные нули
		if digit, _ := res.GetIndex(newLen - 1); digit == 0 {
			newLen--
			res.DeleteLast()
		} else {
			break
		}
		
	}
	return res
}

func (b BigInt) sub(operand BigInt) BigInt {
	var res BigInt
	comp := compare(b, operand)						// сравниваем числа
	if comp == 0 {									// если равны то возвращаем 0
		return BigInt{
			number:  []int{0},
		}
	}
	if comp == 1 {									// если первое больше
		res = subIfFirstBigger(b, operand)
	} else {
		res = subIfFirstBigger(operand, b)			// есил второе больше
	}

	return res
}

func (b BigInt) mult(operand BigInt) BigInt {

	oldLen := b.GetLen()
	newLen := b.GetLen() + operand.GetLen() + 1			// определяем новую длину числа
	
	res := BigInt{
		number: make([]int, newLen),
	}

	for i:= 0; i < oldLen; i++ {
		for j := 0; j < operand.GetLen(); j++ {
			first, _ := b.GetIndex(i)
			second, _ := operand.GetIndex(j)
			old, _ := res.GetIndex(i + j)
			res.SetDigit(i + j, first*second + old)		// увеличиваем i+j цифру на b[i] + operand[j]
		}
	}

    for i := 0; i < newLen; i++{						// проверяем нет ли элементов больших 9
		temp, _ := res.GetIndex(i)
		old, _ := res.GetIndex(i + 1)
		res.SetDigit(i + 1, temp / 10 + old)
		res.SetDigit(i, temp % 10)
	}

	for {												// убираем ненужные нули
		if digit, _ := res.GetIndex(newLen - 1); digit == 0 {
			newLen--
			res.DeleteLast()
		} else {
			break
		}
		
	}
	return res
}

func (b BigInt) div(operand BigInt) BigInt {

	if compare(operand, BigInt{[]int{0}}) == 0 {	// на ноль нельзя делить
		return BigInt{}
	}

	if compare(b, operand) == -1 {					// если первый операнд меньше второго
		return BigInt{[]int{0}}
	}
	newOperand := operand
	res := BigInt{[]int{1}}
	for compare(b, newOperand) > 0 {				// пока больше подбираем ответ
		res = res.plus(BigInt{[]int{1}})
		
		newOperand = newOperand.plus(operand)
	}

	if compare(b, newOperand) == -1 {				// уменьшаем если перескочили
		res = res.sub(BigInt{[]int{1}})
	}

	return res
}

func T22() {
	var a, b string
	fmt.Scan(&a, &b)
	first, err := CreateBigInt(a)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	second, err := CreateBigInt(b)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(first, second)
	
	fmt.Println(first.div(second))
}