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

	filesArgs := flag.Args()

	if len(filesArgs) < 1 {
		printRes(count.Count(os.Stdin), *op, "")
		return
	}

	totalRes := count.Counts{}
	for _, filePath := range filesArgs {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		res := count.Count(file)

		printRes(res, *op, file.Name())

		totalRes.Add(res)
	}

	if len(filesArgs) > 1 {
		printRes(totalRes, *op, "total")
	}
}

func printRes(res count.Counts, op options.ArgOptions, name string) {
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
