package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Original version
func PopCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// With loop
func PopCount2(x uint64) int {
	result := 0
	for i := uint(0); i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

// With removing each bit one by one
func PopCount3(x uint64) int {
	result := 0
	for x > 0 {
		x = x & (x - 1)
		result++
	}
	return result
}
