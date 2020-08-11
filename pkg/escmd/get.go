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
func GetClusterInfo(esClient *elastic7.Client, cfgOpts *opts.CommandOptions) error {

	// Why are InfoRequest and Info() not documented in v7.8.0?
	req := esapi.InfoRequest{
		Human:  true,
		Pretty: true,
	}

	// Flag init
	outputFmt := cfgOpts.OutputFormat
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

// GET /_cluster/health
func GetClusterHealth(esClient *elastic7.Client, cfgOpts *opts.CommandOptions) error {
	req := esapi.ClusterHealthRequest{
		Human:  true,
		Pretty: true,
	}

	// Flag init
	outputFmt := cfgOpts.OutputFormat
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
