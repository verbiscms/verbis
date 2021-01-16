package resolve

import (
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/logger"
	"github.com/ainsleyclark/verbis/api/models"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"testing"
)

type ResolverTestSuite struct {
	suite.Suite
}

type noStringer struct{}

func TestResolver(t *testing.T) {
	suite.Run(t, new(ResolverTestSuite))
}

func (t *ResolverTestSuite) BeforeTest(suiteName, testName string) {
	err := logger.Init(config.Configuration{})
	log.SetOutput(ioutil.Discard)
	t.NoError(err)
}

func (t *ResolverTestSuite) GetValue() *Value {
	return &Value{
		&models.Store{},
	}
}