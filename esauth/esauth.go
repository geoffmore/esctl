package esauth

import (
	elastic7 "github.com/elastic/go-elasticsearch/v7"
)

func EsAuth(cfg elastic7.Config) (client *elastic7.Client, err error) {
	client, err = elastic7.NewClient(cfg)
	return client, err
}
