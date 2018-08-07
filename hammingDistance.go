package main

// HammingDistance calculates the hamming distance btw two integers
func HammingDistance(a, b int) int {
	result := a ^ b
	count := 0
	for ; result > 0; result >>= 1 {
		if result&1 > 0 {
			count++
		}
	}
	return count
}
