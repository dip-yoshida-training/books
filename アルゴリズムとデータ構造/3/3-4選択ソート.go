package main

import (
	"fmt"
)

func main() {
	var N, i, sw int
	var A[100]int

	fmt.Scanf("%d", &N);
	for i = 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}

	sw = selectionSort(A[:], N);

	for i = 0; i < N; i++ {
		if i > 0 {
			fmt.Print(" ");
		}
		fmt.Printf("%d", A[i])
	}
	fmt.Print("\n")
	fmt.Printf("%d", sw)
}

// 選択ソート
func selectionSort(A[] int, N int) int {
	var i, j, t, minJ int
	sw := 0 
	for i = 0; i < N - 1; i++ {
		minJ = i;
		for j = i; j < N; j++ {
			if A[j] < A[minJ] {
				minJ = j
			}
		}
		t = A[i]
		A[i] = A[minJ]
		A[minJ] = t;
		if i != minJ {
			sw++;
		}
	}
	return sw
}