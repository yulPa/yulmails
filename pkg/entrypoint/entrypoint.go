package entrypoint

import (
	"log"
	"net/http"
	"net/mail"

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
