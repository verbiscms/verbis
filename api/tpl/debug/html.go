package debug

import (
	"fmt"
)

var (
	name = ``
	css  = `
		<style>
			.pre.sf-dump {
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

type htmlWriter struct{}

func (h *htmlWriter) write(i interface{}) string {
	return h.write(i)
}

func (h *htmlWriter) writeDeep(i interface{}, name string) string {
	return h.name(name) + h.value(i)
}

func (h *htmlWriter) name(n string) string {
	return fmt.Sprintf(`<span class="sf-dump-name">%s</span>`, n)
}

func (h *htmlWriter) value(val interface{}) string {
	return fmt.Sprintf(`<span class="sf-dump-type">%T</span> <span class="sf-dump-value">%[1]v</span>`, val)
}
