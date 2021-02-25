package recovertemp

import (
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/tpl"
	"github.com/gin-gonic/gin"
)

type Recovery interface {
	Recover(cfg Config) []byte
	HttpRecovery() gin.HandlerFunc
}

// Config defines
type Config struct {
	Code    int
	Context *gin.Context
	Error   interface{}
	TplFile string
	TplExec tpl.TemplateExecutor
	Post    *domain.PostDatum
}
