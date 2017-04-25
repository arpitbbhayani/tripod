// Tripod provides search-by-prefix utilities.
package tripod

import (
	"container/list"
	"fmt"
)

// Represents the PrefixStore which uses an in-memory trie data-structure to
// store keys and efficiently return them when searched by prefix.
// PrefixStoreRuneTrie is optimized for type []rune.
type PrefixStoreRuneTrie struct {
	isLast            bool
	children          map[rune]*PrefixStoreRuneTrie
	maxKeySizeInRunes int
}

// Creates and returns reference to a new instance of PrefixStoreRuneTrie.
// maxKeySizeInRunes is the maximum size of the key ([]rune) that should be
// allowed to be added to the PrefixStore. When tried to put key of length more
// than maxKeySizeInRunes, the method will return the error.
func CreatePrefixStoreRuneTrie(maxKeySizeInRunes int) *PrefixStoreRuneTrie {
	return &PrefixStoreRuneTrie{
		children:          make(map[rune]*PrefixStoreRuneTrie),
		maxKeySizeInRunes: maxKeySizeInRunes,
	}
}

// Adds the []rune key to the PrefixStore and returns if key was succesfully
// added and any error encountered.
// A non nil error is returned if len(key) > maxKeySizeInRunes
func (t *PrefixStoreRuneTrie) Put(key []rune) (bool, error) {
	if len(key) > t.maxKeySizeInRunes {
		return false, fmt.Errorf("max size of key should be %d (%d > %d)",
			t.maxKeySizeInRunes, len(key), t.maxKeySizeInRunes)
	}

	current_node := t
	for _, b := range key {
		child := current_node.children[b]
		if child == nil {
			child = CreatePrefixStoreRuneTrie(t.maxKeySizeInRunes)
			current_node.children[b] = child
		}
		current_node = child
	}

	// If key is empty then current_node == t, and since nothing was added
	// hence returning false.
	if current_node == t {
		return false, nil
	}

	newlyAdded := current_node.isLast == false
	current_node.isLast = true
	return newlyAdded, nil
}

// Checks and returns if given key is present in the PrefixStore
func (t *PrefixStoreRuneTrie) Exists(key []rune) bool {
	if len(key) > t.maxKeySizeInRunes {
		// Shorting the lookup, since Put method does not allow to put key of
		// size > maxKeySizeInRunes, hence shorting this evaluation.
		return false
	}

	current_node := t
	for _, b := range key {
		child := current_node.children[b]
		if child == nil {
			return false
		}
		current_node = child
	}
	return current_node.isLast
}

// For a given instance of PrefixStore t, this method returns a reference to
// subPrefixStore that ends at the key.
func (t *PrefixStoreRuneTrie) get(key []rune) *PrefixStoreRuneTrie {
	if len(key) > t.maxKeySizeInRunes {
		// Shorting the lookup, since Put method does not allow to put key of
		// size > maxKeySizeInRunes, hence shorting this evaluation.
		return nil
	}

	current_node := t
	for _, b := range key {
		child := current_node.children[b]
		if child == nil {
			return nil
		}
		current_node = child
	}
	return current_node
}

// Does a Depth First Search traversal on the PrefixStoreRuneTrie and returns a
// List (Double Linked List) containing the keys present in the
// PrefixStore for the given prefix.
func (t *PrefixStoreRuneTrie) list(prefix []rune) *list.List {
	entries := list.New()
	buffer := make([]rune, 0, t.maxKeySizeInRunes)

	for ch, tt := range t.children {
		buffer = append(buffer, ch)
		_dfs_rune(tt, prefix, buffer, entries)
		buffer = buffer[:len(buffer)-1]
	}
	return entries
}

// The DFS Function which recursively calls the children and as it encounters
// a valid existing key, creates a copy of the temporary buffer
// and appends prefix + copyOfBuffer to the linked list; thus forming the
// complete key
func _dfs_rune(t *PrefixStoreRuneTrie, prefix []rune, buffer []rune, entries *list.List) {
	if t.isLast {
		copyOfBuffer := make([]rune, 0, len(buffer) + len(prefix))

		// TODO: Check if any optimizations can be done here
		copyOfBuffer = append(copyOfBuffer, prefix...)
		copyOfBuffer = append(copyOfBuffer, buffer...)

		entries.PushBack(copyOfBuffer)
	}
	for ch, tt := range t.children {
		buffer = append(buffer, ch)
		_dfs_rune(tt, prefix, buffer, entries)
		buffer = buffer[:len(buffer)-1]
	}
}

// Does the prefix search on the PrefixStore and returns a reference to
// list (*list.List) containings all entries from the store for the given
// prefix. Each element of the list is []rune.
func (t *PrefixStoreRuneTrie) PrefixSearch(prefix []rune) *list.List {
	if len(prefix) > t.maxKeySizeInRunes {
		return list.New()
	}
	subTrie := t.get(prefix)
	if subTrie == nil {
		return list.New()
	}

	entries := subTrie.list(prefix)
	if subTrie.isLast {
		entries.PushFront(prefix)
	}
	return entries
}
