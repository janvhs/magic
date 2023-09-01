package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

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

	executablePaths := make([]string, 0)

	err = filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		isExe, err := magic.IsAnExecutable(path)
		if err != nil {
			return nil
		}

		if isExe {
			executablePaths = append(executablePaths, path)
		}

		return nil
	})
	if err != nil {
		return err
	}

	for _, path := range executablePaths {
		fmt.Fprintln(os.Stdout, path)
	}

	return nil
}
