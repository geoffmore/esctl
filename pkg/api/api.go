// +build api

package api

import (
	"context"
	"errors"
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	//"github.com/elastic/go-elasticsearch/v7/estransport"
	"github.com/geoffmore/esctl-go/esutil"
	//"math/rand"
	"net/http"
	"reflect"
	"strings"
	//	"time"
)

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods
var httpRequestMethods = []string{
	"GET",
	"HEAD",
	"POST",
	"PUT",
	"DELETE",
	"CONNECT",
	"OPTIONS",
	"TRACE",
	"PATCH",
}

//func CmdTest() {
//	fmt.Println("Test")
//}

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
//func Do(esClient *elastic7.Client, method string, endpoint string) error {
//	// Assumes absolutely no errors
//	req, err = buildRequest(esClient, method, endpoint)
//	res, err = makeRequest(esClient, req)
//
//	// I need:
//	//type esRequest interface {
//	//	Do(context.Context, esapi.Transport) (*esapi.Response, error)
//	//}
//	// I should find out how to use esutil, but that has its complexities. I
//	// think it boils down to satisfying this interface: https://godoc.org/github.com/elastic/go-elasticsearch/esapi#Transport
//
//	// Transport interface -> esRequest interface
//}

//type apiFooTransport struct {
//	// Needs some way to get method, url
//	// Also needs a buildURL function
//	// Maybe start with config, select a random url, append the endpoint?
//
//	// Adding a *bufio.Reader to the transport because I think that's what it's
//	// missing
//	// Should This be io.Reader or *bufio.Reader?
//	Reader *bufio.Reader
//}

// type io.Reader interface {
//    Read(p []byte) (n int, err error)
//}

// bufio.NewReader(io.Reader) *bufio.Reader

//func makeAPIFooTransport() apiFooTransport {
//	// bufio.NewReader(io.Reader) *bufio.Reader
//
//	// What if es.Client already has a buffered transport?
//	// Should I be using estransport? It has Perform
//	return apiFooTransport{
//		Reader: bufio.NewReader(reader),
//	}
//
//}

// This is now a esapi.Transport
//func (a apiFooTransport) Perform(req *http.Request) (*http.Response, error) {
//	// Need to take method and url from the transport
//	// Need to move that around I should have a request already
//	//req, err := http.NewRequest(a.method, a.buildURL(), nil)
//	// Technically, I only need to care about StatusCode, Header, and Body
//
//	//func ReadResponse(r *bufio.Reader, req *Request) (*Response, error)
//	res, err := http.ReadResponse(a.Reader, req)
//	if err != nil {
//		return nil, err
//	}
//
//	return res, nil
//}

//func (a apiFooTransport) buildURL() string {}

// Wrapper around http.Request
type apiFooEsRequest struct {
	Request *http.Request
	//Path    strings.Builder
}

// This is now esRequest
func (a apiFooEsRequest) Do(ctx context.Context, transport esapi.Transport) (*esapi.Response, error) {

	// Use the request to get a response
	// Maybe

	// Start with http response
	//res
	// esapi.Response is a http.Response with methods :

	// func (r *Response) IsError() bool
	// func (r *Response) Status() string
	// func (r *Response) String() string
	// Actually, it just has a subset of http.Response fields
	// And it is a struct, not an interface

	// context should be added to req, then transport.Perform(req) should be
	// called

	// Ignoring ctx for now

	res, err := transport.Perform(a.Request)
	if err != nil {
		return nil, err
	}

	esRes := esapi.Response{
		StatusCode: res.StatusCode,
		Header:     res.Header,
		Body:       res.Body,
	}

	// How do I avoid a nil pointer here?
	return &esRes, nil

	// Below is an example of the Do method. I got pretty close
	// https://github.com/elastic/go-elasticsearch/blob/52d1cf7160ac9b92b81ac6c82acec3f20351d8e7/esapi/api.bulk.go#L64

}

//func genRequest() esutil.esRequest

// This is esapi.Response
// type Response struct {
//    StatusCode int
//    Header     http.Header
//    Body       io.ReadCloser
//}

// What if each subcommand of the api subcommand was a http verb?
// I don't like that idea. Too much work

//func Foo(esClient *elastic7.Client, outputFmt string, method string, endpoint string) error {
//
//	// Make an esRequest
//	req := genRequest().(esutil.esRequest)
//	//req := esapi.CatHelpRequest{
//	//	Human:  true,
//	//	Pretty: true,
//
//	//	// This may be a v7 vs v8 issue. Unlocking v7
//	//	Help: &help,
//	//}
//
//	// Boilerplate
//	//changedField := esutil.SetFormat(reflect.ValueOf(&req).Elem(), outputFmt)
//	// // Make a request to get bytes
//	b, err := esutil.RequestNew(req, esClient)
//	if err != nil {
//		return err
//	}
//	// // Print bytes
//	err = esutil.ParseBytes(b, changedField, outputFmt)
//	return err
//}

func newRequest(esClient *elastic7.Client, method, endpoint string) (*http.Request, error) {
	var (
		//method string
		path strings.Builder
	)

	// Validate method
	// https://golangcode.com/check-if-element-exists-in-slice/
	_, found := Find(httpRequestMethods, strings.ToUpper(method))
	// This is incorrect, and I need to generate an error
	if !found {
		return nil, errors.New("InvalidMethodError")
	}

	// Seed randomness
	// https://stackoverflow.com/questions/33994677
	//s := rand.NewSource(time.Now().Unix())
	//r := rand.New(s)

	// Find the stored URLs in esClient
	//urls := esClient.Transport.URLs()
	//esTransportInterface := esClient.Transport.(estransport.Interface)
	// Can I use reflection here?
	//esTransportClient := esTransPortInterface.Perform()
	//tpClient := &esTransport.(estransport.Client)
	//tpClient := esClient.Transport.(estransport.Client)
	//urls := esTransport.URLs()
	//urls := esClient.Transport.(estransport.Client).URLs()
	//urls := []string{"foo", "bar"}
	// Select a random address as url
	//urlIndex := r.Intn(len(urls))
	//url := urls[urlIndex]

	// Expand path for url
	//path.Grow(1 + len(url))
	//path.WriteString(url)
	// Expand path for endpoint
	path.Grow(1 + len("/"+endpoint))
	path.WriteString("/" + endpoint)
	//method, path.String()

	//if err != nil {
	//	return nil, err
	//}

	// I don't need a url since estransport.Client.Perform() modifies an existing
	// URL
	return http.NewRequest(strings.ToUpper(method), path.String(), nil)
}
func Foo1(esClient *elastic7.Client, outputFmt, method, endpoint string) error {
	// Create estransport.Client from elasticserch.Client verify that this is
	// needed
	//transportClient := esClient.(estransport.Client)
	//req, err := http.NewRequest(a.method, a.buildURL(), nil)
	//res, err := transportClient.Perform(req)

	httpReq, err := newRequest(esClient, method, endpoint)
	if err != nil {
		return err
	}

	req := apiFooEsRequest{httpReq}
	// Build a http.Request
	//req := http.Request{
	//	Method: parseMethod(method),
	//	URL:    randURL(esClient),
	//}
	// I literally only need to build a request and pass that

	//type apiFooEsRequest struct {

	// Boilerplate
	changedField := esutil.SetFormat(reflect.ValueOf(&req).Elem(), outputFmt)
	// // Make a request to get bytes
	b, err := esutil.RequestNew(req, esClient)
	if err != nil {
		return err
	}
	// // Print bytes
	err = esutil.ParseBytes(b, changedField, outputFmt)
	return err
}

// https://golangcode.com/check-if-element-exists-in-slice/
// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
