package main

import (
	"log"
	"os"

	"github.com/benwebber/invadicon"
)

func main() {
	var seed string
	if len(os.Args) == 1 {
		seed = ""
	} else {
		seed = os.Args[1]
	}
	i, err := invadicon.New(seed)
	if err != nil {
		log.Fatal(err)
	}
	i.Write(os.Stdout)
}
