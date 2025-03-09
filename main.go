package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "xingtu/api/router"
    "xingtu/config"
    "xingtu/dao"
    "xingtu/model"
    "xingtu/service"
    "xingtu/api/handler"
    _ "xingtu/docs" // 引入 Swagger 文档
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
    r.Run(":18080")
}