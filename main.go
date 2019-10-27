package main

import (
	"log"
	"os"

	"github.com/yulpa/yulmails/cmd"
)

func main() {
	err := cmd.App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
