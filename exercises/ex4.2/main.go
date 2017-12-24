// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

//!+
import "crypto/sha256"
import "crypto/sha512"
import "io/ioutil"
import "os"
import "flag"
import "fmt"

func main() {
	var size int
	flag.IntVar(&size, "size", 256, "Checksum size. Valid options: 256, 384 and 512")

	flag.Parse()

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	switch size {
	case 384:
		fmt.Printf("SHA384: %x\n", sha512.Sum384(in))
	case 512:
		fmt.Printf("SHA512: %x\n", sha512.Sum512(in))
	default:
		fmt.Printf("SHA256: %x\n", sha256.Sum256(in))
	}
}

//!-
