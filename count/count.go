package count

import (
	"bufio"
	"io"
	"log"
	"unicode"
)

// Represents the
type Counts struct {
	Lines      int64
	Words      int64
	Bytes      int64
	Characters int64
}

func Count(in io.Reader) Counts {
	source := bufio.NewReader(in)

	res := Counts{}
	var prev rune

	for {
		r, i, err := source.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		res.Bytes += int64(i)
		res.Characters++

		if unicode.IsSpace(r) && !unicode.IsSpace(prev) {
			res.Words++
		}

		if r == '\n' {
			res.Lines++
		}

		prev = r
	}

	return res
}
