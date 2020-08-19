package esutil

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/geoffmore/esctl/pkg/opts"
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

	// Check if status code is successful
	if !(res.StatusCode >= 200 && res.StatusCode < 300) {
		return fmt.Errorf("Status Code '%v' not in range 200-299. Exiting...\n", res.StatusCode)
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
	// Check if status code is successful
	if !(res.StatusCode >= 200 && res.StatusCode < 300) {
		return b, fmt.Errorf("Status Code '%v' not in range 200-299. Exiting...\n", res.StatusCode)
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

// https://golangcode.com/check-if-element-exists-in-slice/
// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// Take json. Return a reader
func JSONToReader(intf interface{}) (*bytes.Reader, error) {
	b, err := json.Marshal(intf)
	return bytes.NewReader(b), err
}

func FilenameToReader(name string) (*bytes.Reader, error) {

	var r *bytes.Reader

	// Handle stdin
	if name == "-" {
		return r, fmt.Errorf("Reading from stdin is not yet supported")
	} else {
		// Open the file and read its contents
		b, err := ioutil.ReadFile(name)
		if err != nil {
			return r, err
		}
		// Create a *io.Reader from bytes
		r = bytes.NewReader(b)
	}

	return r, nil
}

// Attempt to change the value of an arbitrary boolean type field
// Return whether the field exists and could be set
func SetBool(s reflect.Value, field string, a bool) bool {

	var fieldExists bool = s.FieldByName(field).IsValid()
	var canSet bool = s.CanSet()

	if s.Kind() == reflect.Struct {
		if fieldExists && canSet {
			// Attempt to change value
			s.FieldByName(field).SetBool(a)
		}
	}
	return fieldExists && canSet
}

// Attempt to change the value of an arbitrary string type field
// Return whether the field exists and could be set
func SetString(s reflect.Value, field string, a string) bool {

	var fieldExists bool = s.FieldByName(field).IsValid()
	var canSet bool = s.CanSet()

	if s.Kind() == reflect.Struct {
		if fieldExists && canSet {
			// Attempt to change value
			s.FieldByName(field).SetString(a)
		}
	}
	return fieldExists && canSet
}

// Attempt to change the value of an arbitrary int64 type field
// Return whether the field exists and could be set
func SetInt(s reflect.Value, field string, a int64) bool {

	var fieldExists bool = s.FieldByName(field).IsValid()
	var canSet bool = s.CanSet()

	if s.Kind() == reflect.Struct {
		if fieldExists && canSet {
			// Attempt to change value
			s.FieldByName(field).SetInt(a)
		}
	}
	return fieldExists && canSet
}

// Attempt to set all fields contained in opts.CommandOptions according to the
// map opts.CmdsToFieldNames
func SetAllCmdOpts(v reflect.Value, c *opts.CommandOptions) map[string]bool {

	cmdOpts := opts.CmdsToFieldNames

	// https://stackoverflow.com/questions/18926304
	cv := reflect.ValueOf(c).Elem()
	var changedFields map[string]bool = make(map[string]bool)

	for cmdFieldName, structFieldName := range cmdOpts {
		val := v.FieldByName(structFieldName)
		if val.IsValid() {
			// Type lookup is necessary here for the switch
			switch t := val.Type().String(); t {
			case "string":
				changedFields[structFieldName] = SetString(v, structFieldName, cv.FieldByName(cmdFieldName).String())
			//case "int":
			// reflect's SetInt() expects int64
			case "int64":
				changedFields[structFieldName] = SetInt(v, structFieldName, cv.FieldByName(cmdFieldName).Int())
			case "bool":
				changedFields[structFieldName] = SetBool(v, structFieldName, cv.FieldByName(cmdFieldName).Bool())
			}
		} else {
			// Handle the case where the field doesn't exist in the struct
			changedFields[structFieldName] = false
		}
	}
	return changedFields
}
