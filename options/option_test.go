package options

import (
	"flag"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		args     []string
		expected ArgOptions
	}{
		"No flags (default to -c -l -w)": {
			args:     []string{},
			expected: ArgOptions{L: true, W: true, C: true, M: false},
		},
		"Only -l flag": {
			args:     []string{"-l"},
			expected: ArgOptions{L: true, W: false, C: false, M: false},
		},
		"Only -w flag": {
			args:     []string{"-w"},
			expected: ArgOptions{L: false, W: true, C: false, M: false},
		},
		"Multiple flags -l -w -m": {
			args:     []string{"-l", "-w", "-m"},
			expected: ArgOptions{L: true, W: true, C: false, M: true},
		},
		"All flags -l -w -c -m": {
			args:     []string{"-l", "-w", "-c", "-m"},
			expected: ArgOptions{L: true, W: true, C: true, M: true},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			defer tearDown()

			os.Args = append([]string{"cmd"}, test.args...)

			opts := (&ArgOptions{}).Parse()

			// Check results
			if *opts != test.expected {
				t.Errorf("wrong result: got %v, want %v", opts, test.expected)
			}
		})
	}
}

func tearDown() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}
