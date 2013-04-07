package bloom

import "hash/fnv"
import "github.com/reusee/mmh3"

/*
32bit fnv algorithm hash
*/
func fnvHash(str string) uint32 {
    h := fnv.New32()
    h.Write([]byte(str))
    return h.Sum32()
}

/*
32bit murmur algorithm hash
*/
func murmurHash(str string) uint32 {
    return mmh3.Hash32([]byte(str))
}

/*
holds the bitvector made of 64bit unsigned ints
a number of hash functions and salts
*/
type Filter struct {
    bits []uint64
    hashes []func(string) uint32
    salts []string
}

/*
make a new bloom filter
size is number of 64bit blocks to make bitvector
*/
func New(size int) *Filter {
    hashes := []func(string) uint32 {fnvHash, murmurHash}
    salts := []string {"salt", "pepper"}
    return &Filter{make([]uint64,size), hashes, salts}
}

/*
add a string to set
*/
func (f *Filter) Add(str string) {
    size := uint32(len(f.bits) * 64)
    for _, hash := range f.hashes {
        f.set(size, hash, str)
        for _, salt := range f.salts {
            f.set(size, hash, salt + str)
        }
    }
}

/*
sets bit that resulted from hash of str
*/
func (f *Filter) set(size uint32, hash func(string) uint32, str string) {
    index := hash(str) % size
    f.bits[index/64] |= 1 << (index%64)
}

/*
returns false if str isn't in set
*/
func (f *Filter) Contains(str string) bool {
    size := uint32(len(f.bits) * 64)
    for _, hash := range f.hashes {
        if !f.isSet(size, hash, str) {
            return false
        }
        for _, salt := range f.salts {
            if !f.isSet(size, hash, salt + str) {
                return false
            }
        }
    }
    return true
}

/*
return true if bit result from hash of str is set
*/
func (f *Filter) isSet(size uint32, hash func(string) uint32, str string) bool {
    index := hash(str) % size
    return (f.bits[index/64] & (1 << (index%64))) != 0
}
