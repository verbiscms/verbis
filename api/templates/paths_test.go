package templates

import (
	"testing"
)

func Test_BasePath(t *testing.T) {
	f := newTestSuite()
	oldBasePath := basePath
	defer func() {
		basePath = oldBasePath
	}()

	basePath = "test"
	tmpl := "{{ basePath }}"
	runt(t, f, tmpl, "test")
}

func Test_AdminPath(t *testing.T) {
	f := newTestSuite()
	oldAdminPath := adminPath
	defer func() {
		adminPath = oldAdminPath
	}()

	adminPath = "test"
	tmpl := "{{ adminPath }}"
	runt(t, f, tmpl, "test")
}

func Test_APIPath(t *testing.T) {
	f := newTestSuite()
	oldApiPath := apiPath
	defer func() {
		apiPath = oldApiPath
	}()

	apiPath = "test"
	tmpl := "{{ apiPath }}"
	runt(t, f, tmpl, "test")
}

func Test_ThemePath(t *testing.T) {
	f := newTestSuite()
	oldThemePath := tmplThemePath
	defer func() {
		tmplThemePath = oldThemePath
	}()

	tmplThemePath = "test"
	tmpl := "{{ themePath }}"
	runt(t, f, tmpl, "test")
}

func Test_UploadsPath(t *testing.T) {
	f := newTestSuite()
	oldUploadsPath := uploadsPath
	defer func() {
		uploadsPath = oldUploadsPath
	}()

	uploadsPath = "test"
	tmpl := "{{ uploadsPath }}"
	runt(t, f, tmpl, "test")
}

func Test_AssetsPath(t *testing.T) {
	f := newTestSuite()
	f.themeConfig.AssetsPath = "test"
	tmpl := "{{ assetsPath }}"
	runt(t, f, tmpl, "test")
}

func Test_StoragePath(t *testing.T) {
	f := newTestSuite()
	oldStoragePath := storagePath
	defer func() {
		storagePath = oldStoragePath
	}()

	storagePath = "test"
	tmpl := "{{ storagePath }}"
	runt(t, f, tmpl, "test")
}

func Test_TemplatesPath(t *testing.T) {
	f := newTestSuite()
	f.themeConfig.TemplateDir = "/dir"

	oldThemePath := tmplThemePath
	defer func() {
		tmplThemePath = oldThemePath
	}()

	tmplThemePath = "test"
	tmpl := "{{ templatesPath }}"
	runt(t, f, tmpl, "test/dir")
}

func Test_LayoutsPath(t *testing.T) {
	f := newTestSuite()
	f.themeConfig.LayoutDir = "/dir"

	oldThemePath := tmplThemePath
	defer func() {
		tmplThemePath = oldThemePath
	}()

	tmplThemePath = "test"
	tmpl := "{{ layoutsPath }}"
	runt(t, f, tmpl, "test/dir")
}
