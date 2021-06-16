package static

import (
	"embed"
	"io/fs"
	"os"
)

type FS interface {
	Open(name string) (fs.File, error)
	ReadFile(name string) ([]byte, error)
	ReadDir(name string) ([]fs.DirEntry, error)
}

type ProductionFS struct {
	fs embed.FS
}

func New(path string) FS {

	return nil
}

func (s *ProductionFS) Open(name string) (fs.File, error) {
	return s.fs.Open(name)
}

func (s *ProductionFS) ReadFile(name string) ([]byte, error) {
	return s.fs.ReadFile(name)
}

func (s *ProductionFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return s.fs.ReadDir(name)
}

type DevFS struct {
	path string
}

func (s *DevFS) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func (s *DevFS) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (s *DevFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(name)
}
