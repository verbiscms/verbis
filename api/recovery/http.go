package recovery

import (
	"github.com/gin-gonic/gin"
	"net"
	"os"
	"strings"
)

func (r *Recover) HttpRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				// If the connection is dead, we can't write a status to it.
				if !brokenPipe {
					bytes := r.Recover(Config{
						Context: ctx,
						Error:   err,
					})
					ctx.Data(500, "text/html", bytes)
					return
				}
			}
		}()
		ctx.Next()
	}
}
