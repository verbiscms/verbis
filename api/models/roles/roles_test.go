package roles

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// RolesTestSuite defines the helper used for role
// testing.
type RolesTestSuite struct {
	test.DBSuite
}

// TestRoles
//
// Assert testing has begun.
func TestRoles(t *testing.T) {
	suite.Run(t, &RolesTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

// Setup
//
// A helper to obtain a mock roles database
// for testing.
func (t *RolesTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	if mf != nil {
		mf(t.Mock)
	}
	return New(t.DB)
}

var (
	// The default role used for testing.
	role = domain.Role{
		Id:          1,
		Name:        "Owner",
		Description: "Description",
	}
)
