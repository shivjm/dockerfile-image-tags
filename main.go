package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/asottile/dockerfile"
	"github.com/spf13/cobra"

	"github.com/shivjm/dockerfile-image-tags/pkg/images"
	"github.com/shivjm/dockerfile-image-tags/pkg/input"
)


func main() {
	var unknownMarker string
	var query string
	var occurrence int

	var rootCmd = &cobra.Command{
		Use:   "dockerfile-image-tags",
		Short: "List or query images & tags used in a Dockerfile.",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			file, err := input.GetInput(args)

			if err != nil {
				log.Fatalf("Could not read Dockerfile: %s", err)
			}

			parsed, err := dockerfile.ParseReader(file)

			if err != nil {
				log.Fatalf("Could not parse Dockerfile: %s\n", err)
			}

			allImages := images.GetImages(parsed, unknownMarker)

			val, err := json.Marshal(allImages)

			if err != nil {
				log.Fatalf("Could not serialize images as JSON: %s\n", err)
			}

			if query == "" {
				fmt.Println(string(val))
			} else {
				tag, err := images.GetSingleTag(allImages, query, occurrence)

				if err != nil {
					log.Fatalf("Could not find image in Dockerfile: %s", query)
				}

				fmt.Println(tag)
			}
		},
	}
	rootCmd.Flags().StringVarP(&unknownMarker, "unknown-marker", "m", "?", "string to use to indicate unknown tags")
	rootCmd.Flags().StringVarP(&query, "query", "q", "", "single image to return tag for (see `occurrence`)")
	rootCmd.Flags().IntVarP(&occurrence, "occurrence", "n", 1, "which occurrence of image to return tag for")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
