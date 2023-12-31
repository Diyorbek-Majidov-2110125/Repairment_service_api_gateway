package handlers

import (
	"projects/Repairment_service/Repairment_service_api_gateway/api/http"
	"projects/Repairment_service/Repairment_service_api_gateway/config"
	"projects/Repairment_service/Repairment_service_api_gateway/grpc/client"
	"projects/Repairment_service/Repairment_service_api_gateway/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg      config.Config
	log      logger.LoggerI
	services client.ServiceManagerI
}

func NewHandler(cfg config.Config, log logger.LoggerI, services client.ServiceManagerI) *Handler {
	return &Handler{
		cfg:      cfg,
		log:      log,
		services: services,
	}
}

func (h *Handler) handleResponse(c *gin.Context, status http.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			// logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}
