package api

import "github.com/ainsleyclark/verbis/api/http"

var (
	// DefaultParams represents the default params if
	// none were passed for the API.
	DefaultParams = http.Defaults{
		Page:           1,
		Limit:          15,
		OrderBy:        "created_at",
		OrderDirection: "desc",
	}
)
