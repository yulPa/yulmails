package main

import (
	"log"
	"os"

	"gitlab.com/tortuemat/yulmails/cmd"

)

func main() {
	err := cmd.App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
