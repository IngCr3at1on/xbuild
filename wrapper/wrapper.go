package wrapper

import (
	"errors"
	"flag"

	"github.com/IngCr3at1on/x/build"
)

var (
	// Builder is the wrapped builder.
	Builder *build.Builder
)

func init() {
	flag.Parse()
}

// Run passes args to the defined builder.
func Run(args ...string) error {
	if Builder == nil {
		return errors.New("builder must be set first")
	}

	return Builder.Build(args)
}
