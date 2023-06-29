package tasks

import "fmt"

func partition(arr []int, start, end int) ([]int, int) {		// выбирается последний элемент и все значения, которые меньше его
	pivot := arr[end]											// будут находиться слева от него, а больше справа от него по окончанию функции
	i := start													// выводится индекс разделяющего элемента

	for j := start; j < end; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]
	return arr, i
}

func quickSort(arr []int, start, end int) []int {				// сортирует массив относительно разделяющего элемента
																// и потом сортирует две части относительно разделяющего элемента
	if start < end {
		var pos int

		arr, pos = partition(arr, start, end)
		arr = quickSort(arr, start, pos-1)
		arr = quickSort(arr, pos+1, end)
	}

	return arr
}

func T16() {
	arr := []int{8, 1, 20, 14, 16, 2,-10}
	fmt.Println(arr)
	fmt.Println(quickSort(arr, 0, len(arr)-1))
}