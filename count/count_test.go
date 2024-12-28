package count

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

type cases struct {
	in     io.Reader
	expect Counts
}

func TestCount(t *testing.T) {
	tests := map[string]cases{
		"empty input": {
			in: strings.NewReader(""),
		},
		"two words": {
			in:     bufio.NewReader(strings.NewReader("Two Words ")),
			expect: Counts{Lines: 0, Words: 2, Bytes: 10, Characters: 10},
		},
		"single line": {
			in:     bufio.NewReader(strings.NewReader("this is a POSIX line\n")),
			expect: Counts{Lines: 1, Words: 5, Bytes: 21, Characters: 21},
		},
		"multiple lines": {
			in:     bufio.NewReader(strings.NewReader("first line\nsecond line\n")),
			expect: Counts{Lines: 2, Words: 4, Bytes: 23, Characters: 23},
		},
		"from test.txt file": {
			in: func() io.Reader {
				file, err := os.Open("../test/test.txt")
				if err != nil {
					t.Fatalf("could not open test file: %s", err)
				}
				return file
			}(),
			expect: Counts{
				Lines: 7145, Words: 58164, Bytes: 342190, Characters: 339292,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			defer tearDown(t, test)

			res := Count(test.in)

			if res != test.expect {
				t.Errorf("wrong result. expect: %v got: %v", test.expect, res)
			}
		})
	}
}

func tearDown(t *testing.T, c cases) {
	if close, ok := c.in.(io.Closer); ok {
		err := close.Close()
		if err != nil {
			t.Fatalf("could not tear down test: %s", err)
		}
	}
}
