package esutil

import (
	"context"
	"encoding/json"
	"fmt"
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io/ioutil"
	"strings"
)

type esRequest interface {
	Do(context.Context, esapi.Transport) (*esapi.Response, error)
}

// Deserialize a response into a generic interface
func Des(esRes *esapi.Response) (r map[string]interface{}, err error) {
	//statusCode := esRes.StatusCode
	//header := esRes.Header

	if err = json.NewDecoder(esRes.Body).Decode(&r); err != nil {
		return r, err
	}
	defer esRes.Body.Close()

	return r, err
}

// Display a generic interface as yaml
func MapToYamlish(r map[string]interface{}, s int) {
	// Going to rely on the value of a null int here. Wish me luck...
	// j.k, Golang doesn't like that.
	var nestedMap map[string]interface{}
	spaces := s
	if false {
		fmt.Printf("%v", spaces)
	}
	for k, v := range r {
		switch v.(type) {
		default:
			// Did not find a map as a value
			fmt.Printf("%s%v: '%v'\n", strings.Repeat(" ", 2*(spaces)), k, v)
		case map[string]interface{}:
			// Found a map as a value
			fmt.Printf("%v:\n", k)
			// Explicitly create a map from that value, recurse
			nestedMap = v.(map[string]interface{})
			MapToYamlish(nestedMap, spaces+1)
			// I need a way to use multiple interface types and iterate over them
			// Maybe I add logic to figure out how to determine the type of the map
			// interface thing?
			//case map[int]interface{}:
			//	fmt.Printf("%v:\n", k)
			//	nestedMap = v.(map[int]interface{})
			//	mapToYamlish(nestedMap, spaces+1)
		}
	}
}

// Generic function used to execute requests
func Request(r esRequest, c *elastic7.Client) error {
	// In the future,
	// Request should attempt to insert the format of choice into esRequest
	// (check if it has a Format option) If it does, add it and return the bytes
	// as is. If it does not, use an interface of map[string]interface{} and
	// use methods to convert to desired format. Of note is that responses should
	// be json already, so no action should be required for a default of or
	// explicitly defined json format. This decoupling help the output format
	// future in the future. MapToYamlish can almost certainly go away in this
	// after this body of work.

	res, err := r.Do(context.Background(), c.Transport)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("Status Code is %v rather than 200. Exiting...\n", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	fmt.Printf("%+v\n", string(b))

	return nil
}
