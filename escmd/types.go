package escmd

import (
	"context"
	"fmt"
	elastic7 "github.com/geoffmore/go-elasticsearch/v7"
	"github.com/geoffmore/go-elasticsearch/v7/esapi"
	"io/ioutil"
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

// Generic interface for requests
type esRequest interface {
	Do(context.Context, esapi.Transport) (*esapi.Response, error)
}

// Generic function used to execute requests
func request(r esRequest, c *elastic7.Client) error {
	// Can request take a format json, yaml as an argument?
	// If try to serialize using map[string]interface{} my app immediately
	// becomes less performant, but I wouldn't have to dynamically generate
	// structs for each different type of response
	// Error in request execution
	res, resErr := r.Do(context.Background(), c.Transport)
	if resErr != nil {
		return resErr
	}

	// Read response body
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// Error in http response
	if res.StatusCode != 200 {
		return fmt.Errorf("Status Code is %v rather than 200.\n Printed error is: \n%v\n. Exiting...\n", res.StatusCode, string(b))
	}

	fmt.Printf("%+v\n", string(b))
	return nil

}
