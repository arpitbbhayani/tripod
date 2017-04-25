package test_tripod

import (
	"container/list"
	"github.com/arpitbbhayani/tripod"
	"math/rand"
	"testing"
)

func getRandomUTF8RuneSlice(size int) []rune {
	letterRunes := []rune("aãbãcãdãefghãiãjãkãlmãnopãqrãsããtuãvwãxãyz")
	b := make([]rune, size)
	for i := 0; i < size; i++ {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return b
}

func populatePrefixStoreRuneTrie(tr *tripod.PrefixStoreRuneTrie) {
	for key, _ := range dataset {
		tr.Put([]rune(key))
	}
}

func TestPrefixStoreRuneTriePut(t *testing.T) {
    tr := tripod.CreatePrefixStoreRuneTrie(2)
    if _, err := tr.Put([]rune("aa")); err != nil {
		t.Errorf("adding 2 ascii characters in 2 runes long string should be allowed")
	}

    if _, err := tr.Put([]rune("aaa")); err == nil {
		t.Errorf("adding 3 ascii characters in 3 runes long string should not be allowed")
	}

    if _, err := tr.Put([]rune("aã")); err != nil {
		t.Errorf("adding 1 ascii and 1 utf8 characters in 2 runes long string should be allowed")
	}

    if _, err := tr.Put([]rune("ãã")); err != nil {
		t.Errorf("adding 2 utf8 characters in 2 runes long string should be allowed")
	}

    if _, err := tr.Put([]rune("ããa")); err == nil {
		t.Errorf("adding 1 ascii 2 utf8 characters in 2 runes long string should not be allowed")
	}

    if _, err := tr.Put([]rune("ããã")); err == nil {
		t.Errorf("adding 3 utf8 characters in 2 runes long string should not be allowed")
	}

	tr = tripod.CreatePrefixStoreRuneTrie(128)

	if _, err := tr.Put(getRandomUTF8RuneSlice(129)); err == nil {
		t.Errorf("adding data more than maxSize specified should return an error")
	}

	if newlyAdded, _ := tr.Put([]rune("")); newlyAdded == true {
		t.Errorf("adding empty string to trie, expected %t", false)
	}

	if newlyAdded, _ := tr.Put([]rune("test")); newlyAdded == false {
		t.Errorf("adding key to trie: expected %t", true)
	}

	if newlyAdded, _ := tr.Put([]rune("test")); newlyAdded == true {
		t.Errorf("readding same key to trie: expected %t", false)
	}

	if newlyAdded, _ := tr.Put([]rune("te")); newlyAdded == false {
		t.Errorf("adding key to trie: expected %t", true)
	}

	if newlyAdded, _ := tr.Put([]rune("test123")); newlyAdded == false {
		t.Errorf("adding key to trie: expected %t", true)
	}

	// Testing on huge random data
	hugeDataset := make(map[string]bool)
	for i := 0; i < 20000; i++ {
		x := getRandomUTF8RuneSlice(1 + rand.Intn(127))
		hugeDataset[string(x)] = true
		tr.Put(x)
	}

	for key, _ := range hugeDataset {
		if tr.Exists([]rune(key)) != true {
			t.Errorf("key %s should be there in the PrefixStore %s", key)
		}
	}

}

func TestPrefixStoreRuneTrieExists(t *testing.T) {
	tr := tripod.CreatePrefixStoreRuneTrie(128)
	populatePrefixStoreRuneTrie(tr)

	if isPresent := tr.Exists([]rune("")); isPresent == true {
		t.Errorf("fetching empty key from trie should return false")
	}

	if isPresent := tr.Exists([]rune("doesnotexist")); isPresent == true {
		t.Errorf("fetching non-existent key from trie should return %t", false)
	}

	if isPresent := tr.Exists([]rune("tes")); isPresent == true {
		t.Errorf("fetching non-existent key but for which path exists should return %t", false)
	}

	if isPresent := tr.Exists([]rune("test")); isPresent == false {
		t.Errorf("fetching existing key from trie should return %t", true)
	}

	if isPresent := tr.Exists([]rune("te")); isPresent == false {
		t.Errorf("fetching existing key from trie should return %t", true)
	}

	if isPresent := tr.Exists([]rune("test123")); isPresent == false {
		t.Errorf("fetching existing key from trie should return %t", true)
	}

	if isPresent := tr.Exists([]rune("test123456")); isPresent == true {
		t.Errorf("fetching trie for non-existent key whose prefix exists in trie should return %t", false)
	}
}

func TestPrefixStoreRuneTriePrefixSearch(t *testing.T) {
	tr := tripod.CreatePrefixStoreRuneTrie(128)
	populatePrefixStoreRuneTrie(tr)

	if count := tr.PrefixSearch(getRandomUTF8RuneSlice(129)).Len(); count != 0 {
		t.Errorf("prefix search for a prefix > maxSize should return empty list, but it returned %d", count)
	}

	if count := tr.PrefixSearch([]rune("doesnotexist")).Len(); count != 0 {
		t.Errorf("prefixsearch for path that does not exist should return empty list, but it returned %d", count)
	}

	if count := tr.PrefixSearch([]rune("tesghikl")).Len(); count != 0 {
		t.Errorf("prefixsearch for path that does not exist should return empty list, but it returned %d", count)
	}

	if count := tr.PrefixSearch([]rune("tes")).Len(); count != 2 {
		t.Errorf("expected number of elements in trie for given prefix are %d, but there are %d elements", 2, count)
	}

	if count := tr.PrefixSearch([]rune("test123")).Len(); count != 1 {
		t.Errorf("expected number of elements in trie for given prefix are %d, but there are %d elements", 1, count)
	}

	if count := tr.PrefixSearch([]rune("test1234")).Len(); count != 0 {
		t.Errorf("expected number of elements in trie for given prefix are %d, but there are %d elements", 0, count)
	}

	var results *list.List

	results = tr.PrefixSearch([]rune(""))
	if count := results.Len(); count != 3 {
		t.Errorf("expected elements in trie are %d, but there are %d elements", 3, count)
	}
	for e := results.Front(); e != nil; e = e.Next() {
		val := string(e.Value.([]rune))
		if dataset[val] == false {
			t.Errorf("improper value %d retrieved from trie during full trie PrefixSearch", val)
		}
	}

	results = tr.PrefixSearch([]rune("te"))
	if count := results.Len(); count != 3 {
		t.Errorf("expected elements in trie are %d, but there are %d elements", 3, count)
	}
	for e := results.Front(); e != nil; e = e.Next() {
		val := string(e.Value.([]rune))
		if dataset[val] == false {
			t.Errorf("improper value %d retrieved from trie during full trie PrefixSearch", val)
		}
	}

	results = tr.PrefixSearch([]rune("tes"))
	if count := results.Len(); count != 2 {
		t.Errorf("expected elements in trie are %d, but there are %d elements", 2, count)
	}
	for e := results.Front(); e != nil; e = e.Next() {
		val := string(e.Value.([]rune))
		if prefixDataset[val] == false {
			t.Errorf("improper value %d retrieved from trie during full trie PrefixSearch", val)
		}
	}

	results = tr.PrefixSearch([]rune("test123"))
	if count := results.Len(); count != 1 {
		t.Errorf("expected elements in trie are %d, but there are %d elements", 1, count)
	}
	for e := results.Front(); e != nil; e = e.Next() {
		val := string(e.Value.([]rune))
		if prefixDataset[val] == false {
			t.Errorf("improper value %d retrieved from trie during full trie PrefixSearch", val)
		}
	}

	results = tr.PrefixSearch([]rune("test12345"))
	if count := results.Len(); count != 0 {
		t.Errorf("expected elements in trie are %d, but there are %d elements", 0, count)
	}
}
