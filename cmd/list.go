package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/asottile/dockerfile"
	"github.com/shivjm/dockerfile-image-tags/pkg/images"
	"github.com/shivjm/dockerfile-image-tags/pkg/input"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List images and tags",
		Long:  "Print a list of the images and tags in the Dockerfile as a JSON array. Every `FROM` instruction in the file corresponds to an entry in the array, so the same image might appear multiple times, with the same or different tags. A special marker (default: `?`) will be returned if the tag is not specified.",
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

			fmt.Println(string(val))
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
}
