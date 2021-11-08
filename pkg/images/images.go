package images

import (
	"fmt"
	"strings"

	"github.com/asottile/dockerfile"
)

type Image struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

// GetImages returns the `Image`s used in the given list of Dockerfile commands.
func GetImages(commands []dockerfile.Command, unknownMarker string) []Image {
	images := []Image{}

	for _, cmd := range commands {
		if cmd.Cmd == "FROM" {
			full := cmd.Value
			rawImage := full[0]
			imageParts := strings.Split(rawImage, ":")
			image := imageParts[0]
			var version string

			if len(imageParts) > 1 {
				version = imageParts[1]
			} else {
				version = unknownMarker
			}

			images = append(images, Image{Name: image, Tag: version})
		}
	}

	return images
}

// GetSingleTag returns the tag for the nth `occurrence` of `query` in
// the given list of `Image`s.
func GetSingleTag(images []Image, query string, occurrence int) (string, error) {
	found := 0

	for _, i := range images {
		if i.Name == query {
			found += 1

			if found >= occurrence {
				return i.Tag, nil
			}
		}
	}

	return "", fmt.Errorf("could not find image %s in list", query)
}
