package escmd

import (
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/geoffmore/esctl/pkg/esutil"
	"github.com/geoffmore/esctl/pkg/opts"
	"reflect"
)

// GET /_cluster/info
// This endpoint is special because it doesn't have a <foo>Request struct
func GetClusterInfo(esClient *elastic7.Client, cmdOpts *opts.CommandOptions) error {

	// Why are InfoRequest and Info() not documented in v7.8.0?
	req := esapi.InfoRequest{
		Human:  true,
		Pretty: true,
	}

	// Boilerplate
	r := reflect.ValueOf(&req).Elem()
	// Bring flags to the Request struct
	changedFields := esutil.SetAllCmdOpts(r, cmdOpts)
	// // Make a request to get bytes
	b, err := esutil.RequestNew(req, esClient)
	if err != nil {
		return err
	}
	// // Print bytes
	err = esutil.ParseBytes(b, changedFields["Format"], cmdOpts.OutputFormat)
	return err
}

// GET /_cluster/health
func GetClusterHealth(esClient *elastic7.Client, cmdOpts *opts.CommandOptions) error {
	req := esapi.ClusterHealthRequest{
		Human:  true,
		Pretty: true,
	}

	// Boilerplate
	r := reflect.ValueOf(&req).Elem()
	// Bring flags to the Request struct
	changedFields := esutil.SetAllCmdOpts(r, cmdOpts)
	// // Make a request to get bytes
	b, err := esutil.RequestNew(req, esClient)
	if err != nil {
		return err
	}
	// // Print bytes
	err = esutil.ParseBytes(b, changedFields["Format"], cmdOpts.OutputFormat)
	return err
}
