# GoTree

GoTree is a Golang implementation of the UNIX [tree](https://en.wikipedia.org/wiki/Tree_(Unix)) program, which produces a depth-indented listing of files.

## Installation

Use [`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to install and update:

```sh
$ go get -u github.com/std0/gotree
```

## Usage

Use [`go run`](https://golang.org/cmd/go/#hdr-Compile_and_run_Go_program) or [`go build`](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies) to run or build the program, respectively.

```sh
$ go run main.go PATH [options]
```

```sh
$ go build
$ ./gotree PATH [options]
```

Available options:

```
  -f  optional. Output both files and dirs (default is dirs only)
```

## Example

```sh
$ go run main.go testdata -f
testdata
├───dir1
│   ├───dir1_1
│   │   └───text1_1_1.txt (665b)
│   └───text1_1.txt (835b)
├───dir2
│   └───text2_1.txt (empty)
└───text1.txt (551b)
```

```sh
$ go build
$ ./gotree testdata -f
testdata
├───dir1
│   ├───dir1_1
│   │   └───text1_1_1.txt (665b)
│   └───text1_1.txt (835b)
├───dir2
│   └───text2_1.txt (empty)
└───text1.txt (551b)
```