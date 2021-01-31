/*
* @Time    : 2021-01-31 17:18
* @Author  : CoderCharm
* @File    : main.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/wxyMicro/user/domain/repository"
	service2 "github.com/wxyMicro/user/domain/service"
	"github.com/wxyMicro/user/handler"
	user "github.com/wxyMicro/user/proto/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	//服务参数设置
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)

	//2初始化服务
	srv.Init()

	//3. 初始化中间件 创建数据库连接
	dsn := "root:Admin12345-@tcp(172.16.137.129:3306)/micro?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		fmt.Println(err)
	}
	rp := repository.NewUserRepository(db)

	// 初始化表 只用运行一次
	//_ = rp.InitTable()

	//4.创建服务实例
	userDataService := service2.NewUserDataService(rp)

	//5.注册handler
	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println(err)
	}

	//6. Run service
	if err := srv.Run(); err != nil {
		fmt.Println()
	}

}
