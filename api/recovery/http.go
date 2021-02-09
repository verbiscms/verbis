package recovery

import (
	"github.com/gin-gonic/gin"
	"io"
	"net"
	"os"
	"strings"
)

func (r *Recover) HttpRecovery(out io.Writer) gin.HandlerFunc {
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
					r.Recover(ctx, err)
					//ctx.Data(500, "text/html", ctx.Read)
					return
				}
			}
		}()
		ctx.Next()
	}
}
