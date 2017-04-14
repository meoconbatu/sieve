package sieve

import "math"

const testVersion = 1

func Sieve(input int) (p []int) {
	sqrtInput := int(math.Sqrt(float64(input)))
	a := make([]bool, input)
	for i := 2; i < input; i++ {
		if a[i] == false {
			p = append(p, i)
			if i <= sqrtInput {
				for j := i; j < input; j = j + i {
					a[j] = true
				}
			}
		}
	}
	return
}
