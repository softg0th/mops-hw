package main

import (
	logstash_logger "github.com/KaranJagtiani/go-logstash"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"iotController/internal/exceptions"
	infra "iotController/internal/infra"
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
	// rabbitHost := os.Getenv("RABBIT_HOST")
	rabbitQueuName := os.Getenv("RABBIT_QUEUE_NAME")
	prometheusPort := os.Getenv("PROMETHEUS_PORT")

	logstashPort, err := strconv.Atoi(os.Getenv("LOGSTASH_PORT"))

	if err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "Port", Message: "incorrect port type"})
		return
	}

	rabbitConn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")

	if err != nil {
		log.Printf(err.Error())
		exceptions.HandleError(&exceptions.CMDError{Field: "Rabbit", Message: "failed to connect to RabbitMQ"})
		return
	}

	defer rabbitConn.Close()

	rabbitChannel, err := rabbitConn.Channel()
	newRabbit, err := infra.NewRabbit(rabbitChannel, rabbitQueuName)
	if err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "Queue", Message: "failed to create rabbit queue"})
	}

	logger := logstash_logger.Init("logstash", logstashPort, logstashProtocol, 5)

	conn, err := repository.NewMongoConnection(dbUrl)
	if err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "Mongo", Message: "failed to connect to MongoDB"})
		return
	}

	db := repository.NewDataBase(conn, dbName, collectionName)
	iotService := service.NewService(db, logger, newRabbit)
	listen, err := net.Listen(grpcProtocol, grpcPort)
	if err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "GRPC", Message: "failed to connect to GRPC service"})
		return
	}
	defer listen.Close()

	serv := grpc.NewServer()
	iotServer := server.NewServer(iotService)
	pb.RegisterIotServiceServer(serv, iotServer)

	prometheus.MustRegister(infra.RequestsTotal)
	prometheus.MustRegister(infra.RequestDuration)
	go service.RunMetricServer(iotService, prometheusPort)

	iotService.Logger.Info(map[string]interface{}{
		"message": "Server listening at",
		"addr":    grpcProtocol + ":" + grpcPort,
	})

	if err := serv.Serve(listen); err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "GRPC Run", Message: "failed to run GRPC service"})
		return
	}
}
