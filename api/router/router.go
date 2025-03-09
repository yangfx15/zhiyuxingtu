package router

import (
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "xingtu/api/handler"
    _ "xingtu/docs" // 引入生成的 Swagger 文档
)

func SetupRouter(handler *handler.Handler) *gin.Engine {
    r := gin.Default()
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.POST("/login", handler.Login)
    r.GET("/questions", handler.GetQuestions)
    r.POST("/submit-test", handler.SubmitTest)
    return r
}