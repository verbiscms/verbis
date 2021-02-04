// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/stretchr/testify/assert"
	"html/template"
	"testing"
)

var (
	ns = New(&deps.Deps{})
)

func TestNamespace_Debug(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  template.HTML
	}{
		"String": {
			"test",
			template.HTML("test\n"),
		},
		"Int": {
			1,
			template.HTML("1\n"),
		},
		"Slice": {
			[]interface{}{"test", 123},
			template.HTML("[test 123]\n"),
		},
		"Map": {
			map[string]interface{}{"test": 123},
			template.HTML("map[test:123]\n"),
		},
		"Struct": {
			struct{ Test string }{Test: "test"},
			template.HTML("{Test:test}\n"),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Debug(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Dump(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  interface{}
	}{
		"String": {
			"test",
			template.HTML(fmt.Sprintf(`%s<pre class="sf-dump">%s</pre>`, CSS, `"test"`)),
		},
		"Int": {
			1,
			template.HTML(fmt.Sprintf(`%s<pre class="sf-dump">%s</pre>`, CSS, `1`)),
		},
		"Slice": {
			[]interface{}{"test", 123},
			template.HTML(fmt.Sprintf(`%s<pre class="sf-dump">%s</pre>`, CSS, "[\n\t\"test\",\n\t123\n]")),
		},
		"Map": {
			map[string]interface{}{"test": 123},
			template.HTML(fmt.Sprintf(`%s<pre class="sf-dump">%s</pre>`, CSS, "{\n\t\"test\": 123\n}")),
		},
		"Struct": {
			struct{ Test string }{Test: "test"},
			template.HTML(fmt.Sprintf(`%s<pre class="sf-dump">%s</pre>`, CSS, "{\n\t\"Test\": \"test\"\n}")),
		},
		"Error": {
			map[string]interface{}{"foo": make(chan int)},
			"Templates.Dump: json: unsupported type: chan int",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := ns.Dump(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}
