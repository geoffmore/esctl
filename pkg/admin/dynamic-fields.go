package admin

import (
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	_ "github.com/elastic/go-elasticsearch/v7/estransport"
	_ "net"
)

func getDynamicFields(esClient *elastic7.Client, idxPatterns ...string) (dynamicFields []string, someErr error) {

	// Join idxPatterns to be comma delimited
	//

	// Sam's suggestion - use field caps to determine an exhaustive list of fields
	// Need an inner list of fields defined in a given index template

	// Given a set of index patterns, determine which templates are used
	//- Ignore component templates, runtime fields for now
	//- Ignore priority for now.
	// I think I need a structure {idxPattern, []{template: priority}}, but I'm not sure

	//   Get all templates (need name, index pattern, and mappings and ...)
	// /_template is the api endpoint

	// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-mapping.html#indices-get-mapping

	/*
	{
	  "ilm-history": {
	    "index_patterns": [
	      "ilm-history-1*"
	    ],
	    "mappings": {
	      "dynamic": false,
	      "properties": {
	        "index": {
	          "type": "keyword"
	        },
	        "success": {
	          "type": "boolean"
	        },
	        "index_age": {
	          "type": "long"
	        },
	        "@timestamp": {
	          "type": "date",
	          "format": "epoch_millis"
	        },
	        "state": {
	          "dynamic": true,
	          "type": "object",
	          "properties": {
	 */

	// esapi.Template


	return dynamicFields, someErr
}