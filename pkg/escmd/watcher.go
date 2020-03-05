package escmd

import (
	"context"
	"encoding/json"
	"fmt"
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/geoffmore/esctl/pkg/esutil"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

// Watchers starting with dashes don't work currently because of cobra
// 'pods_pending' is valid
// '-pods_pending' is interpreted as having a flag and does not work

func WatcherPut(esClient *elastic7.Client, watch string) error {
	//req := esapi.WatcherPutWatchRequest{
	//	Format: "json",
	//	Pretty: true,
	//}

	// -f for file or -stdin for standard in
	//err := request(req, esClient)
	fmt.Println("not yet implemented")
	return nil
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

func WatcherList(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.SearchRequest{
		Human:  true,
		Pretty: true,

		Index:      []string{".watches"},
		FilterPath: []string{"hits.hits._id"},
	}

	// Expected output format for this specific query
	var output struct {
		Hits struct {
			Hits []struct {
				ID string `json:"_id"`
			} `json:"hits"`
		} `json:"hits"`
	}

	// escmd.request doesn't work here because we need data manipulation
	//err := request(req, esClient)

	// Error in request execution
	res, resErr := req.Do(context.Background(), esClient.Transport)
	if resErr != nil {
		return resErr
	}

	// Read response body
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Error in http response
	if res.StatusCode != 200 {
		return fmt.Errorf("Status Code is %v rather than 200.\n Printed error is: \n%v\n. Exiting...\n", res.StatusCode, string(b))
	}

	// Make sense of the bytes
	err = json.Unmarshal(b, &output)
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

	for _, watcher := range output.Hits.Hits {
		fc = string(watcher.ID[0])
		if fc == "-" {
			systemWatchers = append(systemWatchers, watcher.ID)
		} else {
			userWatchers = append(userWatchers, watcher.ID)
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

//
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
