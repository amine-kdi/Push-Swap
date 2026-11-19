package main

import (
	"fmt"
	"os"
	"strings"

	"push-swap/pushswap"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	texte := strings.Join(os.Args[1:], " ")

	nombres, err := pushswap.Parse(texte)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	for _, instruction := range pushswap.Sort(nombres) {
		fmt.Println(instruction)
	}
}
