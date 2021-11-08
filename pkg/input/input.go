package input

import (
	"fmt"
	"os"
)

// GetInput opens the file named in `args` if present, `os.Stdin` if not.
func GetInput(args []string) (*os.File, error) {
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
