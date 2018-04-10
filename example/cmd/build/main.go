package main // import "github.com/IngCr3at1on/x/build/example/cmd/build"

import (
	"fmt"
	"os"

	"github.com/IngCr3at1on/x/build"
	"github.com/IngCr3at1on/x/build/wrapper"
)

func main() {
	_pkg := fmt.Sprintf("github.com/IngCr3at1on/x/build/example/cmd/example")

	_wd, err := os.Getwd()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	_out := fmt.Sprintf("%s/out", _wd)
	if err := os.MkdirAll(_out, os.ModeDir|0700); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	wrapper.Builder = &build.Builder{
		Flows: []*build.Flow{
			&build.Flow{
				Name: "build",
				Steps: []func() error{
					func() error {
						return build.Exec(nil, "vgo", "build", "-o", fmt.Sprintf("%s/app", _out), _pkg)
					},
				},
			},
		},
	}

	if err := wrapper.Run(os.Args[1:]...); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
