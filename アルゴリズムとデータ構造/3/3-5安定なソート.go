package main

import (
	"fmt"
	"strconv"
)

type Card struct {
	mark string
	value int
}

func main() {
	var N, i int
	var str[100]string
	var C1[100]Card
	var C2[100]Card

	fmt.Scanf("%d", &N)
	for i = 0; i < N; i++ {
		fmt.Scanf("%s", &str[i])
		C1[i].mark = str[i][0:1]
		C1[i].value, _ = strconv.Atoi(str[i][1:2])	
	}
	C2 = C1

	bubbleSort(C1[:], N)
	selectionSort(C2[:], N)
	
	for i = 0; i < N; i++ {
		fmt.Printf("%s%v ", C1[i].mark, C1[i].value)
	}
	fmt.Print("\nStable\n")

	for i = 0; i < N; i++ {
		fmt.Printf("%s%v ", C2[i].mark, C2[i].value)
	}
	if isStable(C1[:], C2[:], N) {
		fmt.Print("\nStable")
	}else{
		fmt.Print("\nNot Stable")
	}

}

// バブルソート
func bubbleSort(A[] Card, N int) {
	var i, j int
	for i = 0; i < N; i++ {
		for j = N - 1; j >= i + 1; j-- {
			if A[j].value < A[j - 1].value {
				A[j], A[j - 1] = A[j - 1], A[j]
			}
		}	
	}
}

// 選択ソート
func selectionSort(A[] Card, N int) {
	var i, j, minJ int
	for i = 0; i < N - 1; i++ {
		minJ = i;
		for j = i; j < N; j++ {
			if A[j].value < A[minJ].value {
				minJ = j
			}
		}
		A[i], A[minJ] = A[minJ], A[i]
	}
}

func isStable(C1[] Card, C2[] Card, N int) bool {
	var i int
	for i = 0; i < N; i++ {
		if C1[i].mark != C2[i].mark {
			return false
		}
	}
	return true
}