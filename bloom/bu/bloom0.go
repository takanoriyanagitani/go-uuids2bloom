package bunit

type BloomUnit uint8

func (u BloomUnit) IsFull() bool {
	return 0xff == u
}
