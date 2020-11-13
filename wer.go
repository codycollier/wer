// WER provides word error rate and related calculations
//
// The word error rate algorithm is implemented per the
// reference found here:
//   https://martin-thoma.com/word-error-rate-calculation/
//
package wer

// Return the minimum of three ints
func minTrio(a, b, c int) int {
	min := a
	if a > b {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

// Return word error rate and word accuracy for (reference, candidate)
func WER(reference, candidate []string) (int, float64) {

	lr := len(reference)
	lc := len(candidate)

	// initialization
	var D [][]int
	for i := 0; i <= lr; i++ {
		for j := 0; j <= lc; j++ {
			if i == 0 {
				D[0][j] = j
			} else if j == 0 {
				D[i][0] = i
			}
		}
	}

	// calculation
	for i := 0; i <= lr; i++ {
		for j := 0; j <= lc; j++ {
			if reference[i] == candidate[j] {
				D[i][j] = D[i-1][j-1]
			} else {
				sub := D[i-1][j-1] + 1
				ins := D[i][j-1] + 1
				del := D[i-1][j] + 1
				D[i][j] = minTrio(sub, ins, del)
			}
		}
	}

	wer := D[lr][lc]
	wacc := 1 - wer

	return wer, wacc
}
