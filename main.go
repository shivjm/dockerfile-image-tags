package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/asottile/dockerfile"
)

const (
	UNKNOWN_MARKER = "?"
)

type Image struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func main() {
	file, err := getInput(os.Args)

	if err != nil {
		log.Fatalf("Could not read Dockerfile: %s", err)
	}

	parsed, err := dockerfile.ParseReader(file)

	if err != nil {
		log.Fatalf("Could not parse Dockerfile: %s\n", err)
	}

	images := getTags(parsed, UNKNOWN_MARKER)

	val, err := json.Marshal(images)

	if err != nil {
		log.Fatalf("Could not serialize images as JSON: %s\n", err)
	}

	fmt.Println(string(val))
}

// getInput opens the file named in `args` if present, `os.Stdin` if not.
func getInput(args []string) (*os.File, error) {
	if len(args) > 1 {
		name := args[1]
		f, err := os.Open(name)

		if err != nil {
			return nil, fmt.Errorf("Could not open %s: %s", name, err)
		}

		return f, nil
	}

	return os.Stdin, nil
}

// getTags returns the `Image`s used in the given sets of Dockerfile commands.
func getTags(commands []dockerfile.Command, unknownMarker string) []Image {
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
