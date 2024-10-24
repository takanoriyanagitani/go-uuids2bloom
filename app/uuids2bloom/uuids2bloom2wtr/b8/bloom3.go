package bloom3

import (
	"context"
	"iter"

	u2b "github.com/takanoriyanagitani/go-uuids2bloom"

	b3 "github.com/takanoriyanagitani/go-uuids2bloom/bloom/b8"
	o3 "github.com/takanoriyanagitani/go-uuids2bloom/output/bloom/b8"
)

type UuidsToBloomToWriter3 struct {
	b3.UuidToBloom3
	o3.WriteBloom3
}

func (u UuidsToBloomToWriter3) Write(
	ctx context.Context,
	uuids iter.Seq[u2b.Pair[error, u2b.Uuid]],
) error {
	var state b3.Bloom3
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

		if state.IsFull() {
			// no bits can be added anymore
			return u.WriteBloom3(ctx, state)
		}

		state, err = u.UuidToBloom3(ctx, pair.Right, state)
		if nil != err {
			return err
		}
	}
	return u.WriteBloom3(ctx, state)
}
