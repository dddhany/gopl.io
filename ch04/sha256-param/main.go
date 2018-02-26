package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var hash = flag.String("hash", "SHA256", "SHA256, SHA384 or SHA512")

func main() {
	flag.Parse()
	strarr := []byte(flag.Arg(0))
	switch *hash {
	case "SHA512":
		fmt.Printf("sha512(%s): %x\n", strarr, sha512.Sum512(strarr))
	case "SHA384":
		fmt.Printf("sha384(%s): %x\n", strarr, sha512.Sum384(strarr))
	default:
		fmt.Printf("sha256(%s): %x\n", strarr, sha256.Sum256(strarr))
	}
}
