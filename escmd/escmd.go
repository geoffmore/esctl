package escmd

import (
	"context"
	"fmt"
)

import (
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/geoffmore/esctl-go/esutil"
)

// All methods I've tested other than es7.Info() are not found outside of
// esapi. In addition, based on the documentation I read on the
// documentation, I would assume I could pass a elasticsearch.Client as my
// transport, but I need to use the Transport type within the object instead

func GetClusterInfo(esClient *elastic7.Client) {
	resp, _ := esClient.Info()
	rhee := esutil.Des(resp)
	esutil.MapToYamlish(rhee, 0)
	fmt.Println()
}

func GetClusterHealth(esClient *elastic7.Client) {
	req := esapi.ClusterHealthRequest{}
	res, _ := req.Do(context.Background(), esClient.Transport)
	rhee := esutil.Des(res)
	esutil.MapToYamlish(rhee, 0)
	fmt.Println()
}
