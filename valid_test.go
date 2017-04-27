package valid_test

import (
	"fmt"
	"testing"

	"github.com/matchmove/valid"
)

func TestPrintIfFail(t *testing.T) {
	valid.OkResult().PrintIfFail(t.Error)

	buff := ""

	printer := func(a ...interface{}) {
		buff = fmt.Sprint(a)
	}

	if valid.FailResult("Message").PrintIfFail(printer); "[Message]" != buff {
		t.Errorf(
			"Error should be printed and go '[Message]', instead got '%s'",
			buff,
		)
	}
}
