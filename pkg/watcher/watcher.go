package watcher

import (
	"context"
	"encoding/json"
	"fmt"
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/geoffmore/esctl/pkg/esutil"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"reflect"
)

// Watchers starting with dashes don't work currently because of cobra
// 'pods_pending' is valid
// '-pods_pending' is interpreted as having a flag and does not work

// Having a default value that can be empty causes weird behaviour sometimes

func WatcherPut(esClient *elastic7.Client, watch string, r io.Reader, initInactive bool, outputFmt string) error {

	var active bool
	// This flip isn't ideal, but necessary
	active = !initInactive

	req := esapi.WatcherPutWatchRequest{
		WatchID: watch,
		Body:    r,
		Active:  &active,

		Pretty: true,
		Human:  true,
	}

	// Boilerplate
	changedField := esutil.SetFormat(reflect.ValueOf(&req).Elem(), outputFmt)
	b, err := esutil.RequestNew(req, esClient)
	if err != nil {
		return err
	}
	// // Print bytes
	err = esutil.ParseBytes(b, changedField, outputFmt)
	return err
}

func WatcherGet(esClient *elastic7.Client, watch string, outputFmt string) error {
	req := esapi.WatcherGetWatchRequest{
		WatchID: watch,
		Pretty:  true,
		Human:   true,
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
func WatcherDelete(esClient *elastic7.Client, watch string, outputFmt string) error {
	req := esapi.WatcherDeleteWatchRequest{
		WatchID: watch,
		Pretty:  true,
		Human:   true,
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
func WatcherExecute(esClient *elastic7.Client, watch string, outputFmt string) error {
	// Why do I need a body here?
	//req := esapi.WatcherExecuteWatchRequest{
	//	Body:   io.Reader,
	//	Pretty: true,
	//	Human:  true,
	//}

	//err := request(req, esClient)
	fmt.Println("not yet implemented")
	//return err
	return nil
}
func WatcherGetStats(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.WatcherStatsRequest{
		// Need to figure out how to pass a filter here
		//Matric: []string
		Pretty: true,
		Human:  true,
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
func WatcherAck(esClient *elastic7.Client, watch string, outputFmt string) error {
	//req := esapi.CHANGEME{
	//	Format: "json",
	//	Pretty: true,
	//}

	//err := request(req, esClient)
	fmt.Println("not yet implemented")
	//return err
	return nil
}
func WatcherActivate(esClient *elastic7.Client, watch string, outputFmt string) error {
	var watcherDesiredState bool = true

	// Get the current state of the watcher
	isActive, err := isWatchActive(esClient, watch)
	if err != nil {
		return err
	}

	// Call api to change watch state if needed
	if isActive == watcherDesiredState {
		fmt.Printf("Watcher '%s' is in desired state. No action needed.\n", watch)
	} else {
		req := esapi.WatcherActivateWatchRequest{
			WatchID: watch,
			Pretty:  true,
			Human:   true,
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
	}
	return err
}

func isWatchActive(esClient *elastic7.Client, watch string) (bool, error) {
	req := esapi.WatcherGetWatchRequest{
		WatchID:    watch,
		FilterPath: []string{"status.state.active"},
	}

	var output struct {
		Status struct {
			State struct {
				Active bool `json:"active"`
			} `json:"state"`
		} `json:"status"`
	}

	// Error in request execution
	res, err := req.Do(context.Background(), esClient.Transport)
	if err != nil {
		return false, err
	}

	// Read response body
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	// Error in http response
	if res.StatusCode != 200 {
		return false, fmt.Errorf("Status Code is %v rather than 200.\n Printed error is: \n%v\n. Exiting...\n", res.StatusCode, string(b))
	}

	// Make sense of the bytes
	err = json.Unmarshal(b, &output)
	if err != nil {
		return false, err
	}

	return output.Status.State.Active, nil

}
func WatcherDeactivate(esClient *elastic7.Client, watch string, outputFmt string) error {
	var watcherDesiredState bool = false

	// Get the current state of the watcher
	isActive, err := isWatchActive(esClient, watch)
	if err != nil {
		return err
	}

	// Call api to change watch state if needed
	if isActive == watcherDesiredState {
		fmt.Printf("Watcher '%s' is in desired state. No action needed.\n", watch)
	} else {
		req := esapi.WatcherDeactivateWatchRequest{
			WatchID: watch,
			Pretty:  true,
			Human:   true,
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
	}

	return err
}

func WatcherServiceStop(esClient *elastic7.Client) error {
	//req := esapi.CHANGEME{
	//	Format: "json",
	//	Pretty: true,
	//}

	//err := request(req, esClient)
	fmt.Println("not yet implemented")
	//return err
	return nil
}

// /.watches/_search?filter_path=hits.hits._id
func WatcherServiceStart(esClient *elastic7.Client) error {
	//req := esapi.CHANGEME{
	//	Format: "json",
	//	Pretty: true,
	//}

	//err := request(req, esClient)
	fmt.Println("not yet implemented")
	return nil
}

type GetWatcherCountStruct struct {
	Count int `json:"count"`
}

// Get an exact count of watchers (hits) in the .watches index
func getWatcherCount(esClient *elastic7.Client) (int, error) {

	req := esapi.CountRequest{
		Human:  false,
		Pretty: false,
		Index:  []string{".watches"},
	}
	b, err := esutil.RequestNew(req, esClient)
	if err != nil {
		return -1, err

	}

	var output GetWatcherCountStruct

	err = json.Unmarshal(b, &output)
	if err != nil {
		return -1, err
	}

	return output.Count, nil
}

// key-containing struct necessary for WatcherListRes struct
type WatcherHits struct {
	ID string `json:"_id"`
}

// Expected response for WatcherList function
type WatcherListRes struct {
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		Hits []WatcherHits `json:"hits"`
	} `json:"hits"`
}

// Request body to WatcherList function
type WatcherListReq struct {
	Size int `json:"size"`
}

// Output format for WatcherList function
type WatcherListOutput struct {
	Watchers []string `json:"watchers" yaml:"watchers"`
}

// List watchers present on the cluster
func WatcherList(esClient *elastic7.Client, outputFmt string) error {

	// Create the request with a large inital size
	var initSize int = 1000
	r, err := esutil.JSONToReader(
		WatcherListReq{Size: initSize},
	)
	if err != nil {
		// JSON Marshalling error
		return err
	}

	req := esapi.SearchRequest{
		Human:      false,
		Pretty:     false,
		Body:       r,
		Index:      []string{".watches"},
		FilterPath: []string{"hits.hits._id", "hits.total"},
	}

	// Expected output format for this query
	var output WatcherListRes

	b, err := esutil.RequestNew(req, esClient)

	// Make sense of the bytes
	err = json.Unmarshal(b, &output)
	if err != nil {
		return err
	}

	var watcherSize int

	// Get the exact count if the returned size is larger than the
	//preallocated size
	if !(initSize > output.Hits.Total.Value && output.Hits.Total.Relation != "gt") {
		watcherSize, err = getWatcherCount(esClient)
		if err != nil {
			return err
		}
	} else {
		// Discard initSize in favor of the actual count
		watcherSize = output.Hits.Total.Value
	}
	watchers := make([]string, watcherSize)

	for i, watcher := range output.Hits.Hits {
		watchers[i] = watcher.ID
	}

	watcherList := WatcherListOutput{Watchers: watchers}

	// Only supported output formats here are yaml and json
	// support for tabular should come from esutil once ParseBytes or similar can
	// accept a well-defined struct
	if outputFmt == "yaml" {
		b, err = yaml.Marshal(watcherList)
	} else {
		b, err = json.MarshalIndent(watcherList, "", "  ")
	}
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	return nil
}

// GET /.watches/_search?filter_path=hits.hits._id,hits.hits._source.status.state.active
func reduceWatchersTo(esClient *elastic7.Client, desiredState bool) (watchers []string, err error) {
	req := esapi.SearchRequest{
		Human:  false,
		Pretty: false,

		Index:      []string{".watches"},
		FilterPath: []string{"hits.hits._id", "hits.hits._source.status.state.active"},
	}

	var output struct {
		Hits struct {
			Hits []struct {
				ID     string `json:"_id"`
				Source struct {
					Status struct {
						State struct {
							Active bool `json:"active"`
						} `json:"state"`
					} `json:"status"`
				} `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	// Error in request execution
	res, err := req.Do(context.Background(), esClient.Transport)
	if err != nil {
		return []string{}, err
	}

	// Read response body
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []string{}, err
	}
	defer res.Body.Close()

	// Error in http response
	if res.StatusCode != 200 {
		return []string{}, fmt.Errorf("Status Code is %v rather than 200.\n Printed error is: \n%v\n. Exiting...\n", res.StatusCode, string(b))
	}

	// Make sense of the bytes
	err = json.Unmarshal(b, &output)
	if err != nil {
		return []string{}, err
	}

	for _, watcher := range output.Hits.Hits {
		if watcher.Source.Status.State.Active == desiredState {
			watchers = append(watchers, watcher.ID)
		}
	}

	return watchers, nil

}

func WatcherShowActive(esClient *elastic7.Client, outputFmt string) error {
	var watcherDesiredState bool = true

	watchers, err := reduceWatchersTo(esClient, watcherDesiredState)
	if err != nil {
		return err
	}

	// Only include the necessary stuff. List comprehension would be great here
	// Refactor into its own function
	// Cluster created watchers
	systemWatchers := make([]string, 0)
	// User created watchers
	userWatchers := make([]string, 0)
	var fc string
	var b []byte

	for _, watcher := range watchers {
		fc = string(watcher[0])
		if fc == "-" {
			systemWatchers = append(systemWatchers, watcher)
		} else {
			userWatchers = append(userWatchers, watcher)
		}
	}

	// Temporary until output format feature is added
	var format string = "json"

	if format == "json" {
		var watchers struct {
			Watchers struct {
				User   []string `json:"user"`
				System []string `json:"system"`
			} `json:"watchers"`
		}
		watchers.Watchers.System = systemWatchers
		watchers.Watchers.User = userWatchers
		b, err = json.MarshalIndent(watchers, "", "  ")
	}
	if format == "yaml" {
		var watchers struct {
			Watchers struct {
				User   []string `yaml:"user"`
				System []string `yaml:"system"`
			} `json:"watchers"`
		}
		watchers.Watchers.System = systemWatchers
		watchers.Watchers.User = userWatchers
		b, err = yaml.Marshal(watchers)
	}
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	return nil
}

// GET /.watches/_search?filter_path=hits.hits._id,hits.hits._source.status.state.active
func WatcherShowInactive(esClient *elastic7.Client, outputFmt string) error {
	var watcherDesiredState bool = false

	watchers, err := reduceWatchersTo(esClient, watcherDesiredState)
	if err != nil {
		return err
	}

	// Only include the necessary stuff. List comprehension would be great here
	// Refactor into its own function
	// Cluster created watchers
	systemWatchers := make([]string, 0)
	// User created watchers
	userWatchers := make([]string, 0)
	var fc string
	var b []byte

	for _, watcher := range watchers {
		fc = string(watcher[0])
		// ask elastic booth about the dash prefix
		// Dash prefix is specific to elastic cloud clusters
		if fc == "-" {
			systemWatchers = append(systemWatchers, watcher)
		} else {
			userWatchers = append(userWatchers, watcher)
		}
	}

	// Temporary until output format feature is added
	var format string = "json"

	if format == "json" {
		var watchers struct {
			Watchers struct {
				User   []string `json:"user"`
				System []string `json:"system"`
			} `json:"watchers"`
		}
		watchers.Watchers.System = systemWatchers
		watchers.Watchers.User = userWatchers
		b, err = json.MarshalIndent(watchers, "", "  ")
	}
	if format == "yaml" {
		var watchers struct {
			Watchers struct {
				User   []string `yaml:"user"`
				System []string `yaml:"system"`
			} `json:"watchers"`
		}
		watchers.Watchers.System = systemWatchers
		watchers.Watchers.User = userWatchers
		b, err = yaml.Marshal(watchers)
	}
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	return nil

}
