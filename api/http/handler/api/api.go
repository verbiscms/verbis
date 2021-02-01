package api

import (
	"github.com/ainsleyclark/verbis/api/helpers/params"
)

var (
	// DefaultParams represents the default params if
	// none were passed for the API.
	DefaultParams = params.Defaults{
		Page:           1,
		Limit:          15,
		OrderBy:        "created_at",
		OrderDirection: "desc",
	}
)
