package bloomfilter

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"math"
	"math/big"

	"github.com/willf/bitset"
)

// BloomFilter 布隆过滤器
type BloomFilter struct {
	ElemNum      uint64
	ErrRate      float64
	HashFuncNum  uint64
	BloomBitSize uint64

	bitMap *bitset.BitSet
	keys   map[uint32]bool
}

// NewBloomFilter 初始化布隆过滤器实例
func NewBloomFilter(elemNum, bloomBitSize, hashFuncNum uint64, errRate float64) *BloomFilter {
	return &BloomFilter{
		ElemNum:      elemNum,
		BloomBitSize: bloomBitSize,
		HashFuncNum:  hashFuncNum,
		ErrRate:      errRate,
	}
}

// CalBloomSize 计算布隆过滤器位图大小
// elemNum: 元素个数
// errRate: 误判率
func CalBloomSize(elemNum uint64, errRate float64) uint64 {
	var bloomBitSize = float64(elemNum) * math.Log(errRate) / (math.Log(2) * math.Log(2)) * (-1)
	return uint64(math.Ceil(bloomBitSize))
}

// CalHashFuncNum 计算需要的hash函数个数
func CalHashFuncNum(elemNum, bloomBitSize uint64) uint64 {
	var k = math.Log(2) * float64(bloomBitSize) / float64(elemNum)
	return uint64(math.Ceil(k))
}

// CalErrRate 计算布隆过滤器的误判率
func CalErrRate(elemNum, bloomBitSize, hashFuncNum uint64) float64 {
	var y = float64(elemNum) * float64(hashFuncNum) / float64(bloomBitSize)
	return math.Pow(
		float64(1)-math.Pow(math.E, y*float64(-1)),
		float64(hashFuncNum),
	)
}

// Init 初始化布隆过滤器
func (bf *BloomFilter) Init() {
	// 分配布隆过滤器bitmap
	bf.bitMap = bitset.New(uint(bf.BloomBitSize))

	bf.keys = make(map[uint32]bool)
	for uint64(len(bf.keys)) < bf.HashFuncNum {
		randNum, err := rand.Int(rand.Reader, new(big.Int).SetUint64(math.MaxUint32))
		if err != nil {
			panic(err)
		}

		bf.keys[uint32(randNum.Uint64())] = true
	}
}

// Add 将新元素添加到布隆过滤器
func (bf *BloomFilter) Add(elem []byte) {
	var buf [4]byte
	for k := range bf.keys {
		binary.LittleEndian.PutUint32(buf[:], k)
		hashResult := new(big.Int).SetBytes(HMACWithSHA128(elem, buf[:]))
		result := hashResult.Mod(hashResult, big.NewInt(int64(bf.BloomBitSize)))
		bf.bitMap.Set(uint(result.Uint64()))
	}
}

// Contains 判断元素是否命中过滤器
func (bf *BloomFilter) Contains(elem []byte) bool {
	var buf [4]byte
	for k := range bf.keys {
		binary.LittleEndian.PutUint32(buf[:], k)
		hashResult := new(big.Int).SetBytes(HMACWithSHA128(elem, buf[:]))
		result := hashResult.Mod(hashResult, big.NewInt(int64(bf.BloomBitSize)))

		// 查询result对应的bit是否为1
		if !bf.bitMap.Test(uint(result.Uint64())) {
			return false
		}
	}

	return true
}

// HMACWithSHA128 哈希函数
// 通过不同的盐，对同一个输入数据，同一个哈希函数，产生多个不同的哈希值
func HMACWithSHA128(seed, key []byte) []byte {
	hmac512 := hmac.New(sha1.New, key)
	hmac512.Write(seed)
	return hmac512.Sum(nil)
}

// https://blog.csdn.net/liuzhijun301/article/details/83040178
// https://blog.csdn.net/maxdaic/article/details/106977683
// https://zhuanlan.zhihu.com/p/148784654
