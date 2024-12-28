package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/flan6/gwc/count"
	"github.com/flan6/gwc/options"
)

func main() {
	op := (&options.ArgOptions{}).Parse()

	files := flag.Args()

	var res count.Counts
	name := ""

	if len(files) > 0 {
		file, err := os.Open(files[0])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		name = file.Name()
		res = count.Count(file)
	} else {
		res = count.Count(os.Stdin)
	}

	out := ""

	if op.L {
		out += fmt.Sprintf("%d\t", res.Lines)
	}

	if op.W {
		out += fmt.Sprintf("%d\t", res.Words)
	}

	if op.C {
		out += fmt.Sprintf("%d\t", res.Bytes)
	}

	if op.M {
		out += fmt.Sprintf("%d\t", res.Characters)
	}

	fmt.Printf("%s%s\n", out, name)
}
