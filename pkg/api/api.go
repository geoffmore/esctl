package api

import (
	"context"
	"errors"
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/geoffmore/esctl/pkg/esutil"
	"github.com/geoffmore/esctl/pkg/opts"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods
// Static list of valid HTTP request verbs
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

// Inspired by https://github.com/elastic/go-elasticsearch/blob/52d1cf7160ac9b92b81ac6c82acec3f20351d8e7/esapi/api.bulk.go#L64
// This is now an esRequest than can be used by esutil
// Has a similar form to most *Request structs under esapi
func (r GenericRequest) Do(ctx context.Context, transport esapi.Transport) (*esapi.Response, error) {
	var (
		params map[string]string
	)
	params = make(map[string]string)

	if r.Format != "" {
		params["format"] = r.Format
	}

	if len(r.H) > 0 {
		params["h"] = strings.Join(r.H, ",")
	}

	if r.Help != nil {
		params["help"] = strconv.FormatBool(*r.Help)
	}

	if r.Local != nil {
		params["local"] = strconv.FormatBool(*r.Local)
	}

	if len(r.S) > 0 {
		params["s"] = strings.Join(r.S, ",")
	}

	if r.V != nil {
		params["v"] = strconv.FormatBool(*r.V)
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req := r.Request

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	// TODO - This should probably try to detect different content types and convert them to json before sending
	if r.Body != nil {
		req.Header["Content-Type"] = []string{"application/json"}
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(r.Request)
	if err != nil {
		return nil, err
	}

	response := esapi.Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil

}

// inspired by esapi/esapi.request.go
func newRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	var (
		//method string
		path strings.Builder
	)

	// Validate method
	_, found := esutil.Find(httpRequestMethods, strings.ToUpper(method))
	// This is incorrect, and I need to generate an error
	if !found {
		return nil, errors.New("InvalidMethodError")
	}

	// Strip leading '/' if it exists
	if strings.HasPrefix(endpoint, "/") {
		endpoint = strings.TrimLeft(endpoint, "/")
	}

	// Expand path for endpoint
	path.Grow(1 + len("/"+endpoint))
	path.WriteString("/" + endpoint)

	// A full url is not needed since estransport.Client.Perform() modifies an existing
	// url
	return http.NewRequest(strings.ToUpper(method), path.String(), body)
}
func MakeGenericRequest(esClient *elastic7.Client, cmdOpts *opts.CommandOptions) error {

	method := cmdOpts.Args[0]
	endpoint := cmdOpts.Args[1]

	httpReq, err := newRequest(method, endpoint, cmdOpts.Body)
	if err != nil {
		return err
	}

	// Initialize empty struct
	req := GenericRequest{}

	// Add the necessary *http.Request
	req.Request = httpReq

	// Ease of use. Should probably be a flag
	req.Pretty = true
	req.Human = true

	// Boilerplate
	r := reflect.ValueOf(&req).Elem()
	// Bring flags to the Request struct
	// Automatic stuff
	changedFields := opts.SetAllCmdOpts(r, cmdOpts)
	// Manual stuff
	// Avoiding changing opts.SetAllCmdOpts here because handling interfaces will be a pain
	// For this, I _know_ the fields/types already, so reflection isn't as unsafe
	req.Body = cmdOpts.Body

	// // Make a request to get bytes
	b, err := esutil.RequestNew(req, esClient)
	if err != nil {
		return err
	}
	// // Print bytes
	err = esutil.ParseBytes(b, changedFields["Format"], cmdOpts.OutputFormat)
	return err
}
