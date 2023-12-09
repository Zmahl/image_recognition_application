package label

import (
	"github.com/gin-gonic/gin"
)

type Labeller interface {
	LabelImage(*gin.Context, string, string)
	GetLabelCredentials() string
}
