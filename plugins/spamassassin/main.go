package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func testEmailAPI(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	res := spamassassinCheck(body)
	fmt.Fprintf(w, res)
}

func spamassassinCheck(m []byte) string {
	cmd := fmt.Sprintf("echo %s | spamassassin", m)
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		log.Printf("unable to use spamassassin: %v", err)
	}
	return string(output)
}

func main() {
	http.HandleFunc("/check", testEmailAPI)
	log.Fatal(http.ListenAndServe(":12800", nil))
}
