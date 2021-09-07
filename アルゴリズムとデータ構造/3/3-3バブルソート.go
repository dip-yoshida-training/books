package main

import (
	"fmt"
)

func main() {
	var N, i, counta int
	var A [100]int
	var sortA []int

	fmt.Scanf("%d", &N)
	for i = 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}
	
	sortA, counta = bubbleSort(A[:], N)

	// 出力
	for i=0; i < N; i++ {
		fmt.Printf("%d ", sortA[i])
	}
	fmt.Printf("\n%d", counta)

}

// flagを用いたバブルソート
func bubbleSort(A []int, N int) ([]int, int) {
	var i, j int
	var counta int
	flag := true

	for i = 0; flag; i++ {
		flag = false
		for j = N -1; j > i; j-- {
			// 隣接同士を比較し要素を交換
			if A[j] < A[j - 1] {
				A[j], A[j - 1] = A[j - 1], A[j]
				counta++
				flag = true
			}
		}
	}
	return A[:], counta
}