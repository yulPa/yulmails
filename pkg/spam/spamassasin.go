package main

import (
  "log"
  "os/exec"
  "bytes"
)

func init()  {
  /*
    This init method is called each time this package is loaded
    We check if `spamassassin` is installed on the machine
  */
  cmd := exec.Command("which", "spamassassin")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal("`spamassassin` is not installed on this machine: ", err)
	}
}

func main()  {
  log.Printf("Hello, World")
}
