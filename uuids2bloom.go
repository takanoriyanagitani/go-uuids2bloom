package uuids2bloom

type Uuid [16]uint8

type Pair[L, R any] struct {
	Left  L
	Right R
}
