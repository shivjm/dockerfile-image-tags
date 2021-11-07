package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/asottile/dockerfile"
	"github.com/spf13/cobra"
)

type Image struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

func main() {
	var unknownMarker string
	var query string

	var rootCmd = &cobra.Command{
		Use:   "dockerfile-image-tags",
		Short: "List or query images & tags used in a Dockerfile.",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			file, err := getInput(args)

			if err != nil {
				log.Fatalf("Could not read Dockerfile: %s", err)
			}

			parsed, err := dockerfile.ParseReader(file)

			if err != nil {
				log.Fatalf("Could not parse Dockerfile: %s\n", err)
			}

			images := getImages(parsed, unknownMarker)

			val, err := json.Marshal(images)

			if err != nil {
				log.Fatalf("Could not serialize images as JSON: %s\n", err)
			}

			if query == "" {
				fmt.Println(string(val))
			} else {
				tag, err := getSingleTag(images, query)

				if err != nil {
					log.Fatalf("Could not find image in Dockerfile: %s", query)
				}

				fmt.Println(tag)
			}
		},
	}
	rootCmd.Flags().StringVarP(&unknownMarker, "unknown-marker", "m", "?", "string to use to indicate unknown tags")
	rootCmd.Flags().StringVarP(&query, "query", "q", "", "single image to return tag for (first occurrence)")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// getInput opens the file named in `args` if present, `os.Stdin` if not.
func getInput(args []string) (*os.File, error) {
	if len(args) == 1 {
		name := args[0]
		f, err := os.Open(name)

		if err != nil {
			return nil, fmt.Errorf("could not open %s: %s", name, err)
		}

		return f, nil
	}

	return os.Stdin, nil
}

// getImages returns the `Image`s used in the given sets of Dockerfile commands.
func getImages(commands []dockerfile.Command, unknownMarker string) []Image {
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

func getSingleTag(images []Image, query string) (string, error) {
	for _, i := range images {
		if i.Name == query {
			return i.Tag, nil
		}
	}

	return "", fmt.Errorf("could not find image %s in list", query)
}
