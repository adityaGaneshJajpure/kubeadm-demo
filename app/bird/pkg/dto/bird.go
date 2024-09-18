package dto

type Bird struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type BirdImageResponse struct {
	Data BirdImage `json:"data"`
}

type BirdImage struct {
	Image string `json:"image"`
}
