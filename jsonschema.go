package valid

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/xeipuuv/gojsonschema"
)

// JSONSchema represents the json-schema standards used for documentation
type JSONSchema struct {
	Body string
}

const (
	//REGEXJSONSchemaHeader regex format for JSON Schema Header
	REGEXJSONSchemaHeader = `"\$ref": "http:\/\/json-schema\.org\/draft\-[0-9]{2}\/schema#",\n?`
)

// NewJSONSchemaFromFile loads a json-schema from a file
func NewJSONSchemaFromFile(path string) JSONSchema {

	var (
		buff []byte
		err  error
	)

	if buff, err = ioutil.ReadFile(path); err != nil {
		panic(fmt.Sprint("Error in loading JSON schema file(", path, "):", err))
	}

	return NewJSONSchema(string(buff))
}

// NewJSONSchema creates a json-schema from string stripping the header
func NewJSONSchema(s string) JSONSchema {
	reg, _ := regexp.Compile(REGEXJSONSchemaHeader)
	return JSONSchema{reg.ReplaceAllString(s, "")}
}

// Loader returns the JSONLoader of the the string Body
func (data JSONSchema) Loader() gojsonschema.JSONLoader {
	return gojsonschema.NewStringLoader(data.Body)
}

// Schema creates a schema object from string
func (data JSONSchema) Schema() *gojsonschema.Schema {
	var (
		schema *gojsonschema.Schema
		err    error
	)
	if schema, err = gojsonschema.NewSchema(data.Loader()); err != nil {
		panic(fmt.Sprint("Failed to create gojsonschema.NewSchema ERROR:", err))
	}

	return schema
}

// Compare compares a json to a json-schema
func (data JSONSchema) Compare(schema JSONSchema) Result {

	if data.Body == "" {
		return FailResult("JSON.Body cannot be EMPTY")
	}

	var (
		result *gojsonschema.Result
		err    error
	)

	if result, err = schema.Schema().Validate(data.Loader()); err != nil {
		return FailResult(fmt.Sprint("Encountered schema.Validate ERROR:", err))
	}

	if !result.Valid() {
		errors := ""
		for _, err := range result.Errors() {
			// Err implements the ResultError interface
			errors = errors + fmt.Sprintf("\n - %s", err)
		}
		return FailResult(fmt.Sprintf("Multiple schema errors found:%s", errors))
	}

	return OkResult()
}
