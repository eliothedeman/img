package image

// JPG is an image for the .jpg file type
type JPG struct {
	Raw  []byte
	name string
}

// Name get the name of the file
func (j *JPG) Name() string {
	return j.name
}

// Data returns the data for the given file based on it's request
func (j *JPG) Data(req Request) []byte {
	return j.Raw
}
