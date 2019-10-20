package esutil

import (
	"context"
	"encoding/json"
	"fmt"
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
	"strings"
)

type esRequest interface {
	Do(context.Context, esapi.Transport) (*esapi.Response, error)
}

// Deserialize a response into a generic interface
func Des(esRes *esapi.Response) (r map[string]interface{}, err error) {
	//statusCode := esRes.StatusCode
	//header := esRes.Header

	if err = json.NewDecoder(esRes.Body).Decode(&r); err != nil {
		return r, err
	}
	defer esRes.Body.Close()

	return r, err
}

// Display a generic interface as yaml
func MapToYamlish(r map[string]interface{}, s int) {
	var nestedMap map[string]interface{}
	spaces := s
	if false {
		fmt.Printf("%v", spaces)
	}
	for k, v := range r {
		switch v.(type) {
		default:
			// Did not find a map as a value
			fmt.Printf("%s%v: '%v'\n", strings.Repeat(" ", 2*(spaces)), k, v)
		case map[string]interface{}:
			// Found a map as a value
			fmt.Printf("%v:\n", k)
			// Explicitly create a map from that value, recurse
			nestedMap = v.(map[string]interface{})
			MapToYamlish(nestedMap, spaces+1)
			// I need a way to use multiple interface types and iterate over them
			// Maybe I add logic to figure out how to determine the type of the map
			// interface thing?
			//case map[int]interface{}:
			//	fmt.Printf("%v:\n", k)
			//	nestedMap = v.(map[int]interface{})
			//	mapToYamlish(nestedMap, spaces+1)
		}
	}
}

// Convert a generic interface of type map[string]interface{} to yaml bytes
func MapToYamlBytes(r map[string]interface{}) (b []byte, err error) {
	return b, err
}

// Generic function used to execute requests and print results
func Request(r esRequest, c *elastic7.Client) error {

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

// Check if the 'Format' field exists for a struct beneath a provided interface
func FormatExists(i interface{}) bool {
	return reflect.ValueOf(i).FieldByName("Format").IsValid()
}

// Convert bytes into a desired output format
func ParseBytes(b []byte, fmtExists bool, outputFmt string) (err error) {
	if outputFmt == "yaml" && !fmtExists {
		var iface map[string]interface{}
		err := json.Unmarshal(b, &iface)
		if err != nil {
			return err
		}
		b, err = yaml.Marshal(iface)
		if err != nil {
			// MapToYamlBytes needs to correctly handle generic interfaces without
			// recursion (unless said recursion relies on a pointer to a byte
			b, err = MapToYamlBytes(iface)
			if err != nil {
				return err
			}
		}
	}
	fmt.Println(string(b))
	return nil
}

// Generic function used to execute requests and return bytes
func RequestNew(r esRequest, c *elastic7.Client) ([]byte, error) {

	var b []byte

	res, err := r.Do(context.Background(), c.Transport)
	if err != nil {
		return b, err
	}
	if res.StatusCode != 200 {
		return b, fmt.Errorf("Status Code is %v rather than 200. Exiting...\n", res.StatusCode)
	}

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return b, err
	}
	defer res.Body.Close()

	return b, nil
}

// Attempt to change the 'Format' field of a struct if it exists.
// Return whether the field exists and could be set
func SetFormat(s reflect.Value, outputFmt string) bool {
	// https://blog.golang.org/laws-of-reflection
	var matches bool

	var fieldExists bool = s.FieldByName("Format").IsValid()
	var canSet bool = s.CanSet()

	if s.Kind() == reflect.Struct {
		if fieldExists {
			// Attempt to change output format
			switch outputFmt {
			case "json":
				matches = true
			case "yaml":
				matches = true
			}
			if matches && canSet {
				s.FieldByName("Format").SetString(outputFmt)
			}
		}
	}
	return fieldExists && canSet
}
