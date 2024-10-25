package bloom4

import (
	"context"

	u2b "github.com/takanoriyanagitani/go-uuids2bloom"
)

const (
	// Best value when expected number of items is 32.
	NumberOfHash65536bitsDefault int = 6
)

type Bloom4 [8192]uint8 // 65536 bits

type UuidToBloom4 func(context.Context, u2b.Uuid, *Bloom4) error

// Creates 8 "hash functions"(16 bytes) from the [u2b.Uuid].
type UuidToHash4 func(context.Context, u2b.Uuid) ([8]uint16, error)

func UuidIvAsHash(_ context.Context, u u2b.Uuid) ([8]uint16, error) {
	return [8]uint16{
		uint16(u[0x00])<<8 | uint16(u[0x01]),
		uint16(u[0x02])<<8 | uint16(u[0x03]),
		uint16(u[0x04])<<8 | uint16(u[0x05]),
		uint16(u[0x06])<<8 | uint16(u[0x07]),
		uint16(u[0x08])<<8 | uint16(u[0x09]),
		uint16(u[0x0a])<<8 | uint16(u[0x0b]),
		uint16(u[0x0c])<<8 | uint16(u[0x0d]),
		uint16(u[0x0e])<<8 | uint16(u[0x0f]),
	}, nil
}

type AddHashAllToBloom4 func(context.Context, [8]uint16, *Bloom4) error

func (u UuidToHash4) ToUuidToBloom4(a AddHashAllToBloom4) UuidToBloom4 {
	return func(ctx context.Context, id u2b.Uuid, b *Bloom4) error {
		h, e := u(ctx, id)
		if nil != e {
			return e
		}
		return a(ctx, h, b)
	}
}

type AddHashToBloom4 func(context.Context, uint16, *Bloom4) error

func (a AddHashToBloom4) ToAddHashAll() AddHashAllToBloom4 {
	return func(ctx context.Context, hash [8]uint16, b *Bloom4) error {
		for _, h := range hash[:NumberOfHash65536bitsDefault] {
			err := a(ctx, h, b)
			if nil != err {
				return err
			}
		}
		return nil
	}
}

type SetBit4 func(b07 uint8, bloom uint8) uint8

func (s SetBit4) ToAddHashToBloom4() AddHashToBloom4 {
	return func(ctx context.Context, hash uint16, b *Bloom4) error {
		var shifted uint16 = hash >> 3 // up to 8191
		var b8 uint8 = b[shifted]
		var neo uint8 = s(uint8(hash&0x07), b8)
		b[shifted] = neo
		return nil
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

func (b BitIndexToUint8) ToSetBit4() SetBit4 {
	return func(b07 uint8, bloom uint8) uint8 {
		var shifted uint8 = b(b07)
		return shifted | bloom
	}
}

func BitIxToUint8(bix uint8) uint8 {
	return 1 << bix
}
