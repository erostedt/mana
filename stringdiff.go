package main

func Min(elements ...int) int {
    if (len(elements) == 0) {
        panic("Min called with 0 elements")
    }

    min := elements[0]
    for _, element := range elements {
        if (element < min) {
            min = element
        }
    }
    return min
}

func EditDistance(word1 []rune, word2 []rune) int {
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

			dp[i*n+j] = Min(dp[(i-1)*n+j]+1, // deletion
				dp[i*n+j-1]+1,                    // insertion
				dp[(i-1)*n+j-1]+substitutionCost) // substitution

		}
	}

	return dp[m*n-1]
}

type Match struct {
	word     string
	distance int
}
