package valid_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/matchmove/valid"
)

const (
	jsonSchemaFile    = "test-data/sample-jsonschema.json"
	jsonSchemaHeader  = `"$ref": "http://json-schema.org/draft-04/schema#",`
	jsonSchemaPayload = `{
  "$ref": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "email": {
      "type": "string",
      "format": "email"
    },
    "id": {
      "type": "string",
      "pattern": "[A-Z]{11}"
    }
  }
}`
)

func TestCreateJSONSchema(t *testing.T) {
	s := valid.NewJSONSchema(jsonSchemaPayload)

	if s.Body == "" {
		t.Errorf("Post-processed should not be EMPTY")
	}

	if strings.Contains(s.Body, jsonSchemaHeader) {
		t.Errorf("Post-processed schema should not contain header, but found `%s`", jsonSchemaHeader)
	}
}

func TestCreateJSONSchemaFromFile(t *testing.T) {
	s := valid.NewJSONSchemaFromFile(jsonSchemaFile)

	if s.Body == "" {
		t.Errorf("Post-processed should not be EMPTY")
	}

	if strings.Contains(s.Body, jsonSchemaHeader) {
		t.Errorf("Post-processed schema should not contain header, but found `%s`", jsonSchemaHeader)
	}

	defer func() {
		errMsg := "Error in loading JSON schema file(invalid/file/path):open invalid/file/path: no such file or directory"
		if err := recover(); err != nil {
			if fmt.Sprintf("%v", err) != errMsg {
				t.Errorf("Expecting error `%s` to occur but received `%v`", errMsg, err)
			}
		}
	}()

	valid.NewJSONSchemaFromFile("invalid/file/path")
	t.Error("Expecting error to occur when loading an invalid file but nothing happened.")
}

func TestGetSchema(t *testing.T) {
	s := valid.NewJSONSchema("x")

	defer func() {
		errMsg := `Failed to create gojsonschema.NewSchema ERROR:invalid character 'x' looking for beginning of value`
		if err := recover(); err != nil {
			if fmt.Sprintf("%v", err) != errMsg {
				t.Errorf("Expecting error `%s` to occur but received `%v`", errMsg, err)
			}
		}
	}()

	s.Schema()
	t.Error("Expecting error to occur when loading an invalid file but nothing happened.")
}

func TestCompare(t *testing.T) {
	schema := valid.NewJSONSchemaFromFile(jsonSchemaFile)
	if r := schema.Compare(`{"email":"someone@email.com"}`); !r.Status {
		t.Errorf(
			"JSON-schema comparison must result to `true`; comparing:\n`%s`\nto\n`%s`",
			`{"email":"someone@email.com"}`,
			schema.Body,
		)
	}

	// Negative Tests

	if r := schema.Compare(""); "JSON.Body cannot be EMPTY" != r.Message {
		t.Errorf(
			"Expecting error `%s` when Comparing:\nEMPTY\nto\n`%s`\n but got: %s",
			"JSON.Body cannot be EMPTY",
			schema.Body,
			r.Message,
		)
	}

	var errMsg string

	errMsg = "email: Does not match format 'email'"
	if r := schema.Compare(`{"email":"invalid"}`); !strings.Contains(r.Message, errMsg) {
		t.Errorf(
			"Expecting error cotains `%s` when Comparing:\n%s\nto\n`%s`\n but got: `%s`",
			errMsg,
			`{"email":"invalid"}`,
			schema.Body,
			r.Message,
		)
	}

	errMsg = "Encountered schema.Validate ERROR:invalid character 'x' looking for beginning of value"
	if r := schema.Compare("x"); r.Message != errMsg {
		t.Errorf(
			"Expecting error `%s` when Comparing:\nx\nto\n`%s`\n but got: `%s`",
			errMsg,
			schema.Body,
			r.Message,
		)
	}
}
