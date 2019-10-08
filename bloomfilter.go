/*
package bloomfilter provides probabilistic data structure called "Bloom filter".

@doc https://en.wikipedia.org/wiki/Bloom_filter
 */
package bloomfilter

import (
	"fmt"
	"hash"
	"hash/fnv"
)

//-------------------------------
// public interface
//-------------------------------

type BloomFilter struct {
	// bit vector
	bitVector []bool
	n uint // the number of elements
	m uint // the size of bit vector
	// hash functions
	k uint // the number of hash functions
	hashFn []hash.Hash64 // the k-length of hash functions
}

func New(size uint) *BloomFilter {
	hashFn := []hash.Hash64{
		fnv.New64(),
		fnv.New64(),
		fnv.New64(),
	}

	return &BloomFilter{
		bitVector: make([]bool, size),
		n: uint(0),
		m: size,
		k: 3,
		hashFn: hashFn,
	}
}

// Size returns the number of added elements in the bit vector.
func (bf *BloomFilter) Size() uint {
	return bf.n;
}

// Add data to the bit vector.
func (bf *BloomFilter) Add(item []byte) {
	hashes := bf.hash(item)

	fmt.Printf("hashes: %+v\n", hashes)

	for i := uint(0); i < bf.k; i++ {
		pos := uint(hashes[i]) % bf.m
		bf.bitVector[pos] = true
	}

	bf.n += 1
}

// Test returns true if the data is in the bit vector, false otherwise.
func (bf *BloomFilter) Test(item []byte) bool {
	hashes := bf.hash(item)

	for i := uint(0); i < bf.k; i++ {
		pos := uint(hashes[i]) % bf.m
		if !bf.bitVector[pos] {
			return false
		}
	}

	// NOTE: Bloom Filter is false-positive
	return true
}

//-------------------------------
// private interface
//-------------------------------
func (bf *BloomFilter) hash(item []byte) []uint64 {
	var values []uint64

	for _, fn := range bf.hashFn {
		fn.Write(item)
		values = append(values, fn.Sum64())
		fn.Reset()
	}

	return values
}