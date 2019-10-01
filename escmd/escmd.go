package escmd

import (
	"fmt"
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/geoffmore/esctl-go/esutil"
)

// If the objective is to only return output, structs are not needed. Once
// functions requiring data manipulation are introduced, there will be a need
// for these structs to return. For now, they are being left in the code for
// historic reference

//type CatIndicesResponse struct {
//	Health       string `json:"health"`
//	Status       string `json:"status"`
//	Index        string `json:"index"`
//	UUID         string `json:"uuid"`
//	Pri          string `json:"pri"`
//	Rep          string `json:"rep"`
//	DocsCount    string `json:"docs.count"`
//	DocsDeleted  string `json:"docs.deleted"`
//	StoreSize    string `json:"store.size"`
//	PriStoreSize string `json:"pri.store.size"`
//}

//type CatIndicesStruct struct {
//	Indices []CatIndicesResponse
//}

// All methods I've tested other than es7.Info() are not found outside of
// esapi. In addition, based on the documentation I read on the
// documentation, I would assume I could pass a elasticsearch.Client as my
// transport, but I need to use the Transport type within the object instead

// GET /_cluster/info
func GetClusterInfo(esClient *elastic7.Client) error {
	resp, err := esClient.Info()
	if err != nil {
		return err
	}
	rhee, err := esutil.Des(resp)
	if err != nil {
		return err
	}
	esutil.MapToYamlish(rhee, 0)
	fmt.Println()

	return nil
}

// GET /_cluster/health
func GetClusterHealth(esClient *elastic7.Client) error {
	req := esapi.ClusterHealthRequest{
		// There is no Format field here, so the mapToYamlish function will be
		// required for yaml output
		//Format: "json",
		Human:  true,
		Pretty: true,
	}

	err := request(req, esClient)
	return err
}

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
