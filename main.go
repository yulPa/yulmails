package main

import (
	"flag"

	"github.com/yulPa/yulmails/api"
)

var (
	certFile = flag.String("tls-crt-file", "domain.tld.crt", "A certificate file")
	keyFile  = flag.String("tls-key-file", "domain.tld.key", "A key file")
)

func main() {

	flag.Parse()
	api.Start(*certFile, *keyFile)
}
