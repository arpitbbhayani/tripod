package test_tripod

import (
	"container/list"
	"github.com/arpitbbhayani/tripod"
	"math/rand"
	"testing"
)

var dataset = map[string]bool{
	"test":    true,
	"te":      true,
	"test123": true,
}

var prefixDataset = map[string]bool{
	"test":    true,
	"test123": true,
}

func getRandomByteSlice(size int) []byte {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, size)
	for i := 0; i < size; i++ {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func populatePrefixStoreByteTrie(tr *tripod.PrefixStoreByteTrie) {
	for key, _ := range dataset {
		tr.Put([]byte(key))
	}
}

func TestPrefixStoreByteTriePut(t *testing.T) {
	tr := tripod.CreatePrefixStoreByteTrie(128)

	if _, err := tr.Put(make([]byte, 129, 129)); err == nil {
		t.Errorf("adding data more than maxSize specified should return an error")
	}

	if newlyAdded, _ := tr.Put([]byte("")); newlyAdded == true {
		t.Errorf("adding empty string to trie, expected %t", false)
	}

	if newlyAdded, _ := tr.Put([]byte("test")); newlyAdded == false {
		t.Errorf("adding key to trie: expected %t", true)
	}

	if newlyAdded, _ := tr.Put([]byte("test")); newlyAdded == true {
		t.Errorf("readding same key to trie: expected %t", false)
	}

	if newlyAdded, _ := tr.Put([]byte("te")); newlyAdded == false {
		t.Errorf("adding key to trie: expected %t", true)
	}

	if newlyAdded, _ := tr.Put([]byte("test123")); newlyAdded == false {
		t.Errorf("adding key to trie: expected %t", true)
	}

	// Testing on huge random data
	hugeDataset := make(map[string]bool)
	for i := 0; i < 20000; i++ {
		x := getRandomByteSlice(1 + rand.Intn(127))
		hugeDataset[string(x)] = true
		tr.Put([]byte(x))
	}

	for key, _ := range hugeDataset {
		if tr.Exists([]byte(key)) != true {
			t.Errorf("key %s should be there in the PrefixStore %s", key)
		}
	}

}

func TestPrefixStoreByteTrieExists(t *testing.T) {
	tr := tripod.CreatePrefixStoreByteTrie(128)
	populatePrefixStoreByteTrie(tr)

	if isPresent := tr.Exists([]byte("")); isPresent == true {
		t.Errorf("fetching empty key from trie should return false")
	}

	if isPresent := tr.Exists([]byte("doesnotexist")); isPresent == true {
		t.Errorf("fetching non-existent key from trie should return %t", false)
	}

	if isPresent := tr.Exists([]byte("tes")); isPresent == true {
		t.Errorf("fetching non-existent key but for which path exists should return %t", false)
	}

	if isPresent := tr.Exists([]byte("test")); isPresent == false {
		t.Errorf("fetching existing key from trie should return %t", true)
	}

	if isPresent := tr.Exists([]byte("te")); isPresent == false {
		t.Errorf("fetching existing key from trie should return %t", true)
	}

	if isPresent := tr.Exists([]byte("test123")); isPresent == false {
		t.Errorf("fetching existing key from trie should return %t", true)
	}

	if isPresent := tr.Exists([]byte("test123456")); isPresent == true {
		t.Errorf("fetching trie for non-existent key whose prefix exists in trie should return %t", false)
	}
}

func TestPrefixStoreByteTriePrefixSearch(t *testing.T) {
	tr := tripod.CreatePrefixStoreByteTrie(128)
	populatePrefixStoreByteTrie(tr)

	if count := tr.PrefixSearch(make([]byte, 129, 129)).Len(); count != 0 {
		t.Errorf("prefix search for a prefix > maxSize should return empty list, but it returned %d", count)
	}

	if count := tr.PrefixSearch([]byte("doesnotexist")).Len(); count != 0 {
		t.Errorf("prefixsearch for path that does not exist should return empty list, but it returned %d", count)
	}

	if count := tr.PrefixSearch([]byte("tesghikl")).Len(); count != 0 {
		t.Errorf("prefixsearch for path that does not exist should return empty list, but it returned %d", count)
	}

	if count := tr.PrefixSearch([]byte("tes")).Len(); count != 2 {
		t.Errorf("expected number of elements in trie for given prefix are %d, but there are %d elements", 2, count)
	}

	if count := tr.PrefixSearch([]byte("test123")).Len(); count != 1 {
		t.Errorf("expected number of elements in trie for given prefix are %d, but there are %d elements", 1, count)
	}

	if count := tr.PrefixSearch([]byte("test1234")).Len(); count != 0 {
		t.Errorf("expected number of elements in trie for given prefix are %d, but there are %d elements", 0, count)
	}

	var results *list.List

	results = tr.PrefixSearch([]byte(""))
	if count := results.Len(); count != 3 {
		t.Errorf("expected elements in trie are %d, but there are %d elements", 3, count)
	}
	for e := results.Front(); e != nil; e = e.Next() {
		val := string(e.Value.([]byte))
		if dataset[val] == false {
			t.Errorf("improper value %d retrieved from trie during full trie PrefixSearch", val)
		}
	}

	results = tr.PrefixSearch([]byte("te"))
	if count := results.Len(); count != 3 {
		t.Errorf("expected elements in trie are %d, but there are %d elements", 3, count)
	}
	for e := results.Front(); e != nil; e = e.Next() {
		val := string(e.Value.([]byte))
		if dataset[val] == false {
			t.Errorf("improper value %d retrieved from trie during full trie PrefixSearch", val)
		}
	}

	results = tr.PrefixSearch([]byte("tes"))
	if count := results.Len(); count != 2 {
		t.Errorf("expected elements in trie are %d, but there are %d elements", 2, count)
	}
	for e := results.Front(); e != nil; e = e.Next() {
		val := string(e.Value.([]byte))
		if prefixDataset[val] == false {
			t.Errorf("improper value %d retrieved from trie during full trie PrefixSearch", val)
		}
	}

	results = tr.PrefixSearch([]byte("test123"))
	if count := results.Len(); count != 1 {
		t.Errorf("expected elements in trie are %d, but there are %d elements", 1, count)
	}
	for e := results.Front(); e != nil; e = e.Next() {
		val := string(e.Value.([]byte))
		if prefixDataset[val] == false {
			t.Errorf("improper value %d retrieved from trie during full trie PrefixSearch", val)
		}
	}

	results = tr.PrefixSearch([]byte("test12345"))
	if count := results.Len(); count != 0 {
		t.Errorf("expected elements in trie are %d, but there are %d elements", 0, count)
	}
}
