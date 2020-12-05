package bitmap

import (
	"math/rand"
	"testing"
)

const (
	bitmapSize = 1024 * 1024
)

func BenchmarkSetOneMillion(b *testing.B) {
	bm := NewBitMap(bitmapSize)
	datas := rand.Perm(bitmapSize)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bitmapSize; j++ {
			err := bm.Set(int64(datas[j]))
			if err != nil {
				b.Fatal(err)
			}
		}
	}
	b.StopTimer()
}
func BenchmarkGetOneMillion(b *testing.B) {
	bm := NewBitMap(bitmapSize)
	datas := rand.Perm(bitmapSize)

	for i := 0; i < bitmapSize; i++ {
		err := bm.Set(int64(datas[i]))
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bitmapSize; j++ {
			isSed, err := bm.Get(int64(datas[j]))
			if err != nil {
				b.Fatal(err)
			}
			if isSed == false {
				b.Fatalf("with error, get %d should be sed", j)
			}
		}
	}
	b.StopTimer()
}
func BenchmarkResetSetNOneMillion(b *testing.B) {
	bm := NewBitMap(bitmapSize)
	datas := rand.Perm(bitmapSize)
	for j := 0; j < bitmapSize; j++ {
		err := bm.Set(int64(datas[j]))
		if err != nil {
			b.Fatal(err)
		}
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < bitmapSize; j++ {
			_, err := bm.ResetN(int64(datas[j]))
			if err != nil {
				b.Fatal(err)
			}
		}
	}
	b.StopTimer()
}
