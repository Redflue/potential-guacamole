package potentialguacamole

import "math/rand"

func ArrayGenerate(s int) []int {
	array := make([]int, s)

	for i := 0; i < s; i++ {
		array[i] = rand.Intn(10000) + 1
	}

	return array
}
