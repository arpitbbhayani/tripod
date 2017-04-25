# Tripod
Tripod provides fast and optimal search-by-prefix utilities in Go.

## Highlights
PrefixStoreByteTrie
 - 0 bytes of garbage allocation for `Put` and `Exists`
 - minimal and predictable allocations `(3 * (numOfElementsRetrieved) + 2)` for
   `PrefixSearch`

PrefixStoreRuneTrie
 - 0 bytes of garbage allocation for `Put` and `Exists`
 - minimal and predictable allocations `(3 * (numOfElementsRetrieved) + 2)` for
   `PrefixSearch`
 - can safely store UTF-8 characters

## Prefix Stores
PrefixStore is where you will store your prefix. You will call `Put`, `Exists`
and `PrefixSearch` on its instance. There are several type of implementations
available for this. Each implementation of Prefix Store has some tailor-made
optimization on method calls for type of data it can store. Following are the
types available

### PrefixStoreByteTrie
This PrefixStore is implemented via an in-memory trie capable of storing only
`[]byte`. This should be used when you know that the data you would put in only
has ASCII characters. No one is stopping you from putting any kind of `[]byte`.

### PrefixStoreRuneTrie
This PrefixStore is implemented via an in-memory trie capable of storing only
`[]rune`. This is useful when you want to store data that might have
UTF-8 characters.

## Installation
```
go get github.com/arpitbbhayani/tripod
```

## Quickstart: Hello, World!
```go
package main

import (
	"fmt"
	"github.com/arpitbbhayani/tripod"
)

func main() {
	tr := tripod.CreatePrefixStoreByteTrie(8)

	tr.Put([]byte("go"))
	tr.Put([]byte("is"))
	tr.Put([]byte("good"))

	results := tr.PrefixSearch([]byte("g"))
	for e := results.Front(); e != nil; e = e.Next() {
		fmt.Println(string(e.Value.([]byte)))
	}
}
```

## Documentation
http://godoc.org/github.com/arpitbbhayani/tripod

## Running Tests and Benchmarks

### Running Tests
```bash
cd tests
go test
```

### Running Benchmarks
```bash
cd benchmarks
go test -bench=. -benchmem
```

## Benchmarks

### Conventions
 - `BenchmarkByteTriePut8` means Benchmarking on a PrefixStoreByteTrie on
    method `Put` where each key is of size 8 bytes/runes.
 - `BenchmarkByteTriePrefixSearch8_10` means Benchmarking on a
    PrefixStoreByteTrie on method `PrefixSearch` where 10 elements are
    retrieved each of size 8 bytes/runes.

### Results
#### PrefixStoreByteTrie
```
BenchmarkByteTriePut32-4                         3000000               395 ns/op               0 B/op          0 allocs/op
BenchmarkByteTriePut128-4                        1000000              1660 ns/op               0 B/op          0 allocs/op
BenchmarkByteTrieExists32-4                      3000000               454 ns/op               0 B/op          0 allocs/op
BenchmarkByteTriePrefixSearch32_50-4               10000            224612 ns/op            5776 B/op        152 allocs/op
BenchmarkByteTriePrefixSearch128_50-4               2000            926553 ns/op           10576 B/op        152 allocs/op
BenchmarkByteTriePrefixSearch32_200-4               2000            921052 ns/op           22576 B/op        602 allocs/op
BenchmarkByteTriePrefixSearch128_200-4               500           3530201 ns/op           41776 B/op        602 allocs/op
```

#### PrefixStoreRuneTrie
```
BenchmarkRuneTriePut32-4                         3000000               473 ns/op               0 B/op          0 allocs/op
BenchmarkRuneTriePut128-4                        1000000              2038 ns/op               0 B/op          0 allocs/op
BenchmarkRuneTrieExists32-4                      3000000               454 ns/op               0 B/op          0 allocs/op
BenchmarkRuneTrieExists128-4                     1000000              1922 ns/op               0 B/op          0 allocs/op
BenchmarkRuneTriePrefixSearch32_50-4               10000            225612 ns/op           10960 B/op        152 allocs/op
BenchmarkRuneTriePrefixSearch128_50-4               2000            885550 ns/op           30160 B/op        152 allocs/op
BenchmarkRuneTriePrefixSearch32_200-4               2000            742042 ns/op           42160 B/op        602 allocs/op
BenchmarkRuneTriePrefixSearch128_200-4               500           3248185 ns/op          118960 B/op        602 allocs/op
```

## Contribution
In case you loved this utility and have a great feature idea, then feel free to
contribute . The complete utility is written in Go. So for contributing all you
need to have is working knowledge of Go.

## Issues
Please report any glitch, bug, error or an unhandled exception. Feel free
to [create one](https://github.com/arpitbbhayani/tripod/issues/new).
