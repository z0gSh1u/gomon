package core

import (
	"github.com/gin-gonic/gin"
)

func NewCommentRouter(r *gin.Engine) {
	r.POST("/new", func(ctx *gin.Context) {
		// ...
	})
}
