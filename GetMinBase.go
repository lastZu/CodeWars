package main

import "math"

func GetMinBase(n uint64) uint64 {
	var i uint64
	for i = 2; i*i*i < n; i++ {
		if () {
			
		}
	}

	base := n - 1
	for i := 2; uint64(math.Pow(float64(n), float64(1.0/i))) < n; i++ {
		temp := uint64(math.Pow(float64(n), float64(1.0/i)))

	}


	return base
}
