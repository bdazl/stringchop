# strchop

This is a small command line utility to chop strings into chunks.
You must first install the Go compiler toolchain, in order to use it.

```bash
$ # Navigate to the source code directory
$ go run ./cmd/strchop --help

$ # Alternatively, install it first:
$ go install ./cmd/strchop
$ strchop -i "this is a string" -n 4
"this" + 
" is " + 
"a st" + 
"ring"
```

## usage

```bash
$ strchop --help
Usage of strchop:
  -i string
        input string to be chopped
  -join string
        join chunks with (not added at the end) (default " + \n")
  -n uint
        chunk size (default 64)
  -pre string
        prefix to be added for each chunk (default "\"")
  -suf string
        suffix to be added for each chunk (default "\"")
```

You can either put in a string with that param `-i`, or if you don't supply it,
the input can come from stdin:

```bash
$ echo -n "face" | strchop -n 2 -join +
"fa"+"ce"
```
