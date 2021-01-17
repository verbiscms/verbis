package tpl

func (t *TplTestSuite) Test_BasePath() {
	oldBasePath := basePath
	defer func() {
		basePath = oldBasePath
	}()

	basePath = "test"
	tpl := "{{ basePath }}"
	t.RunT(tpl, "test")
}

func (t *TplTestSuite) Test_AdminPath() {
	oldAdminPath := adminPath
	defer func() {
		adminPath = oldAdminPath
	}()

	adminPath = "test"
	tpl := "{{ adminPath }}"
	t.RunT(tpl, "test")
}

func (t *TplTestSuite) Test_APIPath() {
	oldApiPath := apiPath
	defer func() {
		apiPath = oldApiPath
	}()

	apiPath = "test"
	tpl := "{{ apiPath }}"
	t.RunT(tpl, "test")
}

func (t *TplTestSuite) Test_ThemePath() {
	oldThemePath := themePath
	defer func() {
		themePath = oldThemePath
	}()

	themePath = "test"
	tpl := "{{ themePath }}"
	t.RunT(tpl, "test")
}

func (t *TplTestSuite) Test_UploadsPath() {
	oldUploadsPath := uploadsPath
	defer func() {
		uploadsPath = oldUploadsPath
	}()

	uploadsPath = "test"
	tpl := "{{ uploadsPath }}"
	t.RunT(tpl, "test")
}

func (t *TplTestSuite) Test_AssetsPath() {
	t.themeConfig.AssetsPath = "test"
	tpl := "{{ assetsPath }}"
	t.RunT(tpl, "test")
}

func (t *TplTestSuite) Test_StoragePath() {
	oldStoragePath := storagePath
	defer func() {
		storagePath = oldStoragePath
	}()

	storagePath = "test"
	tpl := "{{ storagePath }}"
	t.RunT(tpl, "test")
}

func (t *TplTestSuite) Test_TemplatesPath() {
	t.themeConfig.TemplateDir = "/dir"

	oldThemePath := themePath
	defer func() {
		themePath = oldThemePath
	}()

	themePath = "test"
	tpl := "{{ templatesPath }}"
	t.RunT(tpl, "test/dir")
}

func (t *TplTestSuite) Test_LayoutsPath() {
	t.themeConfig.LayoutDir = "/dir"

	oldThemePath := themePath
	defer func() {
		themePath = oldThemePath
	}()

	themePath = "test"
	tpl := "{{ layoutsPath }}"
	t.RunT(tpl, "test/dir")
}
