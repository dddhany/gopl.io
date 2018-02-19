package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	var buf bytes.Buffer
	for i := 0; i < n; {
		if buf.Len() > 0 {
			buf.WriteByte(',')
			buf.WriteString(s[i : i+3])
			i += 3
		} else {
			firstSegmentLen := n % 3
			if firstSegmentLen == 0 {
				firstSegmentLen = 3
			}
			buf.WriteString(s[0:firstSegmentLen])
			i += firstSegmentLen
		}
	}
	return buf.String()
}

func main() {
	fmt.Printf("%s \n", comma("123"))
	fmt.Printf("%s \n", comma("1234"))
	fmt.Printf("%s \n", comma("12345"))
	fmt.Printf("%s \n", comma("123456"))
	fmt.Printf("%s \n", comma("1234567"))
}
