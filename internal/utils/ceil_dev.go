package utils

import (
	"math"
)

// 向上取整
// fmt.Println(CeilDiv(10, 3)) // 4
// fmt.Println(CeilDiv(7, 2))  // 4
// fmt.Println(CeilDiv(5, 5))  // 1
// fmt.Println(CeilDiv(1, 2))  // 1
func CeilDev(a, b int) int {
	return int(math.Ceil(float64(a) / float64(b)))
}
