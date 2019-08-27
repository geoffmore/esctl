package dbhitf

// dbhitfStruct is the struct containing all variables needed to complete the
// query.
type dbhitfStruct struct {
	hostFieldName   string `json:"host-name"`
	hostFieldValue  string `json:"host-value"`
	index           string `json:"index"`
	time            string `json:"time"`
	existsFieldName string `json:"field-name"`
}

// Minimal search structure according to
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
type minimalSearch struct {
	Size  int   `json:"size"`
	Query Query `json:"query"`
	Aggs  Aggs  `json:"aggs"`
}

type fieldStats struct {
	fieldExistCount int     `json:"field-exist-count"`
	fieldDneCount   int     `json:"field-dne-count"`
	fieldExistPct   float64 `json:"field-exist-pct"`
}

type Query struct {
	Bool Bool `json:"bool"`
}

type Bool struct {
	Must []MustInterface `json:"must"`
}

// Adding a MustInterface here so I can properly serialize range, match and
// future must fields
// This adds the extra level of braces needed for proper query json

// For each desired object under []MustInterface,
// Create the normal struct containing the field configuration, then create a
// struct that contains the field key for the struct under a struct named
// <struct>Json
// Range -> RangeJson
// RangeJson adds `json:"range,omitempty"`

type MustInterface interface {
	MustFoo()
}

// Empty functions to satisfy MustInterface interface
func (m RangeJson) MustFoo() {}
func (m MatchJson) MustFoo() {}

type Must struct {
	Range Range             `json:"range,omitempty"`
	Match map[string]string `json:"match,omitempty"`
}

type Range struct {
	Timestamp Timestamp `json:"@timestamp"`
}
type RangeJson struct {
	Range Range `json:"range,omitempty"`
}

type Timestamp struct {
	Gte string `json:"gte"`
}

type Match struct {
	//AgentHostnameKeyword string `json:"agent.hostname.keyword"`
	//https://stackoverflow.com/questions/18412126/golang-parse-a-json-with-dynamic-key
	hostFieldName map[string]string
}
type MatchJson struct {
	Match map[string]string `json:"match,omitempty"`
}

type Aggs struct {
	FieldExistsBool FieldExistsBool `json:"field_exists_bool"`
}

type FieldExistsBool struct {
	FEBFilters FEBFilters `json:"filters"`
}

type FEBFilters struct {
	OtherBucketKey string  `json:"other_bucket_key"`
	OtherBucket    bool    `json:"other_bucket"`
	Filters        Filters `json:"filters"`
}

type Filters struct {
	FieldExists FieldExists `json:"field_exists"`
}

type FieldExists struct {
	Exists Exists `json:"exists"`
}

type Exists struct {
	Field string `json:"field"`
}

type queryMatch struct {
	//https://stackoverflow.com/questions/18412126/golang-parse-a-json-with-dynamic-key
	//https://golang.org/pkg/encoding/json/#Marshal
	hostFieldName map[string]string
}

// Does this struct and its fields really need to be public?
type DbhitfResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore interface{}   `json:"max_score"`
		Hits     []interface{} `json:"hits"`
	} `json:"hits"`
	Aggregations struct {
		FieldExistsBool struct {
			Buckets struct {
				FieldExists struct {
					DocCount int `json:"doc_count"`
				} `json:"field_exists"`
				FieldDne struct {
					DocCount int `json:"doc_count"`
				} `json:"field_dne"`
			} `json:"buckets"`
		} `json:"field_exists_bool"`
	} `json:"aggregations"`
}
