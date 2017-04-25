// Tripod provides search-by-prefix utilities.
package tripod

import (
	"container/list"
	"fmt"
)

// Represents the PrefixStore which uses an in-memory trie data-structure to
// store keys and efficiently return them when searched by prefix.
// PrefixStoreByteTrie is optimized for type []byte.
type PrefixStoreByteTrie struct {
	isLast            bool
	children          map[int]*PrefixStoreByteTrie
	maxKeySizeInBytes int
}

// Creates and returns reference to a new instance of PrefixStoreByteTrie.
// maxKeySizeInBytes is the maximum size of the key ([]byte) that should be
// allowed to be added to the PrefixStore. When tried to put key of length more
// than maxKeySizeInBytes, the method will return the error.
func CreatePrefixStoreByteTrie(maxKeySizeInBytes int) *PrefixStoreByteTrie {
	return &PrefixStoreByteTrie{
		children:          make(map[int]*PrefixStoreByteTrie),
		maxKeySizeInBytes: maxKeySizeInBytes,
	}
}

// Adds the []byte key to the PrefixStore and returns if key was succesfully
// added and any error encountered.
// A non nil error is returned if len(key) > maxKeySizeInBytes
func (t *PrefixStoreByteTrie) Put(key []byte) (bool, error) {
	if len(key) > t.maxKeySizeInBytes {
		return false, fmt.Errorf("max size of key should be %d (%d > %d)",
			t.maxKeySizeInBytes, len(key), t.maxKeySizeInBytes)
	}

	current_node := t
	for _, b := range key {
		child := current_node.children[int(b)]
		if child == nil {
			child = CreatePrefixStoreByteTrie(t.maxKeySizeInBytes)
			current_node.children[int(b)] = child
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
func (t *PrefixStoreByteTrie) Exists(key []byte) bool {
	if len(key) > t.maxKeySizeInBytes {
		// Shorting the lookup, since Put method does not allow to put key of
		// size > maxKeySizeInBytes, hence shorting this evaluation.
		return false
	}

	current_node := t
	for _, b := range key {
		child := current_node.children[int(b)]
		if child == nil {
			return false
		}
		current_node = child
	}
	return current_node.isLast
}

// For a given instance of PrefixStore t, this method returns a reference to
// subPrefixStore that ends at the key.
func (t *PrefixStoreByteTrie) get(key []byte) *PrefixStoreByteTrie {
	if len(key) > t.maxKeySizeInBytes {
		// Shorting the lookup, since Put method does not allow to put key of
		// size > maxKeySizeInBytes, hence shorting this evaluation.
		return nil
	}

	current_node := t
	for _, b := range key {
		child := current_node.children[int(b)]
		if child == nil {
			return nil
		}
		current_node = child
	}
	return current_node
}

// Does a Depth First Search traversal on the PrefixStoreByteTrie and returns a
// List (Double Linked List) containing the keys present in the
// PrefixStore for the given prefix.
func (t *PrefixStoreByteTrie) list(prefix []byte) *list.List {
	entries := list.New()
	buffer := make([]byte, 0, t.maxKeySizeInBytes)

	for ch, tt := range t.children {
		buffer = append(buffer, byte(ch))
		_dfs(tt, prefix, buffer, entries)
		buffer = buffer[:len(buffer)-1]
	}
	return entries
}

// The DFS Function which recursively calls the children and as it encounters
// a valid existing key, creates a copy of the temporary buffer
// and appends prefix + copyOfBuffer to the linked list; thus forming the
// complete key
func _dfs(t *PrefixStoreByteTrie, prefix []byte, buffer []byte, entries *list.List) {
	if t.isLast {
		copyOfBuffer := make([]byte, 0, len(buffer)+len(prefix))

		// TODO: Check if any optimizations can be done here
		copyOfBuffer = append(copyOfBuffer, prefix...)
		copyOfBuffer = append(copyOfBuffer, buffer...)

		entries.PushBack(copyOfBuffer)
	}
	for ch, tt := range t.children {
		buffer = append(buffer, byte(ch))
		_dfs(tt, prefix, buffer, entries)
		buffer = buffer[:len(buffer)-1]
	}
}

// Does the prefix search on the PrefixStore and returns a reference to
// list (*list.List) containings all entries from the store for the given
// prefix. Each element of the list is []byte.
func (t *PrefixStoreByteTrie) PrefixSearch(prefix []byte) *list.List {
	if len(prefix) > t.maxKeySizeInBytes {
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
