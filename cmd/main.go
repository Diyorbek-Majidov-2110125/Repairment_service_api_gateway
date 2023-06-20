package main

import (
	"projects/Repairment_service/Repairment_service_api_gateway/api"
	"projects/Repairment_service/Repairment_service_api_gateway/api/handlers"
	"projects/Repairment_service/Repairment_service_api_gateway/config"
	"projects/Repairment_service/Repairment_service_api_gateway/grpc/client"
	"projects/Repairment_service/Repairment_service_api_gateway/pkg/logger"

	"github.com/gin-gonic/gin"
	// "google.golang.org/genproto/googleapis/cloud/apigateway/v1"
)

func main() {
	cfg := config.Load()

	var loggerLevel string

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			log.Error("!!!Cleanup->Error-->>>", logger.Error(err))
		}
	}()

	svcs, err := client.NewGrpcClients(cfg)
	if err != nil {
		log.Panic("client.NewGrpcClients", logger.Error(err))
	}

	// pubsubServer, err := events.New(cfg, log)
	// if err != nil {
	// 	log.Panic("error on the event server", logger.Error(err))
	// }

	h := handlers.NewHandler(cfg, log, svcs)

	r := api.SetUpRouter(*h, cfg)

	log.Info("HTTP: Server being started...", logger.String("port", cfg.ServicePort))

	err = r.Run(cfg.ServiceHost + cfg.ServicePort)
	if err != nil {
		panic(err)
	}
}
