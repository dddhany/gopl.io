package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Printf("different bit=%v", countNumberOfDifferentBitSha256("x", "X"))
}

func countNumberOfDifferentBitSha256(s1, s2 string) int {
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))
	differentBit := 0
	for i := range c1 {
		for xorByte := c1[i] ^ c2[i]; xorByte > 0; xorByte = xorByte >> 1 {
			if xorByte&1 == 1 {
				differentBit++
			}
		}
	}
	return differentBit
}
