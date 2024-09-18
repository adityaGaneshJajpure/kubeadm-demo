package utils

import (
	"get-bird-image/pkg/constants"
	"get-bird-image/pkg/dto"
	"os"
)

func DefaultBirdImage() dto.BirdImage {
	return dto.BirdImage{
		Image: os.Getenv(constants.DEFAULT_BIRD_IMAGE),
	}
}
