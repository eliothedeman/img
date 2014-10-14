package image

// Image interface that describes an image
type Image interface {
	Name() string
	Data(req Request) []byte
}
