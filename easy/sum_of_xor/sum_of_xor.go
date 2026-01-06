package main

import "fmt"

func main() {
	expra := []int{1, 3}
	mk := make([][]int, 1 << len(expra))
	// [1, 3], [1, 4], [3, 4], [1, 3, 4]
	tmp := 0
	for i := 0; i < 1 << len(expra); i++ {
		mk[i] = make([]int, 0)
		for j := 0; j < len(expra); j++ {
			if (i & (1 << j)) != 0 {
				mk[i] = append(mk[i], expra[j])
				fmt.Println(expra[j])
			}
		}
	}
	fmt.Println(mk)
}