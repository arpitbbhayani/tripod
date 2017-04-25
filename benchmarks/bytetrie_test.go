package benchmark_tripod

import (
	"github.com/arpitbbhayani/tripod"
	"math/rand"
	"testing"
)

func getRandomByteSlice(size int) []byte {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, size)
	for i := 0; i < size; i++ {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func benchmarkPut(b *testing.B, size int) {
	tr := tripod.CreatePrefixStoreByteTrie(128)
	x := getRandomByteSlice(size)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.Put(x)
	}
}

func BenchmarkByteTriePut8(b *testing.B)   { benchmarkPut(b, 8) }
func BenchmarkByteTriePut16(b *testing.B)  { benchmarkPut(b, 16) }
func BenchmarkByteTriePut32(b *testing.B)  { benchmarkPut(b, 32) }
func BenchmarkByteTriePut64(b *testing.B)  { benchmarkPut(b, 64) }
func BenchmarkByteTriePut128(b *testing.B) { benchmarkPut(b, 128) }

func benchmarkExists(b *testing.B, size int) {
	tr := tripod.CreatePrefixStoreByteTrie(128)
	x := getRandomByteSlice(size)
	tr.Put(x)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.Exists(x)
	}
}

func BenchmarkByteTrieExists8(b *testing.B)   { benchmarkExists(b, 8) }
func BenchmarkByteTrieExists16(b *testing.B)  { benchmarkExists(b, 16) }
func BenchmarkByteTrieExists32(b *testing.B)  { benchmarkExists(b, 32) }
func BenchmarkByteTrieExists64(b *testing.B)  { benchmarkExists(b, 64) }
func BenchmarkByteTrieExists128(b *testing.B) { benchmarkExists(b, 128) }

func populatePrefixStoreByteTrie(tr *tripod.PrefixStoreByteTrie, size int, count int) {
	for i := 0; i < count; i++ {
		tr.Put(getRandomByteSlice(size))
	}
}

func populatePrefixStoreByteTrieForPrefix(tr *tripod.PrefixStoreByteTrie, size int, count int, b byte) {
	for i := 0; i < count; i++ {
		x := getRandomByteSlice(size)
		x[0] = b
		tr.Put(x)
	}
}

func benchmarkPrefixSearch(b *testing.B, size int, count int) {
	tr := tripod.CreatePrefixStoreByteTrie(128)
	x := []byte("a")
	populatePrefixStoreByteTrieForPrefix(tr, size, count, 'a')
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.PrefixSearch(x)
	}
}

func BenchmarkByteTriePrefixSearch8_10(b *testing.B)   { benchmarkPrefixSearch(b, 8, 10) }
func BenchmarkByteTriePrefixSearch16_10(b *testing.B)  { benchmarkPrefixSearch(b, 16, 10) }
func BenchmarkByteTriePrefixSearch32_10(b *testing.B)  { benchmarkPrefixSearch(b, 32, 10) }
func BenchmarkByteTriePrefixSearch64_10(b *testing.B)  { benchmarkPrefixSearch(b, 64, 10) }
func BenchmarkByteTriePrefixSearch128_10(b *testing.B) { benchmarkPrefixSearch(b, 128, 10) }

func BenchmarkByteTriePrefixSearch8_50(b *testing.B)   { benchmarkPrefixSearch(b, 8, 50) }
func BenchmarkByteTriePrefixSearch16_50(b *testing.B)  { benchmarkPrefixSearch(b, 16, 50) }
func BenchmarkByteTriePrefixSearch32_50(b *testing.B)  { benchmarkPrefixSearch(b, 32, 50) }
func BenchmarkByteTriePrefixSearch64_50(b *testing.B)  { benchmarkPrefixSearch(b, 64, 50) }
func BenchmarkByteTriePrefixSearch128_50(b *testing.B) { benchmarkPrefixSearch(b, 128, 50) }

func BenchmarkByteTriePrefixSearch8_100(b *testing.B)   { benchmarkPrefixSearch(b, 8, 100) }
func BenchmarkByteTriePrefixSearch16_100(b *testing.B)  { benchmarkPrefixSearch(b, 16, 100) }
func BenchmarkByteTriePrefixSearch32_100(b *testing.B)  { benchmarkPrefixSearch(b, 32, 100) }
func BenchmarkByteTriePrefixSearch64_100(b *testing.B)  { benchmarkPrefixSearch(b, 64, 100) }
func BenchmarkByteTriePrefixSearch128_100(b *testing.B) { benchmarkPrefixSearch(b, 128, 100) }

func BenchmarkByteTriePrefixSearch8_200(b *testing.B)   { benchmarkPrefixSearch(b, 8, 200) }
func BenchmarkByteTriePrefixSearch16_200(b *testing.B)  { benchmarkPrefixSearch(b, 16, 200) }
func BenchmarkByteTriePrefixSearch32_200(b *testing.B)  { benchmarkPrefixSearch(b, 32, 200) }
func BenchmarkByteTriePrefixSearch64_200(b *testing.B)  { benchmarkPrefixSearch(b, 64, 200) }
func BenchmarkByteTriePrefixSearch128_200(b *testing.B) { benchmarkPrefixSearch(b, 128, 200) }
