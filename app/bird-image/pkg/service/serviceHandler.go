package service

import (
	"context"
	"encoding/json"
	"fmt"
	"get-bird-image/pkg/constants"
	"get-bird-image/pkg/dto"
	"get-bird-image/pkg/http/httpclient"
	"get-bird-image/pkg/logger"
	"get-bird-image/pkg/utils"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

type IServiceHandler interface {
	HandleRequest(c *gin.Context)
}

type ServiceHandler struct {
	logger     logger.ILogger
	httpClient httpclient.IHTTPClient
}

func NewServiceHandler(logger logger.ILogger, httpClient httpclient.IHTTPClient) *ServiceHandler {
	return &ServiceHandler{logger: logger, httpClient: httpClient}
}

func (s *ServiceHandler) HandleRequest(c *gin.Context) {
	ctx := context.Background()

	birdName := c.Query("birdName")
	if birdName == "" {
		s.logger.Errorf(ctx, "birdName parameter not found")
		s.prepareResponse(c, http.StatusBadRequest, false, utils.DefaultBirdImage())
		return
	}

	birdImage, err := s.getBirdImage(ctx, birdName)
	if err != nil {
		s.logger.Errorf(ctx, "Error fetching bird details: %s", err)
		s.prepareResponse(c, http.StatusBadRequest, false, utils.DefaultBirdImage())
		return
	}

	s.prepareResponse(c, http.StatusOK, true, *birdImage)
}

func (s *ServiceHandler) getBirdImage(ctx context.Context, birdName string) (*dto.BirdImage, error) {
	query := fmt.Sprintf(
		"%s/search/photos?page=1&query=%s&client_id=%s&per_page=1",
		os.Getenv(constants.UNSPLASH_API),
		url.QueryEscape(birdName),
		os.Getenv(constants.UNSPLASH_API_KEY),
	)

	res, err := s.httpClient.Get(query)
	if err != nil {
		s.logger.Errorf(ctx, "Error reading image API: %s", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		s.logger.Errorf(ctx, "Error parsing image API response: %s", err)
		return nil, err
	}

	var response dto.ImageResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		s.logger.Errorf(ctx, "Error unmarshalling birdImage: %s", err)
		return nil, err
	}

	return &dto.BirdImage{Image: response.Results[0].Urls.Thumb}, nil
}

func (s *ServiceHandler) prepareResponse(c *gin.Context, statusCode int, success bool, bird dto.BirdImage) {
	c.JSON(statusCode, gin.H{"success": success, "data": bird})
}
