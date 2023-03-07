package core

import (
	"github.com/gin-gonic/gin"
)

func RegisterCreateCommentRouter(r *gin.Engine) {
	r.POST("/create_comment", func(ctx *gin.Context) {
		var comment Comment
		if err := ctx.ShouldBindJSON(&comment); err != nil {
			Fail("bind no json", nil, ctx)
		}
		comment.Censored = true
		CreateComment(&comment)
		Success("success create comment", nil, ctx)
	})
}
