package main

import (
	"log"
	"xingtu/api/handler"
	"xingtu/api/router"
	"xingtu/config"
	"xingtu/dao"
	_ "xingtu/docs" // 引入 Swagger 文档
	"xingtu/model"
	"xingtu/service"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        panic(err)
    }
    db, err := gorm.Open(mysql.Open(cfg.MySQL.DSN), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    err = db.AutoMigrate(&model.User{}, &model.Question{}, &model.TestRecord{})
    if err != nil {
        panic(err)
    }

    dao := dao.NewDAO(db)
    service := service.NewService(dao)
    handler := handler.NewHandler(service)
    r := router.SetupRouter(handler)

    log.Printf("zhiyuxingtu running")
    r.Run(":18080")
}