package esutil

import (
	"encoding/json"
	"fmt"
	"strings"
)

import (
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

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
