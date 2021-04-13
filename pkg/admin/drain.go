// +build optional admin
package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	_ "github.com/elastic/go-elasticsearch/v7/estransport"
	"github.com/geoffmore/esctl/pkg/esutil"
	_ "net"
	"reflect"
	"strings"
)

// Error handling needs to be improved. Maybe different error types based on
// what ES is doing under the hood?

//func Balance(esClient *elastic7.Client, outputFmt string, help bool) error {
//	return nil
//}

// Cordon function doesn't appear to exist in the api, so there is no way to
// implement this

// Helper struct containing useful data node info
type DataNodes struct {
	NodeCount int
	Nodes     []DataNode
}
type DataNode struct {
	ID string
	// User input needs to be validated. Don't think server input does
	//IP   net.IP
	IP   string `json:"ip"`
	Name string `json:"name"`
}
type DataNodeManager struct {
	NodeDefs map[string]DataNode `json:"nodes`
	NodeInfo map[string]int      `json:"_nodes"`
}

// Helper func to convert DataNode.IP to string
//func (node DataNode) GetIPString() string {
//	return node.IP.String()
//}

// Helper function to get a list of data nodes
// This could probaby be public
func getDataNodes(esClient *elastic7.Client) (nodes DataNodes, err error) {

	req := esapi.NodesInfoRequest{
		Human:  false,
		Pretty: false,
		// There is no reason to consider non-data nodes
		Metric:     []string{"data:true"},
		FilterPath: []string{},
	}
	_ = esutil.SetFormat(reflect.ValueOf(&req).Elem(), "json")
	// // Make a request to get bytes
	b, err := esutil.RequestNew(req, esClient)
	if err != nil {
		return nodes, err
	}

	var bytesStruct map[string]interface{}

	json.Unmarshal(b, &bytesStruct)

	// https://stackoverflow.com/questions/44305617
	// What is this madness???
	nodes.NodeCount = int(bytesStruct["_nodes"].(map[string]interface{})["total"].(float64))

	// The length of nodes.DataNodes is known. Does it make sense to treat this
	// as a fixed length array and assign to indices?

	var id, name, ip string

	// If the underlying schema changes, this will surely fail. Can a partial
	// definition exist for a schema? Curious if certain fields can be fully
	// defined and others left as map[string]interface{}

	for k, v := range bytesStruct["nodes"].(map[string]interface{}) {
		id = k
		name = v.(map[string]interface{})["name"].(string)
		ip = v.(map[string]interface{})["ip"].(string)

		nodes.Nodes = append(nodes.Nodes, DataNode{
			ID:   id,
			IP:   ip,
			Name: name,
		})
	}

	return nodes, err
}

// There are a few different methods used to traverse json currently. This is a
// deficiency and a standard shoud be established

// https://stackoverflow.com/questions/21268000/
// Options are unmarshalling and iterating or creating a well-defined struct
type AllocationExclusions struct {
	Transient struct {
		Names string `json:"cluster.routing.allocation.exclude._name"`
		IPs   string `json:"cluster.routing.allocation.exclude._ip"`
	} `json:"transient"`
}

// Helper to get cluster settings
// For functions that rely on this function, an assumption is made that the user
// never wants exclude a node that is invalid and/or does not exist in the
// cluster. If this is not the case, please open an issue against this file
func getExcludeSettings(esClient *elastic7.Client) (AllocationExclusions, error) {
	var flatSettings bool = true
	var includeDefaults bool = true

	var exclusions AllocationExclusions

	// Verbose flag should give the api call
	req := esapi.ClusterGetSettingsRequest{
		Human:  false,
		Pretty: false,
		// Reduce data to necessary path
		FilterPath: []string{"transient.cluster*"},
		// Necessary to get all of the potentially relevant fields
		IncludeDefaults: &includeDefaults,
		// Flat settings are kinda useful, but don't work when combined with
		// a filter path. Why? Is
		// escaping the filter path necessary? Avoiding this question for now
		FlatSettings: &flatSettings,
	}
	_ = esutil.SetFormat(reflect.ValueOf(&req).Elem(), "json")
	// // Make a request to get bytes
	b, err := esutil.RequestNew(req, esClient)
	if err != nil {
		return exclusions, err
	}

	err = json.Unmarshal(b, &exclusions)
	// No error handling here, but if it was, there would be some sort of json
	// unmarshalling error
	return exclusions, err
}

// helper func to replicate "if foo in []bar" syntax
// Shamelessly copied from
// https://stackoverflow.com/questions/15323767
func stringInSlice(a string, list []string) (bool, int) {
	for index, b := range list {
		if b == a {
			return true, index
		}
	}
	// Not sure how to handle the case where something isn't found
	return false, -1
}

// Might as well add the list from the calling function into the signature to
// save a redundant api call

// Feels  kinda dangerous to keep business logic outside of the function that
// makes the api call
// This probably needs an outputFmt field since it makes a call to
// esutil.ParseBytes
func (ae AllocationExclusions) rn(esClient *elastic7.Client, name string) (bool, error) {
	var flatSettings bool = true

	excludedNameList := strings.Split(ae.Transient.Names, ",")

	// Check if values exist in current settings and remove them
	nameExists, index := stringInSlice(name, excludedNameList)
	if nameExists {
		// https://github.com/golang/go/wiki/SliceTricks
		// https://stackoverflow.com/questions/16248241
		excludedNameList = append(excludedNameList[:index], excludedNameList[index+1:]...)
	} else {
		// Name doesn't exist. No need to proceed
		return false, nil
	}
	ae.Transient.Names = strings.Join(excludedNameList, ",")

	// Convert back to bytes
	b, err := json.Marshal(ae)
	if err != nil {
		return false, err
	}

	r := bytes.NewReader(b)
	// Verbose flag should give the api call
	req := esapi.ClusterPutSettingsRequest{
		Human:  false,
		Pretty: false,

		Body: r,

		FlatSettings: &flatSettings,
	}

	// Boilerplate
	//changedField := esutil.SetFormat(reflect.ValueOf(&req).Elem(), outputFmt)
	// // Make a request to get bytes
	b, err = esutil.RequestNew(req, esClient)
	if err != nil {
		return false, err
	}

	err = esutil.ParseBytes(b, true, "json")
	if err != nil {
		return false, err
	}
	return true, nil
}

//func (ae AllocationExclusions) ri(esClient *elastic7.Client, name string) (bool, error) {
//	return false, nil
//}

//func removeName(esClient *elastic7.Client, name string, names []string) error {
//
//	// If the AllocationExclusion object is passed rather than a list of names,
//	// this arguably becomes easier to work with. Moves stuff closer to
//	// interfaces, idk. There needs to be an io.Reader anyway, so maybe this is
//	// the way to go?
//
//	// ClusterPutSettingsRequest
//	var flatSettings bool = true
//	var includeDefaults bool = true
//
//	// Verbose flag should give the api call
//	req := esapi.ClusterGetSettingsRequest{
//		Human:  false,
//		Pretty: false,
//		// Reduce data to necessary path
//		FilterPath: []string{"transient.cluster*"},
//		// Necessary to get all of the potentially relevant fields
//		IncludeDefaults: &includeDefaults,
//		// Flat settings are kinda useful, but don't work when combined with
//		// a filter path. Why? Is
//		// escaping the filter path necessary? Avoiding this question for now
//		FlatSettings: &flatSettings,
//	}
//}

// Might as well add the list from the calling function into the signature to
// save a redundant api call
//func removeIP(esClient *elastic7.Client, ip string, ips []string) error {
//	// ClusterPutSettingsRequest
//
//}

//func addName() {}
//func addIP()   {}

//func (ae AllocationExclusions) libn(esClient *elastic7.Client, name string) (string, error)
//func (ae AllocationExclusions) lnbi(esClient *elastic7.Client, ip string) (string, error)

// Empty string indicates no match found and is acceptable
//func lookupIPByName(esClient *elastic7.Client, name string) (string, error) {
//	// ClusterGetSettingsRequest + getDataNodes + io.Reader
//}

// Empty string indicates no match found and is acceptable
//func lookupNameByIP(esClient *elastic7.Client, ip string) (string, error) {
//	// ClusterGetSettingsRequest + getDataNodes + io.Reader
//}

// Functionally similar to kubectl's uncordon command.
// Given a name OR ip address, removes both node name AND ip address from being excluded
// In the future, this should take arguments <property> and <value> and remove
// the values from the necessary keys
// Uncordon is easier than Drain because it doesn't have the risk of adding
// garbage
// Cluster settings affected
// transient.cluster.routing.allocation.exclude._name
// transient.cluster.routing.allocation.exclude._ip
// note: v1 of this function only has the ability to do removals by name
func Uncordon(esClient *elastic7.Client, outputFmt string, help bool, node string) error {

	// Get current settings
	exclusions, err := getExcludeSettings(esClient)

	if err != nil {
		return err
	}

	// Transform settings into lists

	// Need to run the individual checks on the name/ip types and get a specific
	// boolean value for each to assist later in the program

	// Do a removal by name
	// rn should perform a lookup in the event node is an ip
	// Hint: requires DataNodes struct
	nameRemoved, nameError := exclusions.rn(esClient, node)
	if nameRemoved && nameError == nil {
		fmt.Printf("Name '%s' removed from exclusion list\n", node)
	} else {
		fmt.Printf("Name '%s' not removed from exclusion list. Somehow invalid\n", node)
		return nameError
	}
	// TODO Do a removal by ip
	// ri should perform a lookup in the event node is a name
	//ipRemoved, ipError := exclusions.rn(esClient, node)
	//if ipRemoved && ipError == nil {
	//	fmt.Printf("IP '%s' removed from exclusion list\n", node)
	//} else {
	//	fmt.Printf("IP '%s' not removed from exclusion list. Somehow invalid\n", node)
	//	return ipError
	//}

	//if !(nameExists || ipExists) {
	//	return fmt.Errorf("input not cluster exclude settings")
	//}

	return nil
}

//func Drain(esClient *elastic7.Client, outputFmt string, help bool, node string) error {
//
//	// The exclude can either be via index-level settings or cluster-level settings.
//	// Cordoning a node without removing its existing shards has no clear path yet
//
//	//{ "transient": {
//	//        "cluster.routing.allocation.exclude._ip": ["172.25.12.143", "172.25.137.180"]
//	//    } }
//	// The above syntax is invalid. How can multiple nodes be excluded from shard
//	// allocation?
//
//	// Use flat settings to make everything much easier
//
//	// // Input validation phase
//	var nodeExists bool = false
//	//https://www.elastic.co/guide/en/elasticsearch/reference/master/cluster.html
//	// ^ is cool because I can filter by node role
//	//	esctl api get /_nodes/_all,data:true?filter_path=_nodes.total
//	//{
//	//  "_nodes" : {
//	//    "total" : 3
//	//  }
//	//}
//	//
//	// Get data node count
//	// I _could_ filter /_nodes to get my info so I have the desired length, ip,
//	// intance id, and name. Only drawback is that I have to filter more
//
//	//nodeReq := esapi.CatNodesRequest{
//	//	Human:      false,
//	//	Pretty:     false,
//	//	FilterPath: []string{""},
//	//}
//	// nodeBytes, err := esutil.RequestNew(nodeReq, esClient)
//	// There is kind of a race condition here where I need to figure out the
//	// length of the returned list before the object is instantiated. Hopefully
//	// stuff is automatically extended rather than truncated. For now, a static
//	// number will be used.
//
//	// Remember, only data nodes should be considered here
//	// nodesData := make([]string, 10)
//	if !nodeExists {
//		panic("Node '%s' doesn't exist in cluster, unable to proceed!")
//	}
//
//	// - Verify ip or node name exists in the cluster. Fail if otherwise
//	// Get cat nodes filter ip, name
//	// loop object, start false true if value is in the array
//
//	// // Prerequesite checks phase
//	// - Get cluster allocation settings
//	settingsReq := esapi.ClusterGetSettingsRequest{
//		Human:      false,
//		Pretty:     false,
//		FilterPath: []string{"*.cluster.routing.allocation.exclude"},
//	}
//	// Should there be a struct containing top level fields "defaults",
//	// "transient", and "persistent"? May make analysis a bit easier
//
//	settingsBytes, err := esutil.RequestNew(settingsReq, esClient)
//	if err != nil {
//		return err
//	}
//
//	settingsData := make(map[string]interface{})
//	err = json.Unmarshal(settingsBytes, &settingsData)
//	if err != nil {
//		panic("JSON Unmarshalling error")
//	}
//	// Perform analysis of settings
//	// Get transient settings as its own object
//
//	// Debug
//	//fmt.Println(string(settingsBytes))
//
//	// - Cluster health check
//	req := esapi.CatHealthRequest{
//		Human:  false,
//		Pretty: false,
//		Format: "json",
//	}
//
//	b, err := esutil.RequestNew(req, esClient)
//	if err != nil {
//		return err
//	}
//
//	clusterHealthData := make([]map[string]interface{}, 1)
//	err = json.Unmarshal(b, &clusterHealthData)
//	if err != nil {
//		panic("JSON Unmarshalling error")
//	}
//
//	//Convert static boolean to flag in the future
//	var force bool = false
//	if clusterHealthData[0]["status"] != "green" || force {
//		panic("Use force option to proceed. Cluster health was likely not green")
//	}
//
//	// - Generate shard migration plan (may be able to let Elasticsearch handle
//	// this (storage, shard count)
//
//	// Get previous cluster state
//
//	// // Ensure no shards can land on the node
//	// Stub for Cordon function
//
//	// // Remove existing shards from the node
//
//	// The string input containing the path to the node should be interpreted as an
//	// ip or node name. Depending on the detected type, there should be a different
//	// API call
//	// Kept for future inspiration
//
//	// // Print bytes
//	// Get parsable json output
//	// Need a MarshalBytes function that returns a map[string]interface{} for
//	// further processing
//	// esutil.MarshalBytes()
//
//	//err = esutil.ParseBytes(b, changedField, outputFmt)
//	return err
//
//}
