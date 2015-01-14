package main

import (
	"fmt"
)

func main() {
	sa := []int{1, 2, 3}
	i := 0
	i, sa[i] = 1, 2
	// sets i = 1, sa[0] = 2.
	// Go 1 이전에는 왼쪽 i가 먼저 1이 되고 sa[i] 즉, sa[1]이 2가 할당되어 {1, 2, 3} 이 된다.
	// Go 1에서는 {2, 2, 3}
	fmt.Println(sa)

	sb := []int{1, 2, 3}
	j := 0
	sb[j], j = 2, 1
	// sets sb[0] = 2, j = 1
	// 이건 둘다 결과는 {2 2 3}, j == 1
	fmt.Println(sb, j)

	sc := []int{1, 2, 3}
	sc[0], sc[0] = 1, 2
	// sets sc[0] = 1, then sc[0] = 2 (so sc[0] = 2 at end)
	// 이건 둘다 결과는 {2 2 3}
	fmt.Println(sc)
}
