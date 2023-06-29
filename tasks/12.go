package tasks

import "fmt"

func T12() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, val := range arr {
		if _, ok := set[val]; !ok {
			set[val] = struct{}{}
		}
	}
	fmt.Println(set)
}