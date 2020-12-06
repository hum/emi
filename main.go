package main

import (
	"fmt"
	"os"

	"github.com/hum/emi/repl"
)

func main() {
	fmt.Printf("A REPL session of the Emi programming language.\n")
	repl.Start(os.Stdin, os.Stdout)
}
