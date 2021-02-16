 // Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

 import (
	 "fmt"
	 "github.com/sanity-io/litter"
	 "html/template"
 )

const (
	// CSS describes the styling for dumps to be sent back
	// to the front end.
	CSS = `
		<style>
			.sf-dump {
				background-color: #18171B;
				color: #FF8400;
				line-height: 1.4em;
				font: 12px Menlo, Monaco, Consolas, monospace;
				word-wrap: break-word;
				white-space: pre-wrap;
				position: relative;
				z-index: 99999;
				word-break: break-all;
				white-space: pre-wrap;
				padding: 5px;
				overflow: initial !important;
			}
			.pre.sf-dump .sf-dump-name {
				color: #fff;
			}
			.pre.sf-dump .sf-dump-public {
				color: #fff;
			}
			.pre.sf-dump .sf-dump-value {
				color: #56DB3A;
				font-weight: bold;
			}
		</style>`
)

// Debug
//
// Returns a pretty print of the interface passed.
// This function is a shortcut for fmt.Sprintf
//
// Example: {{ debug .Post }}
func (ns *Namespace) Debug(i interface{}) template.HTML {
	return template.HTML(fmt.Sprintf("%+v\n", i))
}

// Dump
//
// Uses the litter package to dump the passed interface.
// inside a div with CSS attached.
//
// Returns errors.TEMPLATE if the marshal failed.
//
// Example: {{ dump .Post }}
func (ns *Namespace) Dump(i interface{}) (template.HTML, error) {
	dump := litter.Sdump(i)
	el := fmt.Sprintf(`%s<pre class="sf-dump">%s</pre>`, CSS, dump)
	return template.HTML(el), nil
}
