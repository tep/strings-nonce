package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"

	"toolman.org/base/toolman/v2"
	"toolman.org/strings/nonce"
)

var (
	length  = pflag.UintP("length", "l", 32, "Length of generated string")
	verbose = pflag.BoolP("verbose", "V", false, "Be verbose")
)

func main() {
	toolman.InitCLI()

	n, err := nonce.New(int(*length))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	if *verbose {
		fmt.Printf("Nonce: %q (%d)\n", n, len(n))
		return
	}

	fmt.Println(n)
}
