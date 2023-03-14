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

		if err := CreateComment(&comment); err != nil {
			Fail(err.Error(), nil, ctx)
		} else {
			Success("success create comment", nil, ctx)
		}
	})
}

func RegisterSelectCommentByTopicIdRouter(r *gin.Engine) {
	r.GET("/get_comments/:topicId", func(ctx *gin.Context) {
		topicId := ctx.Param("topicId")

		if comments, err := SelectComments(topicId); err != nil {
			Fail(err.Error(), nil, ctx)
		} else {
			Success("get_comments success", comments, ctx)
		}
	})
}
