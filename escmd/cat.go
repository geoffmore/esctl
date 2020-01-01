package escmd

import (
	elastic7 "github.com/geoffmore/go-elasticsearch/v7"
	"github.com/geoffmore/go-elasticsearch/v7/esapi"
	"github.com/geoffmore/esctl-go/esutil"
	"reflect"
)

// https://www.elastic.co/guide/en/elasticsearch/reference/7.2/cat.html

// GET /_cat/indices
func CatIndices(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatIndicesRequest{
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

// GET /_cat/aliases
func CatAliases(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatAliasesRequest{
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

// GET /_cat/templates
func CatTemplates(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatTemplatesRequest{
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

// GET /_cat/allocation
func CatAllocation(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatAllocationRequest{
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

// GET /_cat/shards
func CatShards(esClient *elastic7.Client, outputFmt string) error {
	//var help bool = true
	req := esapi.CatShardsRequest{
		Human:  true,
		Pretty: true,

		// Temporary addition
		//Help: &help,
		H: []string{"*"},
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

// GET /_cat/master
func CatMaster(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatMasterRequest{
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

// GET /_cat/nodes
func CatNodes(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatNodesRequest{
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

// GET /_cat/tasks
func CatTasks(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatTasksRequest{
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

// GET /_cat/segments
func CatSegments(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatSegmentsRequest{
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

// GET /_cat/count
func CatCount(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatCountRequest{
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

// GET /_cat/recovery
func CatRecovery(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatRecoveryRequest{
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

// GET /_cat/health
func CatHealth(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatHealthRequest{
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

// GET /_cat/pending_tasks
func CatPendingTasks(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatPendingTasksRequest{
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

// GET /_cat/thread_pool
func CatThreadPool(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatThreadPoolRequest{
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

// GET /_cat/thread_pool/{thread_pools}

// GET /_cat/plugins
func CatPlugins(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatPluginsRequest{
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

// GET /_cat/fielddata
func CatFielddata(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatFielddataRequest{
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

// GET /_cat/nodeattrs
func CatNodeattrs(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatNodeattrsRequest{
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

// GET /_cat/repositories
func CatRepositories(esClient *elastic7.Client, outputFmt string) error {
	req := esapi.CatRepositoriesRequest{
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
