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
	cli.Usage = Usage(cli)

	var includeGit bool

	cli.BoolVar(&includeGit, "git", false, "Include files located in a \".git\" folder.")

	var includeRust bool

	cli.BoolVar(&includeRust, "rust", false, "Include files ending on the \".rs\" extension (temporary flag)")

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

		isRustFile := filepath.Ext(path) == ".rs"

		if isExe && (!isGitSubfolder || includeGit) && (!isRustFile || includeRust) {
			fmt.Fprintln(os.Stdout, path)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func Usage(f *flag.FlagSet) func() {
	return func() {
		fmt.Fprintf(f.Output(), "Usage of %s: path\n", f.Name())
		f.PrintDefaults()
	}
}
