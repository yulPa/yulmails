package entrypoint

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	reader := bufio.NewReader(os.Stdin)
	emails, err := reader.ReadString('\n')
	fmt.Println(emails, err)
}
