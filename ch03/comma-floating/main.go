package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer
	sPart := s
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		buf.WriteString(s[0:1])
		sPart = s[1:]
	}
	floatingPointS := ""
	if dotPos := strings.Index(sPart, "."); dotPos >= 0 {
		floatingPointS = sPart[dotPos:]
		sPart = sPart[0:dotPos]
	}
	buf.WriteString(commaRecursive(sPart))
	buf.WriteString(floatingPointS)
	return buf.String()
}

func commaRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaRecursive(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Printf("%s \n", comma("1234.567"))
	fmt.Printf("%s \n", comma("1234.567890"))
	fmt.Printf("%s \n", comma("-1234.5678901"))
	fmt.Printf("%s \n", comma("+1.5"))
}
