package options

import "flag"

// represents the available [OPTIONS] for the gwc cli
type ArgOptions struct {
	L bool // represents the -l arg to count lines
	W bool // represents the -w arg to count words
	C bool // represents the -c arg to count bytes
	M bool // represents the -m arg to count characters
}

// Parse wraps flags.Parse function
// read args and sets defaults to [ArgOptions]
func (a *ArgOptions) Parse() *ArgOptions {
	flag.BoolVar(&a.L, "l", false, "Count lines")
	flag.BoolVar(&a.W, "w", false, "Count words")
	flag.BoolVar(&a.C, "c", false, "Count bytes")
	flag.BoolVar(&a.M, "m", false, "Count characters")

	flag.Parse()

	// defaults to -c -l -w
	if !a.L &&
		!a.W &&
		!a.C &&
		!a.M {
		a.L = true
		a.C = true
		a.W = true
	}

	return a
}
