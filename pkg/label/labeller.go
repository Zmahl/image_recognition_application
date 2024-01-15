package label

type Labeller interface {
	LabelImage(string) (*LabelResponse, error)
	GetLabelCredentials() string
}
