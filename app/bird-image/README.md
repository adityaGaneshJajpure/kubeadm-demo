# Bird Image Application

This is a sample application written in Golang and contains 2 APIs:
- `/?birdName=<BIRD-NAME>` -> returns details of the bird based on the `birdName` query parameter
- `/health` -> health check endpoint for the application

## Setup

### Required Environment Variables

- `port`: port on which server would listen
- `LOG_LEVEL`: log level for application
- `UNSPLASH_API`: api.unsplash.com endpoint
- `UNSPLASH_API_KEY`: client_id used in `UNSPLASH_API`
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
                // kafka configs
                "PORT": "8001",
                "LOG_LEVEL": "debug",
                "UNSPLASH_API": "https://api.unsplash.com",
                "UNSPLASH_API_KEY": "<ADD-KEY-HERE>",
                "DEFAULT_BIRD_IMAGE": "https://www.pokemonmillennium.net/wp-content/uploads/2015/11/missingno.png"
            },
        }
    ]
}
```

