package rdr2iter

import (
	"bufio"
	"errors"
	"io"
	"iter"
	"os"

	u2b "github.com/takanoriyanagitani/go-uuids2bloom"
)

func ReaderToUuids(rdr io.Reader) iter.Seq[u2b.Pair[error, u2b.Uuid]] {
	var br io.Reader = bufio.NewReader(rdr)
	return func(yield func(u2b.Pair[error, u2b.Uuid]) bool) {
		for {
			var buf [16]uint8
			_, e := io.ReadFull(br, buf[:])
			if nil != e {
				if errors.Is(e, io.EOF) {
					return
				}

				yield(u2b.Pair[error, u2b.Uuid]{Left: e, Right: u2b.Uuid{}})
				return
			}

			var ok bool = yield(u2b.Pair[error, u2b.Uuid]{
				Left:  nil,
				Right: buf,
			})

			if !ok {
				return
			}
		}
	}
}

func StdinToUuids() iter.Seq[u2b.Pair[error, u2b.Uuid]] {
	return ReaderToUuids(os.Stdin)
}
