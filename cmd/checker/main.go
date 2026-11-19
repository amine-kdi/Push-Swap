package main

import (
	"bufio"
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

	if len(nombres) == 0 {
		return
	}

	machine := pushswap.NewMachine(nombres)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instruction := strings.TrimSpace(scanner.Text())

		if instruction == "" {
			continue
		}

		err := machine.Exec(instruction)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if machine.IsSorted() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
