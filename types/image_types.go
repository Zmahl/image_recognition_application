package types

type ImageRequest struct {
	ImageExtension
}

type ImageResponse struct {
	ObjectName string `json:"object_name"`
}
