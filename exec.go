package build

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

// Exec wraps os.Exec for convenience. If provided out will have cmd.Stdout
// written to it, otherwise cmd.Stdout will be piped into os.Stdout.
// cmd.Stderr will either be wrapped into the error or piped to os.Stderr,
// respective to the Stdout behavior.
func Exec(out *bytes.Buffer, command string, params ...string) error {
	cmd := exec.Command(command, params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	var errb bytes.Buffer
	if out != nil {
		cmd.Stdout = out
		cmd.Stderr = &errb
	}

	if err := cmd.Run(); err != nil {
		if out != nil {
			return errors.Wrapf(err, "error (%s) executing command: %s, %v", errb.String(), command, cmd.Args)
		}

		return errors.Wrapf(err, "error executing command: %s, %v", command, cmd.Args)
	}

	return nil
}
