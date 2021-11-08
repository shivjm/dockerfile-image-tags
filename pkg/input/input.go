package input

import (
	"fmt"
	"os"

	"github.com/asottile/dockerfile"
)

// ParseInput parses either a file (if named) or `os.Stdin` into a
// list of Dockerfile commands.
func ParseInput(args []string) ([]dockerfile.Command, error) {
	input, err := GetInput(args)

	if err != nil {
		return nil, err
	}

	return dockerfile.ParseReader(input)
}

// GetInput opens and returns the file named in `args` if present,
// `os.Stdin` if not. The special file `-` can be passed explicitly to
// read from `Stdin`.
func GetInput(args []string) (*os.File, error) {
	if len(args) == 1 {
		name := args[0]

		if name == "-" {
			return os.Stdin, nil
		}

		f, err := os.Open(name)

		if err != nil {
			return nil, fmt.Errorf("could not open %s: %s", name, err)
		}

		return f, nil
	}

	return os.Stdin, nil
}
