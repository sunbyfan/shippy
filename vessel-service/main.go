package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"log"
	"os"
	pb "shippy/vessel-service/proto/vessel"
)

const(
	defaultHost="localhost:27017"
)
func createDummyData(repo Repository){
	defer repo.Close()
	vessels:=[]*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	for _,v:=range vessels{
		repo.Create(v)
	}
}
func main() {
	// 停留在港口的货船，先写死
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	defer session.Close()
	if err != nil {
		log.Fatalf("Error connecting to datastore: %v", err)
	}
	repo := &VesselRepository{session.Copy()}

	createDummyData(repo)

	srv:=micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
		)
	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(),&service{session})

	if err:=server.Run();err!=nil {
		log.Fatalf("failed to server:%v",err)

	}
}
