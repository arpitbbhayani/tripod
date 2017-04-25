Tripod
---
## What is Tripod?
Tripod provides search-by-prefix utilities.

## Highlights
PrefixStoreByteTrie:
 - 0 bytes of garbage allocation for `Put` and `Exists`
 - minimal and predictable allocations `(3 * (numOfElementsRetrieved) + 2)` for
   `PrefixSearch`

## Prefix Stores
PrefixStore is where you will store your prefix. You will call `Put`, `Exists`
and `PrefixSearch` on its instance. There are several type of implementations
available for this. Each implementation of Prefix Store has some tailor-made
optimization on method calls for type of data it can store. Following are the
types available

### PrefixStoreByteTrie:
This PrefixStore is implemented via an in-memory trie capable of storing only
`[]byte`.

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
    method `Put` where each key is of size 8 bytes.
 - `BenchmarkPrefixSearch8_10` means Benchmarking on a PrefixStoreByteTrie on
    method `PrefixSearch` where 10 elements are retrieved each of size 8 bytes.

### Results
```
BenchmarkByteTriePut8-4                         20000000               101 ns/op               0 B/op          0 allocs/op
BenchmarkByteTriePut16-4                        10000000               205 ns/op               0 B/op          0 allocs/op
BenchmarkByteTriePut32-4                         3000000               395 ns/op               0 B/op          0 allocs/op
BenchmarkByteTriePut64-4                         2000000               698 ns/op               0 B/op          0 allocs/op
BenchmarkByteTriePut128-4                        1000000              1660 ns/op               0 B/op          0 allocs/op
BenchmarkByteTrieExists8-4                      20000000                98.5 ns/op             0 B/op          0 allocs/op
BenchmarkByteTrieExists16-4                     10000000               192 ns/op               0 B/op          0 allocs/op
BenchmarkByteTrieExists32-4                      3000000               454 ns/op               0 B/op          0 allocs/op
BenchmarkByteTrieExists64-4                      2000000               846 ns/op               0 B/op          0 allocs/op
BenchmarkByteTrieExists128-4                     1000000              1910 ns/op               0 B/op          0 allocs/op
BenchmarkByteTriePrefixSearch8_10-4               200000             11880 ns/op            1056 B/op         32 allocs/op
BenchmarkByteTriePrefixSearch16_10-4              100000             22321 ns/op            1136 B/op         32 allocs/op
BenchmarkByteTriePrefixSearch32_10-4               50000             38142 ns/op            1296 B/op         32 allocs/op
BenchmarkByteTriePrefixSearch64_10-4               20000             79904 ns/op            1616 B/op         32 allocs/op
BenchmarkByteTriePrefixSearch128_10-4              10000            175410 ns/op            2256 B/op         32 allocs/op
BenchmarkByteTriePrefixSearch8_50-4                30000             55869 ns/op            4576 B/op        152 allocs/op
BenchmarkByteTriePrefixSearch16_50-4               20000            107306 ns/op            4976 B/op        152 allocs/op
BenchmarkByteTriePrefixSearch32_50-4               10000            224612 ns/op            5776 B/op        152 allocs/op
BenchmarkByteTriePrefixSearch64_50-4                3000            445358 ns/op            7376 B/op        152 allocs/op
BenchmarkByteTriePrefixSearch128_50-4               2000            926553 ns/op           10576 B/op        152 allocs/op
BenchmarkByteTriePrefixSearch8_100-4               20000            102505 ns/op            8976 B/op        302 allocs/op
BenchmarkByteTriePrefixSearch16_100-4              10000            217612 ns/op            9776 B/op        302 allocs/op
BenchmarkByteTriePrefixSearch32_100-4               3000            449359 ns/op           11376 B/op        302 allocs/op
BenchmarkByteTriePrefixSearch64_100-4               2000            918052 ns/op           14576 B/op        302 allocs/op
BenchmarkByteTriePrefixSearch128_100-4              1000           1665095 ns/op           20976 B/op        302 allocs/op
BenchmarkByteTriePrefixSearch8_200-4               10000            211612 ns/op           17776 B/op        602 allocs/op
BenchmarkByteTriePrefixSearch16_200-4               3000            425024 ns/op           19376 B/op        602 allocs/op
BenchmarkByteTriePrefixSearch32_200-4               2000            921052 ns/op           22576 B/op        602 allocs/op
BenchmarkByteTriePrefixSearch64_200-4               1000           1942111 ns/op           28976 B/op        602 allocs/op
BenchmarkByteTriePrefixSearch128_200-4               500           3530201 ns/op           41776 B/op        602 allocs/op
```

## Contribution
In case you loved this utility and have a great feature idea, then feel free to
contribute . The complete utility is written in Go. So for contributing all you
need to have is working knowledge of Go.

## Issues
Please report any glitch, bug, error or an unhandled exception. Feel free
to [create one](https://github.com/arpitbbhayani/tripod/issues/new).
