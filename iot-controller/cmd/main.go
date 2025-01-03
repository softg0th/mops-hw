package main

import (
	"fmt"
	logstash_logger "github.com/KaranJagtiani/go-logstash"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"iotController/internal/exceptions"
	pb "iotController/internal/proto"
	"iotController/internal/repository"
	"iotController/internal/server"
	"iotController/internal/service"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load("/iot-controller/cmd/.env")
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)
	if err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "DotEnv", Message: "failed to load env file"})
		return
	}

	grpcProtocol := os.Getenv("GRPC_PROTOCOL")
	grpcPort := os.Getenv("GRPC_PORT")
	dbUrl := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")
	logstashProtocol := os.Getenv("LOGSTASH_PROTOCOL")

	logstashPort, err := strconv.Atoi(os.Getenv("LOGSTASH_PORT"))

	if err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "Port", Message: "incorrect port type"})
		return
	}

	logger := logstash_logger.Init("logstash", logstashPort, logstashProtocol, 5)

	conn, err := repository.NewMongoConnection(dbUrl)
	if err != nil {
		log.Printf("connected to mongodb failed with error: %v", err)
		exceptions.HandleError(&exceptions.CMDError{Field: "Mongo", Message: "failed to connect to MongoDB"})
		return
	}
	log.Printf("connected to mongodb successfully")
	db := repository.NewDataBase(conn, dbName, collectionName)
	iotService := service.NewService(db, logger)
	listen, err := net.Listen(grpcProtocol, grpcPort)
	if err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "GRPC", Message: "failed to connect to GRPC service"})
		return
	}
	defer listen.Close()

	serv := grpc.NewServer()
	iotServer := server.NewServer(iotService)
	pb.RegisterIotServiceServer(serv, iotServer)

	iotService.Logger.Info(map[string]interface{}{
		"message": "Server listening at",
		"addr":    grpcProtocol + ":" + grpcPort,
	})

	if err := serv.Serve(listen); err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "GRPC Run", Message: "failed to run GRPC service"})
		return
	}
}
