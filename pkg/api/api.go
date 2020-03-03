// +build api

package api

import (
	"fmt"
	//elastic7 "github.com/elastic/go-elasticsearch/v7"
	//"github.com/elastic/go-elasticsearch/v7/esapi"
	//"github.com/geoffmore/esctl-go/esutil"
	//"net/http"
	"context"
)

func CmdTest() {
	fmt.Println("Test")
}

//func IntegrationTest(esClient *elastic7.Client) error {
//	req := esapi.ClusterHealthRequest{
//		Pretty: true,
//		Human:  true,
//	}
//
//	err := esutil.Request(req, esClient)
//	return err
//}

//func buildRequest(esClient *elastic7.Client, method string, endpoint string) (req *http.Request, err error) {
//	// https://pkg.go.dev/net/http?tab=doc#Request
//
//	// Build complete URL
//	// Validate and add method
//	// Add optional body? This needs to be supported by a flag. Maybe -d?
//	// var req http.Request
//
//	// I should probably use IsError() syntax eventually
//	req, err = http.NewRequest(method, url)
//
//	// Add auth?
//
//	// Validate and add URL from esClient
//
//	return req, err
//}

// Do I pass this up? For now, I say yes
//func makeRequest(esClient *elastic7.Client, req *http.Request) (response *http.Response, err error) {
//	//	// func (c *Client) Perform(req *http.Request) (*http.Response, error)
//	//	res, err = esClient.Perform(req)
//	//	return response, err
//}
func Do(esClient *elastic7.Client, method string, endpoint string) error {
	// Assumes absolutely no errors
	req, err = buildRequest(esClient, method, endpoint)
	res, err = makeRequest(esClient, req)

	// I need:
	//type esRequest interface {
	//	Do(context.Context, esapi.Transport) (*esapi.Response, error)
	//}
	// I should find out how to use esutil, but that has its complexities. I
	// think it boils down to satisfying this interface: https://godoc.org/github.com/elastic/go-elasticsearch/esapi#Transport

	// Transport interface -> esRequest interface
}

type apiFooTransport struct {
}

// This is now a esapi.Transport
func (a apiFooTransport) Perform(*http.Request) (*http.Response, error) {
}

type apiFooEsRequest struct{}

// This is now esRequest
func (a apiFooEsRequest) Do(context.Context, esapi.Transport) (*esapi.Response, error)

// This is esapi.Response
// type Response struct {
//    StatusCode int
//    Header     http.Header
//    Body       io.ReadCloser
//}
// What if each subcommand of the api subcommand was a http verb?
// I don't like that idea. Too much work
