package escmd

import (
	"context"
	"fmt"
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io/ioutil"
)

// Generic interface for requests
type esRequest interface {
	Do(context.Context, esapi.Transport) (*esapi.Response, error)
}

// Generic function used to execute requests
func request(r esRequest, c *elastic7.Client) error {
	res, err := r.Do(context.Background(), c.Transport)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("Status Code is %v rather than 200. Exiting...\n", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	fmt.Printf("%+v\n", string(b))

	return nil

}
