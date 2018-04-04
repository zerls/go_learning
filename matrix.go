package main

func matrixNormal(C [][]int, B [][]int, A [][]int, N int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
}
