package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	
    fmt.Scan(&n)

    R := make([]int, n)

    for i := range R {
        fmt.Scan(&R[i])
    }

	maxv := -200000
	minv := R[0]

	for i := 1; i < n; i++ { // R[i]呼び出す
		maxv = max(maxv, R[i] - minv) //最大値更新
		minv = min(minv, R[i]) // 最小値の保持
	}

	fmt.Println(maxv)
}

func min(nums ...int) int {
    if len(nums) == 0 {
        panic("funciton min() requires at least one argument.")
    }
    res := nums[0]
    for i := 0; i < len(nums); i++ {
        res = int(math.Min(float64(res), float64(nums[i])))
    }
    return res
}

func max(nums ...int) int {
    if len(nums) == 0 {
        panic("funciton max() requires at least one argument.")
    }
    res := nums[0]
    for i := 0; i < len(nums); i++ {
        res = int(math.Max(float64(res), float64(nums[i])))
    }
    return res
}
