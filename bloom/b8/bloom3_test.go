package bloom3_test

import (
	"context"
	"testing"

	b3 "github.com/takanoriyanagitani/go-uuids2bloom/bloom/b8"
)

func TestBloom3(t *testing.T) {
	t.Parallel()

	t.Run("BitIxToUint8", func(t *testing.T) {
		t.Parallel()

		var bix2u8 b3.BitIndexToUint8 = b3.BitIxToUint8

		cases := [8]uint16{
			0x0001,
			0x0102,
			0x0204,
			0x0308,
			0x0410,
			0x0520,
			0x0640,
			0x0780,
		}

		for _, cs := range cases {
			var pair uint16 = cs
			var i uint16 = pair >> 8
			var o uint16 = pair & 0xff

			var got uint8 = bix2u8(uint8(i))
			if got != uint8(o) {
				t.Fatalf("expected: %v, got: %v\n", o, got)
			}
		}
	})

	t.Run("ToSetBit3", func(t *testing.T) {
		t.Parallel()

		var bix2u8 b3.BitIndexToUint8 = b3.BitIxToUint8
		var sb3 b3.SetBit3 = bix2u8.ToSetBit3()

		t.Run("zero", func(t *testing.T) {
			t.Parallel()

			var result uint8 = sb3(0, 0)
			if 0x01 != result {
				t.Fatalf("expected: %v, got: %v\n", 0x01, result)
			}
		})

		t.Run("sky", func(t *testing.T) {
			t.Parallel()

			var result uint8 = sb3(0x06, 0x34)
			if 0x74 != result {
				t.Fatalf("expected: %v, got: %v\n", 0x01, result)
			}
		})

		t.Run("tokyo", func(t *testing.T) {
			t.Parallel()

			var result uint8 = sb3(0x03, 0x33)
			if 0x3b != result {
				t.Fatalf("expected: %v, got: %v\n", 0x01, result)
			}
		})
	})

	t.Run("ToAddHashToBloom3", func(t *testing.T) {
		t.Parallel()

		var bix2u8 b3.BitIndexToUint8 = b3.BitIxToUint8
		var sb3 b3.SetBit3 = bix2u8.ToSetBit3()
		var ah2b3 b3.AddHashToBloom3 = sb3.ToAddHashToBloom3()

		t.Run("zero", func(t *testing.T) {
			t.Parallel()

			var bl3 b3.Bloom3
			neo, e := ah2b3(context.Background(), 0, bl3)
			if nil != e {
				t.Fatalf("unexpected error: %v\n", e)
			}

			for i := 1; i < 32; i++ {
				var b8 uint8 = neo[i]
				if 0 != b8 {
					t.Fatalf("expected: %v, got: %v\n", 0, b8)
				}
			}

			if 0x01 != neo[0] {
				t.Fatalf("expected: %v, got: %v\n", 0x01, neo[0])
			}
		})

		t.Run("42", func(t *testing.T) {
			t.Parallel()

			var bl3 b3.Bloom3
			neo, e := ah2b3(context.Background(), 0x42, bl3)
			if nil != e {
				t.Fatalf("unexpected error: %v\n", e)
			}

			for i := 0; i < 32; i++ {
				if 8 == i {
					continue
				}
				var b8 uint8 = neo[i]
				if 0 != b8 {
					t.Fatalf("expected: %v, got: %v\n", 0, b8)
				}
			}

			if 4 != neo[8] {
				t.Fatalf("expected: %v, got: %v\n", 4, neo[8])
			}
		})
	})
}
