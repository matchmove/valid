package valid_test

import (
	"testing"

	"github.com/matchmove/valid"
)

func TestValidTypeMatch(t *testing.T) {
	foo := "bar"
	if result := valid.TypeMatch(foo, "string"); !result.Status {
		t.Error(result.Message)
	}

	if result := valid.TypeMatch(foo, "int"); result.Status {
		t.Errorf("Negative test should result to <false>:" + result.Message)
	}
}
