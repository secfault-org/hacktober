package container

import "strings"

func ResolveImage(prefix, image string) string {
	if prefix == "" {
		return image
	}
	host := ""
	if before, _, ok := strings.Cut(image, "/"); ok {
		host = before
	}
	if strings.ContainsAny(host, ".:") || host == "localhost" {
		return image
	}
	return prefix + "/" + image
}
