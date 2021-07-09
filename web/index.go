package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/config"
	"go-gin/domain"
	"go-gin/web/book_route"
	"go-gin/web/user_route"
)

func Setup() {
	r := gin.Default()
	r.Use(domain.Recover)

	r.GET("/", func(c *gin.Context) {
		domain.Ok(c, nil)
	})

	user_route.Setup(r)
	book_route.Setup(r)

	if config.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	_ = r.Run(fmt.Sprintf(":%s", config.Config.Server.Port))
}
