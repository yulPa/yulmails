package entrypoint

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"os/exec"
	"strings"

	"github.com/yulPa/yulmails/pkg/environment"
	"github.com/yulPa/yulmails/pkg/mongo"
)

// Run method
func Run() {
	http.HandleFunc("/mail", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			email, err := mail.ReadMessage(req.Body)
			if err != nil {
				log.Printf("unable to read message: %v", err)
			}

			log.Print(checkIfAuthorize(email.Header.Get("from")))

			raw, _ := ioutil.ReadAll(req.Body)
			checkSpamAssassin(string(raw))
		}
	})
	http.ListenAndServe(":8080", nil)
}

func checkIfAuthorize(from string) bool {
	var archivingdb = mongo.NewSession("mongodb://archivingdb:27017")

	db := archivingdb.DB("configuration")

	env, err := db.ReadEnvironments("yulpa")
	if err != nil {
		log.Printf("unable to read environments: %v", err)
	}

	return contains(from, env)

}

func contains(s string, a []environment.Environment) bool {
	for _, env := range a {
		for _, el := range env.IPs {
			if s == el {
				return true
			}
		}
	}
	return false
}

func checkSpamAssassin(content string) {
	/*
	   This method will call SA in order to determine if mails are spams
	*/
	subProcess := exec.Command("spamassassin")
	subProcess.Stdin = strings.NewReader(content)

	raw, err := subProcess.Output()
	log.Print(string(raw))
	if err != nil {
		log.Fatalf("unable to get output: %v", err)
	}

	r := strings.NewReader(string(raw))
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatalf("unable to read message: %v", err)
	}
	if strings.Split(m.Header.Get("X-Spam-Status"), ",")[0] == "Yes" {
		log.Println("Spam")
		log.Print(m.Header.Get("X-Spam-Status"))
	} else {
		log.Println("not spam")
	}
}
