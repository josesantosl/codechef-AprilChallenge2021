package main

import (
	"fmt"
	"math"
)

func Worthy(A [][]float64, i0, j0, k int) int {
	var somma float64 = 0
	dict := make(map[float64]float64)
	for i := i0; i < i0+k; i++ {
		for j := j0; j < j0+k; j++ {
			somma += A[i][j]
			dict[A[i][j]] = dict[A[i][j]] + 1
		}

	}
	avg := somma / float64(k*k)
	//fmt.Println(dict)                      //borrar eventualmente.
	//fmt.Println(avg, ":", dict[avg], "\n") //borrar eventualmente.

	var answer float64
	if avg >= float64(k*k) {
		answer = dict[avg]
		//fmt.Println("Valid") //borrar eventualmente.
	} else {
		//fmt.Println("Not Valid") //borrar eventualmente.
		answer = 0
	}
	return int(answer)
}

func main() {
	var tt, N, M, k, counter int
	fmt.Scan(&tt)
	for t := 0; t < tt; t++ {
		counter = 0
		fmt.Scan(&N)
		fmt.Scan(&M)
		fmt.Scan(&k)
		k = int(math.Sqrt(float64(k)))
		A := make([][]float64, N)
		for i := 0; i < N; i++ {
			A[i] = make([]float64, M)
			for j := 0; j < M; j++ {
				fmt.Scan(&A[i][j])
				if A[i][j] >= float64(k*k) {
					counter++
				}
			}
		}
		for i := 0; i <= N-k; i++ {
			for j := 0; j <= M-k; j++ {
				counter += Worthy(A, i, j, k)
			}

		}
		fmt.Println(counter)

	}
}
