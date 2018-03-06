package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int) //counts of Unicode characters
	rangeCounts := make(map[*unicode.RangeTable]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune,
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsDigit(r) {
			rangeCounts[unicode.Number]++
		}
		if unicode.IsLetter(r) {
			rangeCounts[unicode.Letter]++
		}
		if unicode.IsSpace(r) {
			rangeCounts[unicode.Space]++
		}
		if unicode.IsPunct(r) {
			rangeCounts[unicode.Punct]++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	fmt.Print("\nRange Table Count\n")
	fmt.Printf("number: %d\n", rangeCounts[unicode.Number])
	fmt.Printf("letter: %d\n", rangeCounts[unicode.Letter])
	fmt.Printf("space: %d\n", rangeCounts[unicode.Space])
	fmt.Printf("punct: %d\n", rangeCounts[unicode.Punct])
}
