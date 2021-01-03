package location

import (
	"github.com/ainsleyclark/verbis/api/cache"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

type LocationTestSuite struct {
	suite.Suite
	Path string
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(LocationTestSuite))
}

func (t *LocationTestSuite) BeforeTest(suiteName, testName string) {
	cache.Init()

	err := logger.Init(config.Configuration{})
	log.SetOutput(ioutil.Discard)
	t.NoError(err)

	wd, err := os.Getwd()
	t.NoError(err)
	t.Path = filepath.Join(filepath.Dir(wd)+"./..") + "/test/testdata/fields"
}

func TestNewLocation(t *testing.T) {
	oldStoragePath := storagePath
	defer func() {
		storagePath = oldStoragePath
	}()
	storagePath = "test"

	assert.Equal(t, &Location{JsonPath: "test/fields"}, NewLocation())
}

func (t *LocationTestSuite) TestLocation_GetLayout() {

	tt := map[string]struct {
		cacheable bool
		jsonPath  string
		want      interface{}
	}{
		"Bad Path": {
			cacheable: false,
			jsonPath:  "wrongval",
			want:      []domain.FieldGroup{},
		},
		"Not Cached": {
			cacheable: false,
			jsonPath:  "/test-get-layout",
			want:      []domain.FieldGroup{{Title: "title"}},
		},
		"Cacheable Nil": {
			cacheable: true,
			jsonPath:  "/test-get-layout",
			want:      []domain.FieldGroup{{Title: "title"}},
		},
		"Cacheable": {
			cacheable: true,
			jsonPath:  "/test-get-layout",
			want:      []domain.FieldGroup{{Title: "title"}},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			l := &Location{JsonPath: t.Path + test.jsonPath}
			t.Equal(test.want, l.GetLayout(domain.Post{}, domain.User{}, &domain.Category{}, test.cacheable))
		})
	}
}

func (t *LocationTestSuite) TestLocation_GroupResolver() {

	r := "resource"
	uu := uuid.New()

	tt := map[string]struct {
		post     domain.Post
		author   domain.User
		category *domain.Category
		groups   []domain.FieldGroup
		want     interface{}
	}{
		"None": {
			want: []domain.FieldGroup{},
		},
		"Already Added": {
			post: domain.Post{Id: 1, Title: "title", Status: "published"},
			groups: []domain.FieldGroup{
				{Title: "status", UUID: uu},
				{Title: "status", UUID: uu},
			},
			want: []domain.FieldGroup{{Title: "status", UUID: uu}},
		},
		"Status": {
			post: domain.Post{Id: 1, Title: "title", Status: "published"},
			groups: []domain.FieldGroup{
				{
					Title: "status",
					Locations: [][]domain.FieldLocation{
						{{Param: "status", Operator: "==", Value: "published"}},
					},
				},
			},
			want: []domain.FieldGroup{{Title: "status"}},
		},
		"Post": {
			post: domain.Post{Id: 1},
			groups: []domain.FieldGroup{
				{
					Title: "post",
					Locations: [][]domain.FieldLocation{
						{{Param: "post", Operator: "==", Value: "1"}},
					},
				},
			},
			want: []domain.FieldGroup{{Title: "post"}},
		},
		"Page Template": {
			post: domain.Post{PageTemplate: "template"},
			groups: []domain.FieldGroup{
				{
					Title: "post",
					Locations: [][]domain.FieldLocation{
						{{Param: "page_template", Operator: "==", Value: "template"}},
					},
				},
			},
			want: []domain.FieldGroup{{Title: "post"}},
		},
		"Layout": {
			post: domain.Post{PageLayout: "layout"},
			groups: []domain.FieldGroup{
				{
					Title: "post",
					Locations: [][]domain.FieldLocation{
						{{Param: "page_layout", Operator: "==", Value: "layout"}},
					},
				},
			},
			want: []domain.FieldGroup{{Title: "post"}},
		},
		"Resource": {
			post: domain.Post{Resource: &r},
			groups: []domain.FieldGroup{
				{
					Title: "post",
					Locations: [][]domain.FieldLocation{
						{{Param: "resource", Operator: "==", Value: r}},
					},
				},
			},
			want: []domain.FieldGroup{{Title: "post"}},
		},
		"Nil Resource": {
			post: domain.Post{Resource: nil},
			groups: []domain.FieldGroup{
				{
					Title: "post",
					Locations: [][]domain.FieldLocation{
						{{Param: "resource", Operator: "==", Value: "false"}},
					},
				},
			},
			want: []domain.FieldGroup{},
		},
		"Category": {
			category: &domain.Category{Id: 1},
			groups: []domain.FieldGroup{
				{
					Title: "category",
					Locations: [][]domain.FieldLocation{
						{{Param: "category", Operator: "==", Value: "1"}},
					},
				},
			},
			want: []domain.FieldGroup{{Title: "category"}},
		},
		"Nil Category": {
			category: nil,
			groups: []domain.FieldGroup{
				{
					Title: "category",
					Locations: [][]domain.FieldLocation{
						{{Param: "category", Operator: "==", Value: "false"}},
					},
				},
			},
			want: []domain.FieldGroup{},
		},
		"Author": {
			post: domain.Post{UserId: 1},
			groups: []domain.FieldGroup{
				{
					Title: "post",
					Locations: [][]domain.FieldLocation{
						{{Param: "author", Operator: "==", Value: "1"}},
					},
				},
			},
			want: []domain.FieldGroup{{Title: "post"}},
		},
		"Role": {
			author: domain.User{
				UserPart: domain.UserPart{
					Role: domain.UserRole{
						Id: 1,
					},
				},
			},
			groups: []domain.FieldGroup{
				{
					Title: "role",
					Locations: [][]domain.FieldLocation{
						{{Param: "role", Operator: "==", Value: "1"}},
					},
				},
			},
			want: []domain.FieldGroup{{Title: "role"}},
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			l := &Location{Groups: test.groups}
			t.Equal(test.want, l.groupResolver(test.post, test.author, test.category))
		})
	}
}

func (t *LocationTestSuite) TestLocation_fieldGroupWalker() {

	testPath := "/test-field-groups/"

	id, err := uuid.Parse("6a4d7442-1020-490f-a3e2-436f9135bc24")
	t.NoError(err)

	// For bad path
	err = os.Chmod(t.Path+testPath+"open-error/json.txt", 000)
	t.NoError(err)
	defer func() {
		err = os.Chmod(t.Path+testPath+"open-error/json.txt", 777)
		t.NoError(err)
	}()

	tt := map[string]struct {
		path string
		want interface{}
	}{
		"Success": {
			path: testPath + "/success",
			want: []domain.FieldGroup{{UUID: id, Title: "Title"}, {UUID: id, Title: "Title"}},
		},
		"Bad Path": {
			path: testPath + "/wrongval",
			want: "no such file or directory",
		},
		"Unmarshal Error": {
			path: testPath + "/unmarshal",
			want: "json: cannot unmarshal number into Go struct field FieldGroup.title of type string",
		},
		"Open Error": {
			path: testPath + "/open-error",
			want: "permission denied",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {

			l := &Location{
				JsonPath: t.Path + test.path,
			}
			got, err := l.fieldGroupWalker()

			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

func (t *LocationTestSuite) Test_CheckLocation() {

	tt := map[string]struct {
		check    string
		location domain.FieldLocation
		want     bool
	}{
		"Equal Match": {
			check:    "val",
			location: domain.FieldLocation{Operator: "==", Value: "val"},
			want:     true,
		},
		"Equal Not Match": {
			check:    "val",
			location: domain.FieldLocation{Operator: "==", Value: "wrongval"},
			want:     false,
		},
		"Not Equal Match": {
			check:    "val",
			location: domain.FieldLocation{Operator: "!=", Value: "val"},
			want:     false,
		},
		"Not Equal Not Match": {
			check:    "val",
			location: domain.FieldLocation{Operator: "!=", Value: "wrongval"},
			want:     true,
		},
		"Wrong Operator": {
			check:    "val",
			location: domain.FieldLocation{Operator: "wrong", Value: "val"},
			want:     false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, checkLocation(test.check, test.location))
		})
	}
}

func (t *LocationTestSuite) Test_CheckMatch() {

	tt := map[string]struct {
		matches []bool
		want    bool
	}{
		"Matches": {
			matches: []bool{true, false, true, false},
			want:    false,
		},
		"Not Matched": {
			matches: []bool{true, true, true, true},
			want:    true,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, checkMatch(test.matches))
		})
	}
}

func (t *LocationTestSuite) Test_HasBeenAdded() {

	key := uuid.New()

	tt := map[string]struct {
		fg   []domain.FieldGroup
		key  string
		want bool
	}{
		"Added": {
			fg:   []domain.FieldGroup{{UUID: key}},
			key:  key.String(),
			want: true,
		},
		"Not Added": {
			fg:   []domain.FieldGroup{{UUID: uuid.New()}},
			key:  key.String(),
			want: false,
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			t.Equal(test.want, hasBeenAdded(test.key, test.fg))
		})
	}
}
