package tools

import (
	"crypto/sha256"
	"log"
	"math"

	"github.com/klauspost/reedsolomon"
)

func Encode(enc reedsolomon.Encoder, data []byte) ([][]byte, error) {
	shards, err := enc.Split(data)
	if err != nil {
		log.Panic(err)
	}

	if err := enc.Encode(shards); err != nil {
		log.Panic(err)
	}
	return shards, nil
}

func merkleTreeHash(data []byte, others ...[]byte) []byte {
	//note: index root at 1
	s := sha256.New()
	s.Write(data)
	for _, d := range others {
		s.Write(d)
	}
	return s.Sum(nil)
}

func newMerkeleTree(blocks [][]byte) [][]byte {
	bottomRow := int(math.Pow(2, math.Ceil(math.Log2(float64(len(blocks))))))
	result := make([][]byte, 2*bottomRow, 2*bottomRow)
	for i := 0; i < len(blocks); i++ {
		result[bottomRow+i] = merkleTreeHash(blocks[i])
	}
	for i := bottomRow - 1; i > 0; i-- {
		result[i] = merkleTreeHash(result[i*2], result[i*2+1])
	}
	return result
}

func getMerkleTreeBranch(tree [][]byte, index int) (result [][]byte) {
	//note: index from 0, block index not tree item index
	t := index + (len(tree) >> 1)
	for t > 1 {
		result = append(result, tree[t^1])
		t /= 2
	}
	return result
}

func verifyMerkleTree(rootHash []byte, branch [][]byte, block []byte, index int) bool {
	//TODO: add checks
	tmp := merkleTreeHash(block)
	for _, node := range branch {
		if index&1 == 0 {
			tmp = merkleTreeHash(tmp, node)
		} else {
			tmp = merkleTreeHash(node, tmp)
		}
		index /= 2
	}
	return string(rootHash) == string(tmp)
}
