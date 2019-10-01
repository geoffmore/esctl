package escmd

import (
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// https://www.elastic.co/guide/en/elasticsearch/reference/7.2/cat.html

// GET /_cat/indices
func CatIndices(esClient *elastic7.Client) error {
	req := esapi.CatIndicesRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/aliases
func CatAliases(esClient *elastic7.Client) error {
	req := esapi.CatAliasesRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/templates
func CatTemplates(esClient *elastic7.Client) error {
	req := esapi.CatTemplatesRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/allocation
func CatAllocation(esClient *elastic7.Client) error {
	req := esapi.CatAllocationRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/shards
func CatShards(esClient *elastic7.Client) error {
	req := esapi.CatShardsRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/master
func CatMaster(esClient *elastic7.Client) error {
	req := esapi.CatMasterRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/nodes
func CatNodes(esClient *elastic7.Client) error {
	req := esapi.CatNodesRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/tasks
func CatTasks(esClient *elastic7.Client) error {
	req := esapi.CatTasksRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/segments
func CatSegments(esClient *elastic7.Client) error {
	req := esapi.CatSegmentsRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/count
func CatCount(esClient *elastic7.Client) error {
	req := esapi.CatCountRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/recovery
func CatRecovery(esClient *elastic7.Client) error {
	req := esapi.CatRecoveryRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/health
func CatHealth(esClient *elastic7.Client) error {
	req := esapi.CatHealthRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/pending_tasks
func CatPendingTasks(esClient *elastic7.Client) error {
	req := esapi.CatPendingTasksRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/thread_pool
func CatThreadPool(esClient *elastic7.Client) error {
	req := esapi.CatThreadPoolRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/thread_pool/{thread_pools}
// GET /_cat/plugins
func CatPlugins(esClient *elastic7.Client) error {
	req := esapi.CatPluginsRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/fielddata
func CatFielddata(esClient *elastic7.Client) error {
	req := esapi.CatFielddataRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/nodeattrs
func CatNodeattrs(esClient *elastic7.Client) error {
	req := esapi.CatNodeattrsRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

// GET /_cat/repositories
func CatRepositories(esClient *elastic7.Client) error {
	req := esapi.CatRepositoriesRequest{
		Format: "json",
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}
