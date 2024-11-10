package commonsubset

import (
	"crypto/rand"
	"testing"
)

func randData(len int) []byte {
	data := make([]byte, len)
	rand.Read(data)
	return data
}

func TestOnlyOneRBCOneABA(t *testing.T) {
	aba := make([]*binaryagreement, 1)
	aba[0] = &binaryagreement{In: make(chan bool), Out: make(chan bool)}
	go func() {
		data := <-aba[0].In
		aba[0].Out <- data
	}()

	rbc := make([]*ReliableBroadcast, 1)
	rbc[0] = &ReliableBroadcast{In: make(chan []byte), Out: make(chan []byte)}
}
