package commonsubset

import (
	"crypto/rand"
	"testing"
	"time"

	binaryagreement "github.com/EncrypteDL/HoneyBadgerBFT-Golang/Binary_Agreement"
	reliablebroadcast "github.com/EncrypteDL/HoneyBadgerBFT-Golang/Reliable_Broadcast"
	"github.com/sirupsen/logrus"
)

func randData(len int) []byte {
	data := make([]byte, len)
	// token := make([]byte, 4)
	rand.Read(data)
	return data
}

func TestOnlyOneRBCOneABA(t *testing.T) {
	// total := 1
	// tol := 1
	aba := make([]*binaryagreement.BinaryAgreement, 1)
	aba[0] = &binaryagreement.BinaryAgreement{In: make(chan bool), Out: make(chan bool)}
	go func() {
		data := <-aba[0].In
		aba[0].Out <- data
	}()
	rbc := make([]*reliablebroadcast.ReliableBroadcast, 1)
	rbc[0] = &reliablebroadcast.ReliableBroadcast{In: make(chan []byte), Out: make(chan []byte)}
	go func() {
		data := <-rbc[0].In
		rbc[0].Out <- data
	}()

	acs := NewCommonSubset(0, 1, 1, rbc, aba)

	acs.In <- randData(4)
	logrus.Debugf("out:", <-acs.Out)
}

func TestWithNInstances(t *testing.T) {
	for k := 0; k < 10; k++ {
		total := 4
		// tol := 1
		aba := make([]*binaryagreement.BinaryAgreement, total)
		rbc := make([]*reliablebroadcast.ReliableBroadcast, total)

		for i := 0; i < total; i++ {
			aba[i] = &binaryagreement.BinaryAgreement{In: make(chan bool), Out: make(chan bool)}

			go func(i int) {
				data := <-aba[i].In
				logrus.Debugf("ABA IN : %v", data)
				aba[i].Out <- data
			}(i)

			rbc[i] = &reliablebroadcast.ReliableBroadcast{In: make(chan []byte), Out: make(chan []byte)}
			go func(i int) {
				var data []byte
				if i == 0 {
					data = <-rbc[i].In
				} else {
					data = randData(40)
				}
				if i != 3 {
					rbc[i].Out <- data
				}
			}(i)
		}

		acs := NewCommonSubset(0, 4, 1, rbc, aba)

		acs.In <- randData(4)
		time.Sleep(2 * time.Second)
		// data := <-aba[3].In
		// aba[3].Out <- data
		logrus.Debugf("out:", <-acs.Out)
	}
}
