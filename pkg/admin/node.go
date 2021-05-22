package admin

import (
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/geoffmore/esctl/pkg/esutil"
	"github.com/geoffmore/esctl/pkg/opts"
	"reflect"
)

// What does this line do in func list_nodes?
// 'dnodes=$(echo "${output}" | awk '/data|di.*instance/ { print $10 }' | sed 's/.*-00*//' | sort | paste -s -d"," -)'
// Looks like this looks for 'd' in node.role. I guess that means I need to
// transform some json
// But how do I get json _and_ the desired output format in the same request?
// I think I could either make the same request or parse the output based on
// user provided outputFmt flag
// Opting to duplicate the request - Once for w/ user specified output format
// and once for the extra stuff
// I don't like that last output, so I'm going to leave it off. Too much work to
// add that output with not enough value

func NodeList(esClient *elastic7.Client, cmdOpts *opts.CommandOptions) error {

	var fieldArray []string = []string{"ip", "heap.percent", "ram.percent", "cpu", "load_1m", "load_5m", "load_15m", "node.role", "master", "name", "disk.total", "disk.used", "disk.avail", "disk.used_percent"}
	var sortArray []string = []string{"name:asc"}

	req := esapi.CatNodesRequest{
		Pretty: true,
		Human:  true,

		H: fieldArray,
		S: sortArray,

		V: &cmdOpts.Verbose,
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

func ListNodes(esClient *elastic7.Client, outputFmt string, help bool) error {

	var fieldArray []string = []string{"ip", "heap.percent", "ram.percent", "cpu", "load_1m", "load_5m", "load_15m", "node.role", "master", "name", "disk.total", "disk.used", "disk.avail", "disk.used_percent"}
	var sortArray []string = []string{"name:asc"}

	var verbose bool = true

	req := esapi.CatNodesRequest{
		Human:  true,
		Pretty: true,

		H: fieldArray,
		S: sortArray,

		// This may be a v7 vs v8 issue. Unlocking v7
		Help: &help,
		V:    &verbose,
	}

	// Kept for future inspiration

	// // Print bytes
	// Get parsable json output
	// Need a MarshalBytes function that returns a map[string]interface{} for
	// further processing
	// esutil.MarshalBytes()

	// Boilerplate
	changedField := esutil.SetFormat(reflect.ValueOf(&req).Elem(), outputFmt)
	// // Make a request to get bytes
	b, err := esutil.RequestNew(req, esClient)
	if err != nil {
		return err
	}

	err = esutil.ParseBytes(b, changedField, outputFmt)
	return err

}

func ListNodesStorage(esClient *elastic7.Client, outputFmt string, help bool) error {

	var fieldArray []string = []string{"ip", "node.role", "master", "name", "disk.total", "disk.used", "disk.avail", "disk.used_percent"}
	var sortArray []string = []string{"disk.used_percent:desc"}

	var verbose bool = true

	req := esapi.CatNodesRequest{
		Human:  true,
		Pretty: true,

		H: fieldArray,
		S: sortArray,

		// This may be a v7 vs v8 issue. Unlocking v7
		Help: &help,
		V:    &verbose,
	}

	// Boilerplate
	changedField := esutil.SetFormat(reflect.ValueOf(&req).Elem(), outputFmt)
	// // Make a request to get bytes
	b, err := esutil.RequestNew(req, esClient)
	if err != nil {
		return err
	}

	err = esutil.ParseBytes(b, changedField, outputFmt)
	return err

}

// Skipping list_nodes_zenoss_alarms because dependency on zencli cannot be
// satisfied in this package
