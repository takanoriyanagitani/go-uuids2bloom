package bloom4

import (
	"context"
	"iter"

	u2b "github.com/takanoriyanagitani/go-uuids2bloom"

	b4 "github.com/takanoriyanagitani/go-uuids2bloom/bloom/b16"
	o4 "github.com/takanoriyanagitani/go-uuids2bloom/output/bloom/b16"
)

type UuidsToBloomToWriter4 struct {
	b4.UuidToBloom4
	o4.WriteBloom4
}

func (u UuidsToBloomToWriter4) Write(
	ctx context.Context,
	uuids iter.Seq[u2b.Pair[error, u2b.Uuid]],
) error {
	var state b4.Bloom4
	var err error
	for pair := range uuids {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if nil != pair.Left {
			return pair.Left
		}

		err = u.UuidToBloom4(ctx, pair.Right, &state)
		if nil != err {
			return err
		}
	}
	return u.WriteBloom4(ctx, &state)
}
