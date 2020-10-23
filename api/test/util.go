package test

import "fmt"

func Format(expect interface{}, got interface{}) string {
	return fmt.Sprintf("Test failed, expecting %v, got %v", expect, got)
}
