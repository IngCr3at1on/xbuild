package build

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type (
	// Builder is a super silly go project builder.
	Builder struct {
		Flows []*Flow
	}

	// Flow defines a build flow.
	Flow struct {
		Name  string
		Steps []func() error
	}
)

// Build runs a given build step.
func (b *Builder) Build(args []string) error {
	if len(b.Flows) == 0 {
		return errors.New("No build configurations defined")
	}

	found := false
	for _, flow := range b.Flows {
		if len(args) == 0 {
			fmt.Fprintf(os.Stdout, "Name: %s\n", flow.Name)
		}

		if len(args) == 1 && strings.EqualFold(strings.TrimSpace(args[0]), flow.Name) {
			found = true

			if err := run(flow.Steps); err != nil {
				return err
			}
			break
		}
	}

	if !found && len(args) > 0 {
		return errors.New("build step not recognized")
	}

	return nil
}

func run(steps []func() error) error {
	for _, step := range steps {
		if err := step(); err != nil {
			return err
		}
	}

	return nil
}
