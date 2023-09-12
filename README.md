# Magic

Magic is a tool to recursively list all executable files from a given root path,
ignoring `.git` directories.

Currently, it only supports Mach-O and Fat files on booth big and little endian
architectures.
Scripts, executed via a shebang are recognized as well, but currently there is a
bug, which let's rust files starting with e.g. `#![no_std]` be recognized as well.

## Quick Start

```sh
$ git clone git@github.com:bode-fun/magic.git
$ cd magic
$ go build ./cmd/magic
$ ./magic --help
```
