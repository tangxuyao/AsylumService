package main

import (
	"../MongoData"
	"github.com/micro/go-micro"
	"./src"
	"proto/asylum"
	"fmt"
)

func main() {
	mg := &mongo.MongoDB{}
	if err := mg.Dial(""); err != nil {
		panic("连接[mongo]数据失败，请检查相关参数")
	}

	service := micro.NewService(micro.Name("AsylumService"))
	service.Init()

	defer func() {
		service.Server().Deregister()
	}()

	handler := &asylum.AsylumService{M: mg}
	asylum_api.RegisterAsylumServiceHandler(service.Server(), handler)

	fmt.Println("AsylumService starting...")

	if err := service.Run(); err != nil {
		panic(err)
	}
}
