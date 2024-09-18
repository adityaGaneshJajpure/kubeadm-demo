package constants

const (
	HEALTH_ENDPOINT            = "/health"
	STATUS_OK                  = "Ok"
	ENVIRONMENT                = "env"
	DEV_ENV                    = "dev"
	PROD_ENV                   = "prod"
	TEST_ENV                   = "test"
	CONTEXT_TIMEOUT_IN_SECONDS = 30
)

const (
	// required envs
	PORT                       = "PORT"
	BIRDS_API                  = "BIRDS_API"
	BIRD_IMAGE_SERVER_ENDPOINT = "BIRD_IMAGE_SERVER_ENDPOINT"
	DEFAULT_BIRD_NAME          = "DEFAULT_BIRD_NAME"
	DEFAULT_BIRD_IMAGE         = "DEFAULT_BIRD_IMAGE"
)
