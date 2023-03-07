package main

import (
	"github.com/z0gSh1u/gomon/server/core"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	core.RegisterCreateCommentRouter(r)

	r.Run()
}
