package fields

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/logger"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"testing"
)

type FieldTestSuite struct {
	suite.Suite
}

type noStringer struct{}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(FieldTestSuite))
}

func (t *FieldTestSuite) BeforeTest(suiteName, testName string) {
	err := logger.Init(config.Configuration{})
	log.SetOutput(ioutil.Discard)
	t.NoError(err)
}

func (t *FieldTestSuite) GetMockService(fields []domain.PostField, fnc func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)) *Service {
	fieldsMock := &mocks.FieldsRepository{}
	categoryMock := &mocks.CategoryRepository{}

	fnc(fieldsMock, categoryMock)

	s := t.GetService(fields)
	s.store = &models.Store{
		Categories: categoryMock,
		Fields:     fieldsMock,
	}

	return s
}

func (t *FieldTestSuite) GetTypeMockService(fnc func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository)) *Service {

	categoryMock := &mocks.CategoryRepository{}
	mediaMock := &mocks.MediaRepository{}
	postsMock := &mocks.PostsRepository{}
	userMock := &mocks.UserRepository{}

	fnc(categoryMock, mediaMock, postsMock, userMock)

	s := t.GetService(nil)
	s.store = &models.Store{
		Categories: categoryMock,
		Media:      mediaMock,
		Posts:      postsMock,
		User:       userMock,
	}

	return s
}

func (t *FieldTestSuite) GetService(fields []domain.PostField) *Service {
	return &Service{
		fields: fields,
	}
}
