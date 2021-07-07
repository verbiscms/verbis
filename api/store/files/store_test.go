package files

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/store/config"
	"github.com/ainsleyclark/verbis/api/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// FilesTestSuite defines the helper used for file
// testing.
type FilesTestSuite struct {
	test.DBSuite
}

// TestFiles
//
// Assert testing has begun.
func TestFiles(t *testing.T) {
	suite.Run(t, &FilesTestSuite{
		DBSuite: test.NewDBSuite(t),
	})
}

const (
	// The default file ID used for testing.
	fileID = "1"
)

// Setup
//
// A helper to obtain a mock files database
// for testing.
func (t *FilesTestSuite) Setup(mf func(m sqlmock.Sqlmock)) *Store {
	t.Reset()
	if mf != nil {
		mf(t.Mock)
	}
	return New(&config.Config{
		Driver: t.Driver,
	})
}

var (
	// The default file used for testing.
	file = domain.File{
		Id:       1,
		Url:      "/uploads/2020/01/file.jpg",
		Name:     "file.jpg",
		Path:     "uploads/2020/01",
		Provider: domain.StorageLocal,
	}
	// The default categories used for testing.
	files = domain.Files{
		{
			Id:       1,
			Url:      "/uploads/2020/01/file.jpg",
			Name:     "file.jpg",
			Path:     "uploads/2020/01",
			Provider: domain.StorageLocal,
		},
		{
			Id:       2,
			Url:      "/uploads/2020/01/file-2.jpg",
			Name:     "file-2.jpg",
			Path:     "uploads/2020/01",
			Provider: domain.StorageLocal,
		},
	}
)
