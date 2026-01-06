package main

import "fmt"

/*
	You are given an integer array order of length n and an integer array friends.
	order contains every integer from 1 to n exactly once, representing the IDs of the participants of a race in their finishing order.
	friends contains the IDs of your friends in the race sorted in strictly increasing order. Each ID in friends is guaranteed to appear in the order array.
*/

func main() {
	order := []int{3,1,2,5,4}
	firends := []int{1,3,4}
	idx := 0
	finishing := make([]int, 0)
	for i := 0; i < len(firends); i++ {
		if order[idx] == firends[i] {
			finishing = append(finishing, order[i])
			i = 0
			idx = idx + 1
			continue
		}
		if idx == len(order) - 1 && i == len(firends) - 1 {
			break
		}
	}
	fmt.Println(finishing)
}