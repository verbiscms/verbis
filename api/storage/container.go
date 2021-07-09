package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/storage/internal"
	"github.com/graymeta/stow"
)

type container struct {
	*internal.Config
}

func (c *container) SetBucket(id string) error {
	const op = "Storage.SetBucket"

	if c.Options.StorageProvider == domain.StorageLocal {
		id = ""
	}

	container, err := c.Provider.Container(id)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error setting bucket", Operation: op, Err: err}
	}
	c.Bucket = container

	err = c.OptionsRepo.Update("storage_bucket", id)
	if err != nil {
		return &errors.Error{Code: errors.INTERNAL, Message: "Error updating options table with new bucket", Operation: op, Err: err}
	}

	return nil
}

func (c *container) ListBuckets() (domain.Buckets, error) {
	const op = "Container.ListBuckets"

	var buckets = make(domain.Buckets, 0)
	err := stow.WalkContainers(c.Provider, stow.NoPrefix, 100, func(c stow.Container, err error) error {
		if err != nil {
			return err
		}
		buckets = append(buckets, domain.Bucket{
			Id:   c.ID(),
			Name: c.Name(),
		})
		return nil
	})

	if err != nil {
		return nil, &errors.Error{Code: errors.INVALID, Message: "Error obtaining buckets", Operation: op, Err: err}
	}

	return nil, nil
}

func (c *container) CreateBucket(name string) error {
	const op = "Container.CreateBucket"

	_, err := c.Provider.CreateContainer(name)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error creating bucket: " + name, Operation: op, Err: err}
	}

	return nil
}

func (c *container) DeleteBucket(name string) error {
	const op = "Container.DeleteBucket"

	err := c.Provider.RemoveContainer(name)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error deleting bucket: " + name, Operation: op, Err: err}
	}

	return nil
}
