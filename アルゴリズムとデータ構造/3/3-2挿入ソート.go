package main

import (
	"fmt"
)

func main() {
	var N, i int
	var A [100]int

	fmt.Scanf("%d", &N)
	for i = 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}

	// arrayをスライスに変換して渡す
	trace(A[:], N)
	insertionSort(A[:], N)
}

// 配列の要素を順に出力
func trace(A []int, N int){
	var i int
	for i = 0; i < N; i++ {
		if i > 0 {
			fmt.Print(" ") // 隣接する要素の間に1つの空白を出力
		}
		fmt.Printf("%d", A[i])
	}
	fmt.Print("\n");
}

// 挿入ソート
func insertionSort(A []int, N int){
	var j, i, v int
	for i = 1; i < N; i++ {
		v = A[i]
		j = i - 1
		for j >= 0 && A[j] > v {
			A[j + 1] = A[j]
			j--
		}
		A[j + 1] = v;
		trace(A, N)
	}
}