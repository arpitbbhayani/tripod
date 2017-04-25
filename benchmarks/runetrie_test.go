package benchmark_tripod

import (
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

func benchmarkRunTriePut(b *testing.B, size int) {
	tr := tripod.CreatePrefixStoreRuneTrie(128)
	x := getRandomUTF8RuneSlice(size)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.Put(x)
	}
}


func BenchmarkRuneTriePut8(b *testing.B)   { benchmarkRunTriePut(b, 8) }
func BenchmarkRuneTriePut16(b *testing.B)  { benchmarkRunTriePut(b, 16) }
func BenchmarkRuneTriePut32(b *testing.B)  { benchmarkRunTriePut(b, 32) }
func BenchmarkRuneTriePut64(b *testing.B)  { benchmarkRunTriePut(b, 64) }
func BenchmarkRuneTriePut128(b *testing.B) { benchmarkRunTriePut(b, 128) }

func benchmarkRunTrieExists(b *testing.B, size int) {
	tr := tripod.CreatePrefixStoreRuneTrie(128)
	x := getRandomUTF8RuneSlice(size)
	tr.Put(x)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.Exists(x)
	}
}

func BenchmarkRuneTrieExists8(b *testing.B)   { benchmarkRunTrieExists(b, 8) }
func BenchmarkRuneTrieExists16(b *testing.B)  { benchmarkRunTrieExists(b, 16) }
func BenchmarkRuneTrieExists32(b *testing.B)  { benchmarkRunTrieExists(b, 32) }
func BenchmarkRuneTrieExists64(b *testing.B)  { benchmarkRunTrieExists(b, 64) }
func BenchmarkRuneTrieExists128(b *testing.B) { benchmarkRunTrieExists(b, 128) }

func populatePrefixStoreRuneTrie(tr *tripod.PrefixStoreRuneTrie, size int, count int) {
	for i := 0; i < count; i++ {
		tr.Put(getRandomUTF8RuneSlice(size))
	}
}

func populatePrefixStoreRuneTrieForPrefix(tr *tripod.PrefixStoreRuneTrie, size int, count int, b rune) {
	for i := 0; i < count; i++ {
		x := getRandomUTF8RuneSlice(size)
		x[0] = b
		tr.Put(x)
	}
}

func benchmarkRuneTriePrefixSearch(b *testing.B, size int, count int) {
	tr := tripod.CreatePrefixStoreRuneTrie(128)
	x := []rune("a")
	populatePrefixStoreRuneTrieForPrefix(tr, size, count, 'a')
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		tr.PrefixSearch(x)
	}
}

func BenchmarkRuneTriePrefixSearch8_10(b *testing.B)   { benchmarkRuneTriePrefixSearch(b, 8, 10) }
func BenchmarkRuneTriePrefixSearch16_10(b *testing.B)  { benchmarkRuneTriePrefixSearch(b, 16, 10) }
func BenchmarkRuneTriePrefixSearch32_10(b *testing.B)  { benchmarkRuneTriePrefixSearch(b, 32, 10) }
func BenchmarkRuneTriePrefixSearch64_10(b *testing.B)  { benchmarkRuneTriePrefixSearch(b, 64, 10) }
func BenchmarkRuneTriePrefixSearch128_10(b *testing.B) { benchmarkRuneTriePrefixSearch(b, 128, 10) }

func BenchmarkRuneTriePrefixSearch8_50(b *testing.B)   { benchmarkRuneTriePrefixSearch(b, 8, 50) }
func BenchmarkRuneTriePrefixSearch16_50(b *testing.B)  { benchmarkRuneTriePrefixSearch(b, 16, 50) }
func BenchmarkRuneTriePrefixSearch32_50(b *testing.B)  { benchmarkRuneTriePrefixSearch(b, 32, 50) }
func BenchmarkRuneTriePrefixSearch64_50(b *testing.B)  { benchmarkRuneTriePrefixSearch(b, 64, 50) }
func BenchmarkRuneTriePrefixSearch128_50(b *testing.B) { benchmarkRuneTriePrefixSearch(b, 128, 50) }

func BenchmarkRuneTriePrefixSearch8_100(b *testing.B)   { benchmarkRuneTriePrefixSearch(b, 8, 100) }
func BenchmarkRuneTriePrefixSearch16_100(b *testing.B)  { benchmarkRuneTriePrefixSearch(b, 16, 100) }
func BenchmarkRuneTriePrefixSearch32_100(b *testing.B)  { benchmarkRuneTriePrefixSearch(b, 32, 100) }
func BenchmarkRuneTriePrefixSearch64_100(b *testing.B)  { benchmarkRuneTriePrefixSearch(b, 64, 100) }
func BenchmarkRuneTriePrefixSearch128_100(b *testing.B) { benchmarkRuneTriePrefixSearch(b, 128, 100) }

func BenchmarkRuneTriePrefixSearch8_200(b *testing.B)   { benchmarkRuneTriePrefixSearch(b, 8, 200) }
func BenchmarkRuneTriePrefixSearch16_200(b *testing.B)  { benchmarkRuneTriePrefixSearch(b, 16, 200) }
func BenchmarkRuneTriePrefixSearch32_200(b *testing.B)  { benchmarkRuneTriePrefixSearch(b, 32, 200) }
func BenchmarkRuneTriePrefixSearch64_200(b *testing.B)  { benchmarkRuneTriePrefixSearch(b, 64, 200) }
func BenchmarkRuneTriePrefixSearch128_200(b *testing.B) { benchmarkRuneTriePrefixSearch(b, 128, 200) }
