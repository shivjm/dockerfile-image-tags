package cmd

import (
	"fmt"
	"log"

	"github.com/shivjm/dockerfile-image-tags/pkg/images"
	"github.com/shivjm/dockerfile-image-tags/pkg/input"
	"github.com/spf13/cobra"
)

var (
	occurrence int

	queryCmd = &cobra.Command{
		Use:   "query",
		Short: "Find single image tag",
		Long:  "Print the tag for a specific image if found, exit with an error otherwise. The image may appear multiple times in the Dockerfile, in which case `occurrence` (default: 1) specifies which instance to return the tag for.",
		Args:  cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, rawArgs []string) {
			var args []string
			var query string

			if len(rawArgs) == 2 {
				args, query = rawArgs[0:1], rawArgs[1]
			} else {
				query = rawArgs[0]
				args = []string{}
			}

			parsed, err := input.ParseInput(args)

			if err != nil {
				log.Fatalf("Could not parse Dockerfile: %s\n", err)
			}

			allImages := images.GetImages(parsed, unknownMarker)

			tag, err := images.GetSingleTag(allImages, query, occurrence)

			if err != nil {
				log.Fatalf("Could not find image in Dockerfile: %s", query)
			}

			fmt.Println(tag)
		},
	}
)

func init() {
	queryCmd.Flags().IntVarP(&occurrence, "occurrence", "n", 1, "which occurrence of image to return tag for")

	rootCmd.AddCommand(queryCmd)
}
