package dto

type BirdImage struct {
	Image string `json:"image"`
}

type Urls struct {
	Thumb string `json:"thumb"`
}

type Links struct {
	Urls Urls `json:"urls"`
}

type ImageResponse struct {
	Results []Links `json:"links"`
}
