package escmd

import (
	"fmt"
	elastic7 "github.com/geoffmore/go-elasticsearch/v7"
	"github.com/geoffmore/go-elasticsearch/v7/esapi"
	"github.com/geoffmore/esctl-go/esutil"
	"reflect"
)

// All methods I've tested other than es7.Info() are not found outside of
// esapi. In addition, based on the documentation I read on the
// documentation, I would assume I could pass a elasticsearch.Client as my
// transport, but I need to use the Transport type within the object instead

// GET /_cluster/info
func GetClusterInfo(esClient *elastic7.Client) error {
	resp, err := esClient.Info()
	if err != nil {
		return err
	}
	rhee, err := esutil.Des(resp)
	if err != nil {
		return err
	}
	esutil.MapToYamlish(rhee, 0)
	fmt.Println()

	return nil
}

// GET /_cluster/health
func GetClusterHealth(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.ClusterHealthRequest{
		Human:  true,
		Pretty: true,
	}

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
