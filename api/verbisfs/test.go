package verbisfs

import (
	"embed"
	"github.com/ainsleyclark/verbis/admin"
	"github.com/ainsleyclark/verbis/api"
	"github.com/ainsleyclark/verbis/api/helpers/mime"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"github.com/ainsleyclark/verbis/api/www"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type FS interface {
	Open(name string) (fs.File, error)
	ReadFile(name string) ([]byte, error)
	ReadDir(name string) ([]fs.DirEntry, error)
	HTTP(path, root string) gin.HandlerFunc
}

type FileSystem struct {
	SPA FS
	Web FS
}

const (
	SpaDistFolder = "dist"
)

func New() *FileSystem {
	p := paths.Get()
	if api.SuperAdmin {
		return &FileSystem{
			SPA: &osFS{path: p.Admin},
			Web: &osFS{path: p.Web},
		}
	}
	return &FileSystem{
		SPA: &embedFS{fs: admin.SPA, prefix: SpaDistFolder},
		Web: &embedFS{fs: www.Web, prefix: ""},
	}
}

type embedFS struct {
	fs     embed.FS
	prefix string
}

func (s *embedFS) Open(name string) (fs.File, error) {
	return s.fs.Open(s.prefix + name)
}

func (s *embedFS) ReadFile(name string) ([]byte, error) {
	return s.fs.ReadFile(s.prefix + name)
}

func (s *embedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return s.fs.ReadDir(s.prefix + name)
}

func (s *embedFS) HTTP(path, root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file := strings.ReplaceAll(ctx.Request.URL.String(), path, "")

		bytes, err := s.ReadFile(filepath.Join(root, file))
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.Data(http.StatusOK, mime.TypeByExtension(file), bytes)
	}
}

type osFS struct {
	path string
}

func (s *osFS) Open(name string) (fs.File, error) {
	return os.Open(filepath.Join(s.path, name))
}

func (s *osFS) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(filepath.Join(s.path, name))
}

func (s *osFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(filepath.Join(s.path, name))
}

func (s *osFS) HTTP(path, root string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file := strings.ReplaceAll(ctx.Request.URL.String(), path, "")

		bytes, err := s.ReadFile(filepath.Join(root, file))
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.Data(http.StatusOK, mime.TypeByExtension(file), bytes)
	}
}
