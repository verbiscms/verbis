package recovery

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetError(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  *errors.Error
	}{
		"Non Pointer": {
			errors.Error{Code: errors.INTERNAL, Message: "test", Operation: "op"},
			&errors.Error{Code: errors.INTERNAL, Message: "test", Operation: "op"},
		},
		"Pointer": {
			&errors.Error{Code: errors.INTERNAL, Message: "test", Operation: "op"},
			&errors.Error{Code: errors.INTERNAL, Message: "test", Operation: "op"},
		},
		"Standard Error": {
			fmt.Errorf("error"),
			&errors.Error{Code: errors.TEMPLATE, Message: "error", Operation: "", Err: fmt.Errorf("error")},
		},
		"Pointer Standard Error": {
			nil,
			&errors.Error{Code: errors.TEMPLATE, Message: "Internal Verbis error, please report", Operation: "", Err: fmt.Errorf("internal verbis error")},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := getError(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
