package entrypoint

import (
	"bufio"
	"log"
	"os"
)

func Run() {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		log.Print(reader.Text())
	}

}
