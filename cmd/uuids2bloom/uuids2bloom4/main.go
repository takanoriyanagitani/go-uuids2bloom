package main

import (
	"context"
	"iter"
	"log"

	u2b "github.com/takanoriyanagitani/go-uuids2bloom"

	ui "github.com/takanoriyanagitani/go-uuids2bloom/input/uuid/binary/iter"

	b4 "github.com/takanoriyanagitani/go-uuids2bloom/bloom/b16"
	o4 "github.com/takanoriyanagitani/go-uuids2bloom/output/bloom/b16"

	a4 "github.com/takanoriyanagitani/go-uuids2bloom/app/uuids2bloom/uuids2bloom2wtr/b16"
)

var bix2u8 b4.BitIndexToUint8 = b4.BitIxToUint8
var sb4 b4.SetBit4 = bix2u8.ToSetBit4()
var ah2b4 b4.AddHashToBloom4 = sb4.ToAddHashToBloom4()
var aha2b4 b4.AddHashAllToBloom4 = ah2b4.ToAddHashAll()

var uuid2hash4 b4.UuidToHash4 = b4.UuidIvAsHash
var u2b4 b4.UuidToBloom4 = uuid2hash4.ToUuidToBloom4(aha2b4)

var bloom2writer o4.WriteBloom4 = o4.BloomToStdoutNew4()

var uuids iter.Seq[u2b.Pair[error, u2b.Uuid]] = ui.StdinToUuids()

var u2b2wtr a4.UuidsToBloomToWriter4 = a4.UuidsToBloomToWriter4{
	UuidToBloom4: u2b4,
	WriteBloom4:  bloom2writer,
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
