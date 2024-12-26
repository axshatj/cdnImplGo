package main

import (
	"hash/fnv"
	"sort"
	"sync"
	"fmt"
)

type ConsistentHash struct {
	nodes       map[uint32]string
	sortedKeys  []uint32
	virtualNodes int
	mutex       sync.RWMutex
}

func NewConsistentHash(nodes []string, virtualNodes int) *ConsistentHash {
	ch := &ConsistentHash{
		nodes:       make(map[uint32]string),
		virtualNodes: virtualNodes,
	}
	for _, node := range nodes {
		ch.AddNode(node)
	}
	return ch
}

func (ch *ConsistentHash) AddNode(node string) {
	ch.mutex.Lock()
	defer ch.mutex.Unlock()

	for i := 0; i < ch.virtualNodes; i++ {
		key := ch.hash(fmt.Sprintf("%s:%d", node, i))
		ch.nodes[key] = node
	}
	ch.updateSortedKeys()
}

func (ch *ConsistentHash) RemoveNode(node string) {
	ch.mutex.Lock()
	defer ch.mutex.Unlock()

	for i := 0; i < ch.virtualNodes; i++ {
		key := ch.hash(fmt.Sprintf("%s:%d", node, i))
		delete(ch.nodes, key)
	}
	ch.updateSortedKeys()
}

func (ch *ConsistentHash) GetNode(key string) string {
	ch.mutex.RLock()
	defer ch.mutex.RUnlock()

	if len(ch.sortedKeys) == 0 {
		return ""
	}

	hash := ch.hash(key)
	idx := sort.Search(len(ch.sortedKeys), func(i int) bool {
		return ch.sortedKeys[i] >= hash
	})

	if idx == len(ch.sortedKeys) {
		idx = 0
	}

	return ch.nodes[ch.sortedKeys[idx]]
}

func (ch *ConsistentHash) hash(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (ch *ConsistentHash) updateSortedKeys() {
	ch.sortedKeys = make([]uint32, 0, len(ch.nodes))
	for k := range ch.nodes {
		ch.sortedKeys = append(ch.sortedKeys, k)
	}
	sort.Slice(ch.sortedKeys, func(i, j int) bool {
		return ch.sortedKeys[i] < ch.sortedKeys[j]
	})
}

