package bloom4

import (
	"context"
	"io"
	"os"

	b4 "github.com/takanoriyanagitani/go-uuids2bloom/bloom/b16"
)

type WriteBloom4 func(context.Context, *b4.Bloom4) error

func BloomToWriterNew4(wtr io.Writer) WriteBloom4 {
	return func(_ context.Context, b *b4.Bloom4) error {
		var s []uint8 = b[:]
		_, e := wtr.Write(s)
		return e
	}
}

func BloomToStdoutNew4() WriteBloom4 { return BloomToWriterNew4(os.Stdout) }
