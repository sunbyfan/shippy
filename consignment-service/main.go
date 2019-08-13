package main

import (
	"github.com/micro/go-micro"
	"log"
	"os"
	pb "shippy/consignment-service/proto/consignment"
	vesselPb "shippy/vessel-service/proto/vessel"
)

const(
	defaultHost = "localhost:27017"
)

func main() {
	// 获取容器设置的数据库地址环境变量的值
	dbHost := os.Getenv("DB_HOST")
	if dbHost == ""{
		dbHost = defaultHost
	}
	session, err := CreateSession(dbHost)
	// 创建于 MongoDB 的主会话，需在退出 main() 时候手动释放连接
	defer session.Close()
	if err != nil {
		log.Fatalf("create session error: %v\n", err)
	}

	srv := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	vesselClient := vesselPb.NewVesselService("go.micro.srv.vessel", srv.Client())
	// 解析命令行参数
	srv.Init()
	// 作为 vessel-service 的客户端
	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
