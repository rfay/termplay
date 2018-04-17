package main

import (
	"fmt"
	"os"
	"golang.org/x/crypto/ssh/terminal"
)


func main() {


	isTerm := terminal.IsTerminal(1)
	fmt.Fprintf(os.Stdout, "isTerm=%v", isTerm)
	//fmt.Fprintf(os.Stdout, "\x1b[33mhmm\x1b0m\n")

}
