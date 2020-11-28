package bloomfilter

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestBloomFunc(t *testing.T) {
	var elemNum uint64 = 10000000000
	// 99.99%
	errRate := 0.00001
	bloomBitSize := CalBloomSize(elemNum, errRate)

	fmt.Println("bloom filter size in bit:", bloomBitSize)
	fmt.Println("bloom filter size in Mbyte:", bloomBitSize/8/1000/1000)

	hashFuncNum := CalHashFuncNum(elemNum, bloomBitSize)
	fmt.Println("bloom filter Hash Func Num", hashFuncNum)

	errRate = CalErrRate(elemNum, bloomBitSize, hashFuncNum)
	fmt.Println("Error Rate", errRate)
}

func TestBloomFilter(t *testing.T) {
	file, err := os.Open("word-list.txt")
	if err != nil {
		t.Error(err)
		return
	}

	defer file.Close()

	var wordList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}

	var elemNum uint64 = uint64(len(wordList))
	errRate := 0.0001

	bloomBitSize := CalBloomSize(elemNum, errRate)
	hashFuncNum := CalHashFuncNum(elemNum, bloomBitSize)

	bloomFilter := NewBloomFilter(elemNum, bloomBitSize, hashFuncNum, errRate)
	bloomFilter.Init()
	for _, elem := range wordList {
		bloomFilter.Add([]byte(elem))
	}

	var testCase = []struct {
		Elem   string
		Result bool
	}{
		{"hello", false},
		{"He", true},
		{"make", false},
		{"great", true},
		{"again", false},
		{"yesterday", true},
	}

	for _, oneCase := range testCase {
		got := bloomFilter.Contains([]byte(oneCase.Elem))
		fmt.Printf("word list contains %v : %v \n", oneCase.Elem, got)
		if got != oneCase.Result {
			if got {
				t.Error("Contains", oneCase.Elem)
			} else {
				t.Error("should not Contains", oneCase.Elem)
			}
			t.Error("got != result")
			return
		}
	}

}
