package tasks

import "fmt"

func binarySearch(neededVal int, arr []int) int {
	start, end := 0, len(arr)-1

	for start <= end {
		mid := start + (end-start)/2

		if arr[mid] < neededVal {			// если текущий элемент меньше нужного, значит он находится справа от текущего и мы увеличиваем стартпоинтер
			start = mid + 1
		} else if arr[mid] > neededVal {	// если текущий элемент больше нужного, значит он находится слева от текущего элемента и мы уменьшаем ендпоинтер
			end = mid - 1
		} else {							// мы нашли нужный элемент
			return mid
		}
	}
	return -1								// такого нет
}

func T17() {
	arr := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(binarySearch(8, arr))
	fmt.Println(binarySearch(9, arr))
}