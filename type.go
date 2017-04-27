package valid

import (
	"fmt"
	"reflect"
)

// TypeMatch checks if the `actual` variable has the same type as `expected`.
func TypeMatch(actual interface{}, expected string) Result {
	actualType := reflect.TypeOf(actual).String()

	if expected == actualType {
		return OkResult()
	}

	return FailResult(fmt.Sprintf("Expecting variable to be type <%s> instead of <%s>", expected, actualType))
}
