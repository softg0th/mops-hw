package main

import (
	logstash_logger "github.com/KaranJagtiani/go-logstash"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"os"
	"ruleEngine/internal/exceptions"
	"ruleEngine/internal/infra"
	"ruleEngine/internal/service"
	"strconv"
)

func main() {
	err := godotenv.Load("/rule-engine/cmd/.env")
	if err != nil {
		exceptions.HandleError(&exceptions.CMDError{Field: "DotEnv", Message: "failed to load env file"})
		return
	}

	logstashProtocol := os.Getenv("LOGSTASH_PROTOCOL")
	rabbitQueueName := os.Getenv("RABBIT_QUEUE_NAME")
	prometheusPort := os.Getenv("PROMETHEUS_PORT")

	logstashPort, err := strconv.Atoi(os.Getenv("LOGSTASH_PORT"))

	if err != nil {
		log.Printf(err.Error())
		exceptions.HandleError(&exceptions.CMDError{Field: "Port", Message: "incorrect port type"})
		return
	}

	rabbitConn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")

	if err != nil {
		log.Printf(err.Error())
		exceptions.HandleError(&exceptions.CMDError{Field: "Rabbit", Message: "failed to connect to RabbitMQ"})
		return
	}

	rabbitChannel, err := rabbitConn.Channel()
	newRabbit, err := infra.NewRabbit(rabbitChannel, rabbitQueueName)
	if err != nil {
		log.Printf(err.Error())
		exceptions.HandleError(&exceptions.CMDError{Field: "Queue", Message: "failed to create rabbit queue"})
		return
	}

	logger := logstash_logger.Init("logstash", logstashPort, logstashProtocol, 5)

	ruleService := service.NewService(logger, newRabbit)

	ruleService.Logger.Info(map[string]interface{}{
		"message": "Successfully connected to RabbitMQ"})

	http.Handle("/metrics", promhttp.Handler())
	prometheus.MustRegister(infra.InstantRulesTotal, infra.DurationRulesTotal)
	go func() {
		ruleService.Logger.Error(map[string]interface{}{
			"message": http.ListenAndServe(prometheusPort, nil),
			"error":   true,
		})
	}()

	ruleService.ReadFromRabbitMQ(5, 10)
	defer rabbitConn.Close()
}
