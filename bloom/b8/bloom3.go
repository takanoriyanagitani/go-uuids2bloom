package bloom3

import (
	"context"

	u2b "github.com/takanoriyanagitani/go-uuids2bloom"
)

const (
	// Best value when expected number of items is 32.
	NumberOfHash256bitsDefault int = 6
)

type Bloom3 [32]uint8 // 256 bits

type UuidToBloom3 func(context.Context, u2b.Uuid, Bloom3) (Bloom3, error)

// Creates 8 "hash functions"(8 bytes) from the [u2b.Uuid].
type UuidToHash3 func(context.Context, u2b.Uuid) ([8]uint8, error)

func UuidIvAsHash(_ context.Context, u u2b.Uuid) ([8]uint8, error) {
	return [8]uint8{
		u[0x00] ^ u[0x08],
		u[0x01] ^ u[0x09],
		u[0x02] ^ u[0x0a],
		u[0x03] ^ u[0x0b],
		u[0x04] ^ u[0x0c],
		u[0x05] ^ u[0x0d],
		u[0x06] ^ u[0x0e],
		u[0x07] ^ u[0x0f],
	}, nil
}

type AddHashAllToBloom3 func(context.Context, [8]uint8, Bloom3) (Bloom3, error)

func (u UuidToHash3) ToUuidToBloom3(a AddHashAllToBloom3) UuidToBloom3 {
	return func(ctx context.Context, id u2b.Uuid, b Bloom3) (Bloom3, error) {
		h, e := u(ctx, id)
		if nil != e {
			return b, e
		}
		return a(ctx, h, b)
	}
}

// Apply the hash.
//
//   - 0x00 .. 0x07: modify Bloom3[0x00]
//   - 0x08 .. 0x0f: modify Bloom3[0x01]
//   - 0x10 .. 0x17: modify Bloom3[0x02]
//   - ...
//   - 0xf0 .. 0xf7: modify Bloom3[0x1e]
//   - 0xf8 .. 0xff: modify Bloom3[0x1f]
type AddHashToBloom3 func(context.Context, uint8, Bloom3) (Bloom3, error)

func (a AddHashToBloom3) ToAddHashAll() AddHashAllToBloom3 {
	return func(ctx context.Context, hash [8]uint8, b Bloom3) (Bloom3, error) {
		var state Bloom3 = b
		var err error
		for _, h := range hash[:NumberOfHash256bitsDefault] {
			state, err = a(ctx, h, state)
			if nil != err {
				return b, err
			}
		}
		return state, nil
	}
}

type SetBit3 func(b07 uint8, bloom uint8) uint8

func (s SetBit3) ToAddHashToBloom3() AddHashToBloom3 {
	return func(ctx context.Context, hash uint8, b Bloom3) (Bloom3, error) {
		var shifted uint8 = hash >> 3
		var b8 uint8 = b[shifted]
		var neo uint8 = s(hash&0x07, b8)
		b[shifted] = neo
		return b, nil
	}
}

// Converts the index of bit to "flag".
//
// # Conversion Table
//
//	| input | output |
//	|:-----:|:------:|
//	|   0   | 0x01   |
//	|   1   | 0x02   |
//	|   2   | 0x04   |
//	|   3   | 0x08   |
//	|   4   | 0x10   |
//	|   5   | 0x20   |
//	|   6   | 0x40   |
//	|   7   | 0x80   |
type BitIndexToUint8 func(bix uint8) uint8

func (b BitIndexToUint8) ToSetBit3() SetBit3 {
	return func(b07 uint8, bloom uint8) uint8 {
		var shifted uint8 = b(b07)
		return shifted | bloom
	}
}

func BitIxToUint8(bix uint8) uint8 {
	return 1 << bix
}
