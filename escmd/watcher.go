package escmd

import (
	"context"
	"encoding/json"
	"fmt"
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func WatcherPut(esClient *elastic7.Client) error {
	//req := esapi.CHANGEME{
	//	Format: "json",
	//	Pretty: true,
	//}

	//err := request(req, esClient)
	fmt.Println("not yet implemented")
	return nil
}

func WatcherGet(esClient *elastic7.Client) error {
	//req := esapi.CHANGEME{
	//	Format: "json",
	//	Pretty: true,
	//}

	//err := request(req, esClient)
	fmt.Println("not yet implemented")
	return nil
}
func WatcherDelete(esClient *elastic7.Client) error {
	//req := esapi.CHANGEME{
	//	Format: "json",
	//	Pretty: true,
	//}

	//err := request(req, esClient)
	fmt.Println("not yet implemented")
	return nil
}
func WatcherExecute(esClient *elastic7.Client) error {
	//req := esapi.CHANGEME{
	//	Format: "json",
	//	Pretty: true,
	//}

	//err := request(req, esClient)
	fmt.Println("not yet implemented")
	return nil
}
func WatcherGetStats(esClient *elastic7.Client) error {
	//req := esapi.CHANGEME{
	//	Format: "json",
	//	Pretty: true,
	//}

	//err := request(req, esClient)
	fmt.Println("not yet implemented")
	return nil
}
func WatcherAck(esClient *elastic7.Client) error {
	//req := esapi.CHANGEME{
	//	Format: "json",
	//	Pretty: true,
	//}

	//err := request(req, esClient)
	fmt.Println("not yet implemented")
	return nil
}
func WatcherActivate(esClient *elastic7.Client) error {
	//req := esapi.CHANGEME{
	//	Format: "json",
	//	Pretty: true,
	//}

	//err := request(req, esClient)
	fmt.Println("not yet implemented")
	return nil
}
func WatcherDeactivate(esClient *elastic7.Client) error {
	//req := esapi.CHANGEME{
	//	Format: "json",
	//	Pretty: true,
	//}

	//err := request(req, esClient)
	fmt.Println("not yet implemented")
	return nil
}
func WatcherServiceStop(esClient *elastic7.Client) error {
	//req := esapi.CHANGEME{
	//	Format: "json",
	//	Pretty: true,
	//}

	//err := request(req, esClient)
	fmt.Println("not yet implemented")
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

func WatcherList(esClient *elastic7.Client) error {
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
