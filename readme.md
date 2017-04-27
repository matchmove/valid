![Codeship CI Status](https://codeship.com/projects/f00d5830-0afd-0135-7622-4abc4c11ded6/status?branch=master)

    import rest "gopkg.in/matchmove/rest.v2"

# valid
--
    import "github.com/matchmove/valid"


## Usage

```go
const (
	//REGEXJSONSchemaHeader regex format for JSON Schema Header
	REGEXJSONSchemaHeader = `"\$ref": "http:\/\/json-schema\.org\/draft\-[0-9]{2}\/schema#",\n?`
)
```

#### type JSONSchema

```go
type JSONSchema struct {
	Body string
}
```

JSONSchema represents the json-schema standards used for documentation

#### func  NewJSONSchema

```go
func NewJSONSchema(s string) JSONSchema
```
NewJSONSchema creates a json-schema from string stripping the header

#### func  NewJSONSchemaFromFile

```go
func NewJSONSchemaFromFile(path string) JSONSchema
```
NewJSONSchemaFromFile loads a json-schema from a file

#### func (JSONSchema) Compare

```go
func (data JSONSchema) Compare(schema JSONSchema) Result
```
Compare compares a json to a json-schema

#### func (JSONSchema) Loader

```go
func (data JSONSchema) Loader() gojsonschema.JSONLoader
```
Loader returns the JSONLoader of the the string Body

#### func (JSONSchema) Schema

```go
func (data JSONSchema) Schema() *gojsonschema.Schema
```
Schema creates a schema object from string

#### type Result

```go
type Result struct {
	Message string
	Status  bool
}
```

Result represents a testing result

#### func  FailResult

```go
func FailResult(message string) Result
```
FailResult creates a failed result with message

#### func  OkResult

```go
func OkResult() Result
```
OkResult creates a successful result

#### func  TypeMatch

```go
func TypeMatch(actual interface{}, expected string) Result
```
TypeMatch checks if the `actual` variable has the same type as `expected`.

#### func (Result) PrintIfFail

```go
func (r Result) PrintIfFail(fn func(...interface{}))
```
PrintIfFail executes a print function, like t.Error, when status is false
