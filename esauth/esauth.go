package esauth

// Use a config, return a client

import (
	"log"
)

import (
	elastic7 "github.com/elastic/go-elasticsearch/v7"
)

func EsAuth() *elastic7.Client {
	es, err := elastic7.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return es
}

//// Future work to satisfy esctl login
//func esLogin() {}
