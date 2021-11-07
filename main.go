package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/asottile/dockerfile"
)

type Image struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func main() {
	file := os.Stdin

	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])

		if err != nil {
			log.Fatalf("Could not open %s: %s", os.Args[1], err)
		}

		file = f
	}

	parsed, err := dockerfile.ParseReader(file)

	if err != nil {
		log.Fatalf("Could not parse Dockerfile: %s\n", err)
	}

	images := []Image{}

	for _, cmd := range parsed {
		if cmd.Cmd == "FROM" {
			full := cmd.Value
			rawImage := full[0]
			imageParts := strings.Split(rawImage, ":")
			image := imageParts[0]
			var version string

			if len(imageParts) > 1 {
				version = imageParts[1]
			} else {
				version = "?"
			}

			images = append(images, Image{Name: image, Tag: version})
		}
	}

	val, err := json.Marshal(images)

	fmt.Println(string(val))
}
