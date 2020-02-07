package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var (
		inputptr  = flag.String("i", "", "input string to be chopped")
		chunkptr  = flag.Uint("n", 64, "chunk size")
		prefixptr = flag.String("pre", "\"", "prefix to be added for each chunk")
		suffixptr = flag.String("suf", "\"", "suffix to be added for each chunk")
		joinptr   = flag.String("join", " + \n", "join chunks with (not added at the end)")
	)

	// parse cmd line options
	flag.Parse()

	// set input stream
	var input io.ReadCloser = os.Stdin
	if *inputptr != "" {
		input = ioutil.NopCloser(strings.NewReader(*inputptr))
	}
	defer input.Close()

	// chop some input
	out := chop(input,
		*prefixptr,
		*suffixptr,
		*joinptr,
		int(*chunkptr))

	// print to output
	fmt.Println(out)
}

// chop reads until EOF then chops stream into chunks. Each chunk may
// or may not be wrapped with a prefix and a suffix.
func chop(r io.Reader, prefix, suffix, join string, chunk int) string {
	// do this the simple way
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	bld := strings.Builder{}
	bld.Grow(chunk * 10)
	for a := 0; a < len(buf); a += chunk {
		// make sure we terminate early when chunk is too large
		cnt := min(chunk, len(buf)-a)

		bld.WriteString(prefix)
		bld.Write(buf[a : a+cnt])
		bld.WriteString(suffix)

		// if there is more data to be read, add join
		if a+chunk < len(buf) {
			bld.WriteString(join)
		}
	}

	return bld.String()
}

// min returns the minimum value of a and b
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
