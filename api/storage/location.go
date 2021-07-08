package storage

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/graymeta/stow"
)

type container struct {
	config
}

func (c *container) Set() {
	//	const op = "Storage.SetBucket"
	//
	//	if s.opts.StorageProvider == domain.StorageLocal {
	//		id = ""
	//	}
	//
	//	container, err := s.provider.Container(id)
	//	if err != nil {
	//		return &errors.Error{Code: errors.INVALID, Message: "Error setting bucket", Operation: op, Err: err}
	//	}
	//	s.bucket = container
	//
	//	err = s.optsRepo.Update("storage_bucket", id)
	//	if err != nil {
	//		return &errors.Error{Code: errors.INTERNAL, Message: "Error updating options table with new bucket", Operation: op, Err: err}
	//	}
}

func (c *container) List() (domain.Buckets, error) {
	const op = "Container.List"

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

func (c *container) Create(name string) error {
	const op = "Container.Create"

	_, err := c.Provider.CreateContainer(name)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error creating bucket: " + name, Operation: op, Err: err}
	}

	return nil
}

func (c *container) Delete(name string) error {
	const op = "Container.Delete"

	err := c.Provider.RemoveContainer(name)
	if err != nil {
		return &errors.Error{Code: errors.INVALID, Message: "Error deleting bucket: " + name, Operation: op, Err: err}
	}

	return nil
}
