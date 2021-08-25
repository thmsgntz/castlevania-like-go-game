package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func IntToArray(x int) (arr []int) {
	xs := strings.Split(fmt.Sprint(x), "")
	// fmt.Println(xs, len(xs))

	arr = make([]int, 0, len(xs))
	for i := 0; i < len(xs); i++ {
		j, _ := strconv.Atoi(xs[i])
		arr = append(arr, j)
	}
	return arr
}
