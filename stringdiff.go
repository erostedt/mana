package main

func Min3(a int, b int, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func levenstein(word1 []rune, word2 []rune) int {
	m := len(word1)
	n := len(word2)
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
			if word1[i-1] == word2[j-1] {
				substitutionCost = 0
			} else {
				substitutionCost = 1
			}

			dp[i*n+j] = Min3(dp[(i-1)*n+j]+1, // deletion
				dp[i*n+j-1]+1,                    // insertion
				dp[(i-1)*n+j-1]+substitutionCost) // substitution

		}
	}

	return dp[m*n-1]
}

type Match struct {
	distance int
	word     string
}
