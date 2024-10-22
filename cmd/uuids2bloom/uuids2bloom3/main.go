package main

import (
	"context"
	"iter"
	"log"

	u2b "github.com/takanoriyanagitani/go-uuids2bloom"

	ui "github.com/takanoriyanagitani/go-uuids2bloom/input/uuid/binary/iter"

	b3 "github.com/takanoriyanagitani/go-uuids2bloom/bloom/b8"
	o3 "github.com/takanoriyanagitani/go-uuids2bloom/output/bloom/b8"

	a3 "github.com/takanoriyanagitani/go-uuids2bloom/app/uuids2bloom/uuids2bloom2wtr/b8"
)

var bix2u8 b3.BitIndexToUint8 = b3.BitIxToUint8
var sb3 b3.SetBit3 = bix2u8.ToSetBit3()
var ah2b3 b3.AddHashToBloom3 = sb3.ToAddHashToBloom3()
var aha2b3 b3.AddHashAllToBloom3 = ah2b3.ToAddHashAll()

var uuid2hash3 b3.UuidToHash3 = b3.UuidIvAsHash
var u2b3 b3.UuidToBloom3 = uuid2hash3.ToUuidToBloom3(aha2b3)

var bloom2writer o3.WriteBloom3 = o3.BloomToStdoutNew3()

var uuids iter.Seq[u2b.Pair[error, u2b.Uuid]] = ui.StdinToUuids()

var u2b2wtr a3.UuidsToBloomToWriter3 = a3.UuidsToBloomToWriter3{
	UuidToBloom3: u2b3,
	WriteBloom3:  bloom2writer,
}

func main() {
	e := u2b2wtr.Write(
		context.Background(),
		uuids,
	)
	if nil != e {
		log.Printf("%v\n", e)
	}
}
