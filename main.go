package main

import (
	"fmt"
	"os"
)

func main() {
	e, err := NewEnsemble(".")
	if err != nil {
		fmt.Printf("Failed to initiate ensemble: %s\n", err)
		os.Exit(2)
	}
	if len(os.Args) < 2 {
		fmt.Printf("Provide commands to ensemble: start, next")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "start":
		if len(os.Args) < 3 {
			fmt.Println("Branch required")
			os.Exit(1)
		}
		e.Start(os.Args[2])
	case "next":
		e.Next()
	}
}
