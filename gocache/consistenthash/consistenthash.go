package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash           //Hash functions used for calculation
	keys     []int          //array storing all hash keys
	mapping  map[int]string //map the relationship between every virtual node (key) and its corresponding real node's name(value)
	replicas int            //the number of virtual nodes each real node has
}

func NewMap(rep int, hashfunc Hash) *Map {
	m := &Map{
		replicas: rep,
		hash:     hashfunc,
		mapping:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add adds some keys to the hash. Keys are the real nodes been added to the circle (as values but not keys in the map)
func (m *Map) Add(keys ...string) {
	for _, key := range keys { //traverse the newly added nodes one by one
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key))) //Adding the strings representing two different number to integrate into one single string representing the new number
			m.keys = append(m.keys, hash)
			m.mapping[hash] = key
		}
	}
	sort.Ints(m.keys)
}

// Get gets the closest item in the hash to the provided key
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	return m.mapping[m.keys[idx%len(m.keys)]]
}
