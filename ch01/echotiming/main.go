package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", " "
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(s)
	fmt.Printf("slow: %dns\n", time.Since(start).Nanoseconds())

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("fast: %dns\n", time.Since(start).Nanoseconds())
}
