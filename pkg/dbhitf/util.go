package dbhitf

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func queryToStr(s minimalSearch) string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func queryToPrettyStr(s minimalSearch) string {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func debugReqIfResFaulty(res *esapi.Response, s minimalSearch) {
	// if response is not 200, show the query that was used for debugging.
	// Should this be broadened to a larger set of Http Response Codes?

	// erroMsg should be of type error in good design, but I assign it to string
	errorMsg := fmt.Sprintf("Received non-200 response '%v'.\nQuery used:\n%v\n",
		res.StatusCode,
		// Currently too lazy to convert esapi.SearchRequest.Body from io.Reader to
		// []byte, so I'm instead going to rely prettyPrintQuery(s) for now
		// Ideally, reqToJson(req) string is the correct function here
		queryToPrettyStr(s),
	)
	if res.StatusCode != 200 {
		panic(errorMsg)
	}
}
