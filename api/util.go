package api

// GenPath creates a path for a given request
func GenPath(id string, codec string, size string) string {
	return "data/" + id + "/" + codec + "/" + size
}
