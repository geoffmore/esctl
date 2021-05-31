package esutil

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	elastic7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

// TODO - rename esutil to util or figure out a better name for this package
// TODO - split this until its own types file
type esRequest interface {
	Do(context.Context, esapi.Transport) (*esapi.Response, error)
}

// TODO - remove this in the future if not needed
/*
// Check if the 'Format' field exists for a struct beneath a provided interface
func FormatExists(i interface{}) bool {
	return reflect.ValueOf(i).FieldByName("Format").IsValid()
}
*/
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
			return err
		}
	}
	fmt.Println(string(b))
	return nil
}

func GetResponse(r esRequest, c *elastic7.Client) (*esapi.Response, error) {
	res, err := r.Do(context.Background(), c.Transport)
	return res, err
}

// Generic function used to execute requests and return bytes
func RequestNew(r esRequest, c *elastic7.Client) ([]byte, error) {

	var b []byte

	res, err := r.Do(context.Background(), c.Transport)
	if err != nil {
		return b, err
	}
	// TODO - read first. Return response body
	// Check if status code is successful
	if !(res.StatusCode >= 200 && res.StatusCode < 300) {
		return b, fmt.Errorf("Status Code '%v' not in range 200-299. Exiting...\n", res.StatusCode)
	}

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return b, err
	}
	// https://tour.golang.org/moretypes/25
	// https://stackoverflow.com/questions/57740428/handling-errors-in-defer
	defer func() {
		err = res.Body.Close()
	}()

	if err != nil {
		return b, err
	}

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
