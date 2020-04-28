// +build optional admin

package admin

import (
	// Standard
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/geoffmore/esctl/pkg/esutil"
	"reflect"
	// Package specific
	"encoding/json"
)

func ShowShards(esClient *elastic7.Client, outputFmt string, help bool) error {

	var sortArray []string = []string{"store:desc,index,shard"}

	var verbose bool = true

	req := esapi.CatShardsRequest{
		Human:  true,
		Pretty: true,

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

// Show the largest shards on data nodes by node
func ShowBigShards(esClient *elastic7.Client, outputFmt string, help bool, shardsPerNode int) error {

	// Test, needs to be an optional param with default
	var verbose bool = true

	// Array is sorted server-side, so logic is simplified here
	var sortArray []string = []string{"node:desc,store:desc"}

	req := esapi.CatShardsRequest{
		Human:  false,
		Pretty: false,
		// Implied for nice numbers
		//Bytes: "b",

		S: sortArray,

		// This may be a v7 vs v8 issue. Unlocking v7
		Help: &help,
		V:    &verbose,
	}

	// Start with json
	// Ignoring this has consequences, probably, but ParseBytes needs to know to
	// convert to the desired outputFmt
	//changedField := esutil.SetFormat(reflect.ValueOf(&req).Elem(), "json")
	_ = esutil.SetFormat(reflect.ValueOf(&req).Elem(), "json")

	// https://stackoverflow.com/questions/48340687
	//type nJson map[string]interface{}
	//var jsonData []map[string]interface{}
	jsonData := []map[string]interface{}{}

	// // Make a request to get bytes
	b, err := esutil.RequestNew(req, esClient)
	if err != nil {
		return err
	}

	// Marshal json into variable
	err = json.Unmarshal(b, &jsonData)
	if err != nil {
		return err
	}

	dataMap := map[string][]map[string]interface{}{}

	err = shardFunc(jsonData, dataMap, shardsPerNode)
	if err != nil {
		return err
	}

	b1, err := json.Marshal(dataMap)
	err = esutil.ParseBytes(b1, false, outputFmt)
	return err
}

// Show the smallest shards on data nodes by node
func ShowSmallShards(esClient *elastic7.Client, outputFmt string, help bool, shardsPerNode int) error {

	// Array is sorted server-side, so logic is simplified here
	var sortArray []string = []string{"store:asc,index,shard"}

	var verbose bool = true

	req := esapi.CatShardsRequest{
		// Compact for performance
		Human:  false,
		Pretty: false,

		// Implied for nice numbers
		// Bytes: "b",

		S: sortArray,

		// This may be a v7 vs v8 issue. Unlocking v7
		Help: &help,
		V:    &verbose,
	}

	// Start with json
	// Ignoring this has consequences, probably, but ParseBytes needs to know to
	// convert to the desired outputFmt
	//changedField := esutil.SetFormat(reflect.ValueOf(&req).Elem(), "json")
	_ = esutil.SetFormat(reflect.ValueOf(&req).Elem(), "json")

	// https://stackoverflow.com/questions/48340687
	//type nJson map[string]interface{}
	//var jsonData []map[string]interface{}
	jsonData := []map[string]interface{}{}

	// // Make a request to get bytes
	b, err := esutil.RequestNew(req, esClient)
	if err != nil {
		return err
	}

	// Marshal json into variable
	err = json.Unmarshal(b, &jsonData)
	if err != nil {
		return err
	}

	dataMap := map[string][]map[string]interface{}{}

	err = shardFunc(jsonData, dataMap, shardsPerNode)
	if err != nil {
		return err
	}

	b1, err := json.Marshal(dataMap)
	err = esutil.ParseBytes(b1, false, outputFmt)
	return err
}

// Shared routine for ShowBigShards and ShowSmallShards
// This may not work without interface{}{} in jsonData and mapData
func shardFunc(jsonData []map[string]interface{}, dataMap map[string][]map[string]interface{}, shardsPerNode int) error {

	var currentNode, previousNode string
	var i int
	var ok bool

	for _, v := range jsonData {
		currentNode, ok = v["node"].(string)
		// Validate 'node' field present
		if ok {
			// Reset counter when node changes
			if currentNode != previousNode {
				i = 0
			}
			// Always check array length to prevent overflow and deduplicate
			// currentNode-previousNode logic
			if i < shardsPerNode {
				// Map
				_, nodeKeyExists := dataMap[currentNode]
				if !nodeKeyExists {
					dataMap[currentNode] = make([]map[string]interface{}, shardsPerNode)
				}
				dataMap[currentNode][i] = make(map[string]interface{})
				// Assign
				dataMap[currentNode][i] = v

				//Filter
				// Ignoring for now. Need to remove 'ip' field since it
				// is static per node
				delete(dataMap[currentNode][i], "node")
			}
			i += 1
		}
		previousNode = currentNode
	}
	return nil
}
