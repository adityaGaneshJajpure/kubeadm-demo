package utils

import (
	"fmt"
	"get-bird/pkg/constants"
	"get-bird/pkg/dto"
	"os"
)

func DefaultBird(err error) dto.Bird {
	return dto.Bird{
		Name:        os.Getenv(constants.DEFAULT_BIRD_NAME),
		Description: fmt.Sprintf("This bird is in disguise because: %s", err),
		Image:       os.Getenv(constants.DEFAULT_BIRD_IMAGE),
	}
}
