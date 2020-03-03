// +build api

package api

import (
	"fmt"
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/geoffmore/esctl-go/esutil"
	"net/http"
)

func CmdTest() {
	fmt.Println("Test")
}

func IntegrationTest(esClient *elastic7.Client) error {
	req := esapi.ClusterHealthRequest{
		Pretty: true,
		Human:  true,
	}

	err := esutil.Request(req, esClient)
	return err
}

func buildRequest(esClient *elastic7.Client, method string, endpoint string) (req *http.Request, err error) {
	// https://pkg.go.dev/net/http?tab=doc#Request

	// Build complete URL
	// Validate and add method
	// Add optional body? This needs to be supported by a flag. Maybe -d?
	// var req http.Request

	// I should probably use IsError() syntax eventually
	req, err = http.NewRequest(method, url)

	// Add auth?

	// Validate and add URL from esClient

	return req, err
}
func makeRequest(esClient *elastic7.Client) (response *http.Response, err error) {}
