package service

import (
	"context"
	"encoding/json"
	"fmt"
	"get-bird/pkg/constants"
	"get-bird/pkg/dto"
	"get-bird/pkg/http/httpclient"
	"get-bird/pkg/logger"
	"get-bird/pkg/utils"
	"io"
	"math/rand/v2"
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

	bird, err := s.getBird(ctx)
	if err != nil {
		s.logger.Errorf(ctx, "Error fetching bird details: %s", err)
		s.prepareResponse(c, http.StatusBadRequest, false, utils.DefaultBird(err))
		return
	}

	s.prepareResponse(c, http.StatusOK, true, *bird)
}

func (s *ServiceHandler) getBird(ctx context.Context) (*dto.Bird, error) {
	fmt.Println(fmt.Sprintf("%s%d", os.Getenv(constants.BIRDS_API), rand.IntN(50)))
	res, err := s.httpClient.Get(fmt.Sprintf("%s%d", os.Getenv(constants.BIRDS_API), rand.IntN(50)))
	if err != nil {
		s.logger.Errorf(ctx, "Error reading bird API: %s", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		s.logger.Errorf(ctx, "Error parsing bird API response: %s", err)
		return nil, err
	}

	var bird dto.Bird
	err = json.Unmarshal(body, &bird)
	if err != nil {
		s.logger.Errorf(ctx, "Error unmarshalling bird: %s", err)
		return nil, err
	}

	birdImage, err := s.getBirdImage(ctx, bird.Name)
	if err != nil {
		s.logger.Errorf(ctx, "Error in getting bird image: %s", err)
		return nil, err
	}

	bird.Image = birdImage
	return &bird, nil
}

func (s *ServiceHandler) getBirdImage(ctx context.Context, birdName string) (string, error) {
	res, err := http.Get(fmt.Sprintf("%s?birdName=%s", os.Getenv(constants.BIRD_IMAGE_SERVER_ENDPOINT), url.QueryEscape(birdName)))
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)

	var bird dto.BirdImageResponse
	err = json.Unmarshal(body, &bird)
	if err != nil {
		s.logger.Errorf(ctx, "Error unmarshalling bird: %s", err)
		return "", err
	}

	return bird.Data.Image, nil
}

func (s *ServiceHandler) prepareResponse(c *gin.Context, statusCode int, success bool, bird dto.Bird) {
	c.JSON(statusCode, gin.H{"success": success, "data": bird})
}
