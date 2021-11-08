package cmd

import (
	"github.com/spf13/cobra"
)

const (
	ProgramName = "dockerfile-image-tags"
)

var (
	unknownMarker string

	rootCmd = &cobra.Command{
		Use:   ProgramName,
		Short: "List or query images and tags used in a Dockerfile.",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&unknownMarker, "unknown-marker", "m", "?", "string to use to indicate unknown tag")
}
