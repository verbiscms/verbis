package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {

	tt := map[string]struct {
		depth    int
		traverse int
		want     int
	}{
		"Single":   {1, 0, 1},
		"Multiple": {3, 0, 3},
		"Traverse": {3, 1, 2},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Stack(test.depth, test.traverse)
			assert.Equal(t, test.want, len(got))
		})
	}
}

func TestFileStack_Lines(t *testing.T) {

	tt := map[string]struct {
		input []*FileStack
		want  []*FileLine
	}{
		"Single": {
			[]*FileStack{
				{Line: 1, Contents: "test\ntest"},
			},
			[]*FileLine{
				{Line: 2, Content: "test"},
			},
		},
		"Multiple": {
			[]*FileStack{
				{Line: 1, Contents: "test"},
				{Line: 2, Contents: "test\ntest"},
			},
			[]*FileLine{
				{Line: 1, Content: "test"},
				{Line: 2, Content: "test\ntest"},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			if len(test.input) == 0 {
				t.Fail()
			}
			got := test.input[0].Lines()
			for _, v := range got {
				for _, line := range test.want {
					assert.Equal(t, *line, *v)
				}
			}
		})
	}
}
