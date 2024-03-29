package files

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"github.com/verbiscms/verbis/api/domain"
	"github.com/verbiscms/verbis/api/store/config"
	"github.com/verbiscms/verbis/api/test"
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
		ID:       1,
		URL:      "/uploads/2020/01/file.jpg",
		Name:     "file.jpg",
		BucketID: "uploads/2020/01/file.jpg",
		Provider: domain.StorageLocal,
	}
	// The default files used for testing.
	files = domain.Files{
		{
			ID:       1,
			URL:      "/uploads/2020/01/file.jpg",
			Name:     "file.jpg",
			BucketID: "uploads/2020/01/file.jpg",
			Provider: domain.StorageLocal,
		},
		{
			ID:       2,
			URL:      "/uploads/2020/01/file-2.jpg",
			Name:     "file-2.jpg",
			BucketID: "uploads/2020/01/file.jpg",
			Provider: domain.StorageLocal,
		},
	}
)
