package util

// GenPath creates a path for a given request
func GenPath(id, codec, size, prefix string) string {
	return prefix + id + "/" + codec + "/" + size
}
