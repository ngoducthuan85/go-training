package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

var width = flag.Int("w", 256, "hash width (256 or 384 or 512)")

func sha(b []byte) []byte {
	switch *width {
	case 256:
		h := sha256.Sum256(b)
		return h[:]
	case 384:
		h := sha512.Sum384(b)
		return h[:]
	case 512:
		h := sha512.Sum512(b)
		return h[:]
	default:
		log.Fatal("Unexpected width specified.")
	}
	return []byte{}
}

func main() {
	flag.Parse()
	b := []byte(os.Args[1])
	fmt.Printf("%x\n", sha(b))
}
