package server

import (
	"net/http"
	"os"

	"get-bird/pkg/constants"
	"get-bird/pkg/service"

	"get-bird/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IServer interface {
	InitServer() error
	SetupRoutes() error
	RunServer() error
}

type Server struct {
	GinServer *gin.Engine
	Logger    logger.ILogger
}

func (s *Server) InitServer() error {
	//init gin server
	env := os.Getenv(constants.ENVIRONMENT)
	s.GinServer = gin.Default()

	if env == constants.PROD_ENV {
		gin.SetMode(gin.ReleaseMode)
	} else if env == constants.DEV_ENV {
		gin.SetMode(gin.DebugMode)
	} else if env == constants.TEST_ENV {
		gin.SetMode(gin.TestMode)
	}

	return nil
}

func (s *Server) SetupRoutes(serviceHandler *service.ServiceHandler) error {
	s.GinServer.GET("/health", HealthCheck)
	s.GinServer.GET("/", serviceHandler.HandleRequest)
	return nil
}

func (s *Server) RunServer() error {
	port := os.Getenv(constants.PORT)
	err := s.GinServer.Run(":" + port)
	if err != nil {
		s.Logger.Error("error while running the server", zap.Error(err))
		return err
	}
	s.Logger.Info("gin server is started!")
	return nil
}

func HealthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, constants.STATUS_OK)
}
