package reliablebroadcast

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"

	ab "github.com/EncrypteDL/HoneyBadgerBFT-Golang/proto/orderer"
	"github.com/klauspost/reedsolomon"
	"github.com/sirupsen/logrus"
)

func TestReedSolman(t *testing.T) {
	total := 64
	tol := 4
	data := make([]byte, 400)
	// token := make([]byte, 4)
	rand.Read(data)
	logrus.Debugf("IN data : %v", data)
	var K = total - 2*tol
	enc, _ := reedsolomon.New(K, total-K)
	blocks, _ := Encode(enc, data)
	logrus.Debugf("Encoded : %v", blocks)
	enc.Reconstruct(blocks)
	logrus.Debugf("Decoded : %v", blocks)
	var value []byte
	for _, data := range blocks[:K] {
		logrus.Debugf("%v", data)
		value = append(value, data...)
	}
	logrus.Debugf("Out data : %v", value[:400])
	logrus.Debugf("compare : %v", bytes.Compare(data, value[:400]))

}

func TestIntialization(t *testing.T) {

	sendFunc := func(index int, msg ab.HoneyBadgerBFTMessage) {
	}
	broadcastFunc := func(msg ab.HoneyBadgerBFTMessage) {

	}
	rbcInstanceRecvMsgChannels := make(chan *ab.HoneyBadgerBFTMessage, 666666)
	instanceIndex := 0 // NOTE important to copy i
	componentSendFunc := func(index int, msg ab.HoneyBadgerBFTMessage) {
		msg.Instance = uint64(instanceIndex)
		sendFunc(index, msg)
	}
	componentBroadcastFunc := func(msg ab.HoneyBadgerBFTMessage) {
		msg.Instance = uint64(instanceIndex)
		broadcastFunc(msg)
	}
	rbc := NewReliableBroadcast(0, 4, 1, 0, 0, rbcInstanceRecvMsgChannels, componentSendFunc, componentBroadcastFunc) // TODO: a better way to eval whether i is leader, ListenAddress may difference with address in connectionlist

	if rbc != nil {
		fmt.Printf("Success")
	} else {
		fmt.Errorf("Fail")
	}
}

func TestEncodingandBroadcast(t *testing.T) {
}

func TestOneInstanceRBCChannel(t *testing.T) {
}
