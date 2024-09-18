# Bird Application

This is a sample application written in Golang and contains 2 APIs:
- `/` -> returns details of a random bird
- `/health` -> health check endpoint for the application

## Setup

### Required Environment Variables

- `port`: port on which server would listen
- `LOG_LEVEL`: log level for application
- `BIRDS_API`: freetestapi endpoint to get bird details
- `BIRD_IMAGE_SERVER_ENDPOINT`: server endpoint for `bird-image` application
- `DEFAULT_BIRD_NAME`: default bird name to be returned in case of error
- `DEFAULT_BIRD_IMAGE`: default bird image to be returned in case of error

Example launch.json configuration
```json
{
    "version": "0.2.0",
    "configurations": [
        
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",

            "program": "${fileDirname}",
            "env": {
                "port": "8000",
                "LOG_LEVEL": "debug",
                "BIRDS_API": "https://freetestapi.com/api/v1/birds/",
                "BIRD_IMAGE_SERVER_ENDPOINT": "http://localhost:8001",
                "DEFAULT_BIRD_NAME": "Bird in disguise",
                "DEFAULT_BIRD_IMAGE": "https://www.pokemonmillennium.net/wp-content/uploads/2015/11/missingno.png"
            },
        }
    ]
}
```

