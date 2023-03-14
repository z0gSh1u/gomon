package main

import (
	"github.com/z0gSh1u/gomon/server/core"
	"github.com/z0gSh1u/gomon/server/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware("*", "GET, POST, OPTIONS, PUT"))

	core.RegisterCreateCommentRouter(r)
	core.RegisterSelectCommentByTopicIdRouter(r)

	r.Run()
}
