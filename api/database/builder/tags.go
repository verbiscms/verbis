package builder

import (
	"reflect"
	"strings"
)

const (
	tagName    = "sqlb"
	skipAll    = "skipAll"
	createTime = "autoCreateTime"
	updateTime = "autoUpdateTime"
)

func getTags(field reflect.StructField) (string, bool) {
	return field.Tag.Lookup(tagName)
}

func splitTags(str string) []string {
	return strings.Split(str, ",")
}

func isSkipped(tags []string) bool {
	for _, v := range tags {
		if v == skipAll {
			return true
		}
	}
	return false
}

func isAutoCreateTime(value string) bool {
	return strings.Contains(value, createTime)
}

func isAutoUpdateTime(value string) bool {
	return strings.Contains(value, updateTime)
}
