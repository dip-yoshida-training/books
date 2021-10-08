package main

import (
	"fmt"
)

type P struct {
	name string
	t    int
}

const LEN = 10005

var Q [LEN]P
var head, tail, n int

func enqueue(x P) {
	Q[tail] = x
	tail = (tail + 1) % LEN
}

func dequeue() P {
	x := Q[head]
	head = (head + 1) % LEN
	return x
}

// 最小値の返却
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	elaps := 0
	var c, i, q int
	var u P
	fmt.Scanf("%d %d", &n, &q)

	// 全てのプロセスをキューに順番に追加
	for i = 1; i <= n; i++ {
		fmt.Scanf("%s", &Q[i].name)
		fmt.Scanf("%d", &Q[i].t)
	}
	head = 1
	tail = n + 1

	for {
		if head == tail {
			break
		}
		u = dequeue()
		// q または必要な時間 u.t だけ処理を行う
		c = min(q, u.t)
		// 残りの必要時間を計算
		u.t -= c
		 // 経過時間を加算
		elaps += c

		// 処理が完了していなければキューに追加
		if u.t > 0 {
			enqueue(u)
		} else {
			fmt.Printf("%s %d\n", u.name, elaps)
		}
	}
}
