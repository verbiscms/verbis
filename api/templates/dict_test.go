package templates

import (
	"testing"
)

func TestDic(t *testing.T) {
	f := newTestSuite()
	tpl := `{{ dict "test" 123 }}`
	runt(t, f, tpl, "map[test:123]")
}

func TestDic_Invalid(t *testing.T) {
	f := newTestSuite()
	tpl := `{{ dict "test" }}`
	runt(t, f, tpl, "")
}

func TestDic_InvalidString(t *testing.T) {
	f := newTestSuite()
	tpl := `{{ dict 2 3 }}`
	runt(t, f, tpl, "")
}