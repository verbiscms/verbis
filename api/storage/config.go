package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/store/files"
	"github.com/ainsleyclark/verbis/api/store/options"
	"github.com/graymeta/stow"
	_ "github.com/graymeta/stow/azure"
	"github.com/graymeta/stow/google"
	_ "github.com/graymeta/stow/google"
	"github.com/graymeta/stow/local"
	"github.com/graymeta/stow/s3"
	_ "github.com/graymeta/stow/s3"
	"io/ioutil"
)

type config struct {
	ProviderName domain.StorageProvider
	Environment  *environment.Env
	OptionsRepo  options.Repository
	FilesRepo    files.Repository
	Options      *domain.Options
	Paths        paths.Paths
	Provider     stow.Location
	Bucket       stow.Container
}

var (
	// memory of ioutil
	gcpJson *string
)

func (c *config) GetProvider(provider domain.StorageProvider) (stow.Location, error) {
	var (
		cont stow.Location
		err  error
	)

	switch provider {
	case domain.StorageLocal:
		cont, err = stow.Dial(local.Kind, stow.ConfigMap{
			local.ConfigKeyPath: c.Paths.Storage,
		})
	case domain.StorageAWS:
		cont, err = stow.Dial(s3.Kind, stow.ConfigMap{
			s3.ConfigAccessKeyID: c.Environment.AWSAccessKey,
			s3.ConfigSecretKey:   c.Environment.AWSSecret,
		})
	case domain.StorageGCP:
		json, err := ioutil.ReadFile(c.Environment.GCPJson)
		if err != nil {
			return nil, err
		}
		cont, err = stow.Dial(google.Kind, stow.ConfigMap{
			google.ConfigJSON:      string(json),
			google.ConfigProjectId: c.Environment.GCPProjectId,
		})
	}

	return cont, err
}

func (c *config) GetBucket(file domain.File) (stow.Container, error) {
	provider, err := c.GetProvider(file.Provider)
	if err != nil {
		// TODO varf( err etc
		return nil, err
	}

	bucket, err := provider.Container(file.Bucket)
	if err != nil {
		// TODO varf( err etc
		return nil, err
	}

	return bucket, nil
}

func (c *config) SetLocation(location stow.Location) {
	c.Provider = location
	//c.ProviderName = location.
}
