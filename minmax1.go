package main

import "fmt"

type game struct {
	nama     string
	populasi int
	nilai    float64
}

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	var T [20]game

	for i := 0; i < N; i++ {
		fmt.Scan(&T[i].nama)
	}

	for i := 0; i < N; i++ {
		T[i].populasi = K / 10
		K = K - T[i].populasi

		digitTerakhir := T[i].populasi % 1000
		T[i].nilai = float64(digitTerakhir) * 0.025

		if T[i].nilai > 5.0 {
			fmt.Printf("%s dengan tingkat kejahatan: %.2f%%\n", T[i].nama, T[i].nilai)
		}
	}
}