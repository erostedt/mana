package mana

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

func MakeMatrix[T any](rows, cols int) [][]T {
	matrix := make([][]T, rows)
    for i := range matrix {
        matrix[i] = make([]T, cols)
    }
    return matrix
}


func EditDistance(word1 []rune, word2 []rune) int {
	rows := len(word1)
	cols := len(word2)

    if rows == 0 {
        return cols
    }
    if cols == 0 {
        return rows
    }

    dp := MakeMatrix[int](rows+1, cols+1)
	for r := 0; r <= rows; r++ {
		dp[r][0] = r
	}

	for c := 0; c <= cols; c++ {
		dp[0][c] = c
	}

	for r := 1; r <= rows; r++ {
		for c := 1; c <= cols; c++ {
			if word1[r-1] == word2[c-1] {
                dp[r][c] = dp[r-1][c-1]
			} else {
                dp[r][c] = 1 + Min(
                    dp[r-1][c],   // deletion
                    dp[r][c-1],   // insertion
                    dp[r-1][c-1]) // substitution
			}


		}
	}

	return dp[rows][cols]
}

type Match struct {
	word     string
	distance int
}
