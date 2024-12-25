package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type handle struct {
	logger *log.Helper
}

func NewRouter(chain string, r *gin.Engine) {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)

	handle := &handle{
		logger: log.NewHelper(logger),
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	loadMfileDIDMoudles(r.Group("/mfile"), handle)
	loadDIDDomainMoudles(r.Group("/domain"), handle)
}
