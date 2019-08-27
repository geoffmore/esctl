package dbhitf

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// initDbhitfStruct returns a dbhitfStruct
func initDbhitfStruct(hostFieldName, hostFieldValue, index, time, existsFieldName string) dbhitfStruct {
	//dbs = dbshiftStruct{
	return dbhitfStruct{
		hostFieldName:   hostFieldName,
		hostFieldValue:  hostFieldValue,
		index:           index,
		time:            time,
		existsFieldName: existsFieldName,
	}
}

// genQuery generates a minimalSearch from a dbhitfStruct for output as a json
// document
func genQuery(dbs dbhitfStruct) (s minimalSearch) {
	// How can I make m part of the Match initialization?
	m := make(map[string]string)
	m[dbs.hostFieldName] = dbs.hostFieldValue

	// How can I clean this up?
	s = minimalSearch{
		Size: 0,
		Query: Query{
			Bool: Bool{
				Must: []MustInterface{
					MatchJson{Match: m},
					RangeJson{
						Range{
							Timestamp: Timestamp{
								Gte: dbs.time,
							},
						},
					},
				},
			},
		},
		Aggs: Aggs{
			FieldExistsBool: FieldExistsBool{
				FEBFilters: FEBFilters{
					OtherBucketKey: "field_dne",
					OtherBucket:    true,
					Filters: Filters{
						FieldExists: FieldExists{
							Exists: Exists{
								Field: dbs.existsFieldName,
							},
						},
					},
				},
			},
		},
	}
	return s
}

// Construct search
func constructSearch(es *elastic7.Client, dbs dbhitfStruct, s minimalSearch) (req esapi.SearchRequest) {
	// I have a json.Valid() here, but does it add value?
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	isValid := json.Valid(b)
	if !isValid {
		panic("Marshalled json is invalid")
	}
	reader := bytes.NewReader(b)
	req = esapi.SearchRequest{
		Index: []string{dbs.index},
		Body:  reader,
	}
	return req
}

// Execute search
func executeSearch(es *elastic7.Client, req esapi.SearchRequest) (res *esapi.Response) {
	var err error
	res, err = req.Do(context.TODO(), es.Transport)
	if err != nil {
		panic(err)
	}
	return res
}

// Generate stats from request body
func processRes(res *esapi.Response) fieldStats {
	var dbr DbhitfResponse
	//var output string

	// I will admit ioutil is a bit lazy here. Not sure if I should use a byte
	// buffer or another solution
	// https://medium.com/@matryer/golang-advent-calendar-day-seventeen-io-reader-in-depth-6f744bb4320b
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	if json.Valid(b) {
		json.Unmarshal(b, &dbr)
	} else {
		panic("Invalid response. Expecting JSON")
	}

	fieldExistsHits := dbr.Aggregations.FieldExistsBool.Buckets.FieldExists.DocCount
	fieldDneHits := dbr.Aggregations.FieldExistsBool.Buckets.FieldDne.DocCount

	// In the case where the field does not exist, fieldDneHits has all of the
	// documents, fieldExistsHits has 0. Since the presence of the field is
	// boolean, there is no need worrying about dividing by 0
	fieldExistsPct := float64(fieldExistsHits) / float64(fieldExistsHits+fieldDneHits)

	return fieldStats{
		fieldExistCount: fieldExistsHits,
		fieldDneCount:   fieldDneHits,
		fieldExistPct:   fieldExistsPct}

}

// Display overall results
func displayResults(dbs dbhitfStruct, stats fieldStats) string {
	// dbs will be used when there is an added verbosity setting
	// Using Sprintf allows me to use option flags for output in the future

	// Normal method
	output := fmt.Sprintf(
		"Hits where field exists: %v\nHits where field does not exist: %v\nPercentage of documents with field: %.3f%%",
		stats.fieldExistCount,
		stats.fieldDneCount,
		stats.fieldExistPct*100,
	)

	// Json method
	//bo, err := json.Marshal(outputStuff)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%+v\n", outputStuff)
	//fmt.Printf("%+v\n", string(bo))

	return output

}

// Main function called by cobra
func Search(es *elastic7.Client, hostFieldName, hostFieldValue, indexName, timeBack, existsFieldName string) {

	// I am not a fan of the panic pattern in this package
	dbs := initDbhitfStruct(hostFieldName, hostFieldValue, indexName, timeBack, existsFieldName)
	query := genQuery(dbs)
	// Do I need defers on these requests?
	req := constructSearch(es, dbs, query)
	res := executeSearch(es, req)

	// Panic and show query if response code is not 200
	debugReqIfResFaulty(res, query)
	// Show stats
	stats := processRes(res)

	fmt.Printf("%s\n", displayResults(dbs, stats))

}
