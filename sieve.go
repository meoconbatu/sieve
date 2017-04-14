package sieve

import "math"

const testVersion = 1

func Sieve(input int) []int {
	sqrtInput := int(math.Sqrt(float64(input)))
	a := make([]bool, input)
	for i := 2; i <= sqrtInput; i++ {
		if a[i] == false {
			for j := i * i; j < input; j = j + i {
				a[j] = true
			}
		}
	}
	np := 0
	p := make([]int, input)
	for i := 2; i < input; i++ {
		if a[i] == false {
			p[np] = i
			np++
		}
	}
	return p[0:np]
}
