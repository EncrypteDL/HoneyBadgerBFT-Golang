package honeybadgerblock

import (
	ab "github.com/EncrypteDL/HoneyBadgerBFT-Golang/proto/orderer"
)

type HoneyBadgerBlock struct {
	total        int
	maxMalicious int
	channel      chan *ab.HoneyBadgerBFTMessage
}
