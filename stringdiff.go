package main

import "fmt"

func Min(a int, b int, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func levenstein(asciiSlice1 []byte, asciiSlice2 []byte) ([]int, int, int) {
	m := len(asciiSlice1)
	n := len(asciiSlice2)
	dp := make([]int, (m+1)*(n+1))

	dp[0] = 0
	for i := 1; i < m+1; i++ {
		dp[i*n] = i
	}

	for j := 1; j < n+1; j++ {
		dp[j] = j
	}

	substitutionCost := 0
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if asciiSlice1[i-1] == asciiSlice2[j-1] {
				substitutionCost = 0
			} else {
				substitutionCost = 1
			}

			dp[i*n+j] = Min(dp[(i-1)*n+j]+1, // deletion
				dp[i*n+j-1]+1,                    // insertion
				dp[(i-1)*n+j-1]+substitutionCost) // substitution

		}
	}

	return dp, m, n
}

func PrintDp(dp []int, m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", dp[i*n+j])
		}
		fmt.Println("")
	}
}
