package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"git.bode.fun/magic"
)

func main() {
	if err := mainE(); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
}

func mainE() error {
	cli := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	if err := cli.Parse(os.Args[1:]); err != nil {
		return err
	}

	rootPath := cli.Arg(0)

	rootPath, err := filepath.Abs(rootPath)
	if err != nil {
		return err
	}

	err = filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil //nolint
		}

		if d.IsDir() {
			return nil
		}

		isExe, err := magic.IsAnExecutable(path)
		if err != nil {
			return nil //nolint
		}

		isGitSubfolder := strings.Contains(path, ".git")

		if isExe && !isGitSubfolder {
			fmt.Fprintln(os.Stdout, path)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
