package bloom3

import (
	"context"
	"io"
	"os"

	b3 "github.com/takanoriyanagitani/go-uuids2bloom/bloom/b8"
)

type WriteBloom3 func(context.Context, b3.Bloom3) error

func BloomToWriterNew3(wtr io.Writer) WriteBloom3 {
	return func(_ context.Context, b b3.Bloom3) error {
		var s []uint8 = b[:]
		_, e := wtr.Write(s)
		return e
	}
}

func BloomToStdoutNew3() WriteBloom3 { return BloomToWriterNew3(os.Stdout) }
