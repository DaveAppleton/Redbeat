package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var a, b, c, d uint64
var mu sync.Mutex

func BenchmarkAtomic(bch *testing.B) {
	bch.ResetTimer()
	for j := 0; j < bch.N; j++ {
		e := atomic.LoadUint64(&a)
		f := atomic.LoadUint64(&b)
		g := atomic.LoadUint64(&c)
		h := atomic.LoadUint64(&d)
		atomic.StoreUint64(&a, h)
		atomic.StoreUint64(&b, g)
		atomic.StoreUint64(&c, f)
		atomic.StoreUint64(&d, e)
	}
}

func BenchmarkI(b *testing.B) {
	myData := []string{
		"one hundred days at sea",
		"one hundred days at sea",
		"one hundred days at sea",
		"one hundred days at sea",
		"one hundred days at sea",
	}
	count := 0
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for row := 0; row < len(myData); row++ {
			if myData[row] == "one hundred days at sea" {
				count++
			}
		}
	}
	fmt.Println(count)
}

func BenchmarkV(b *testing.B) {
	myData := []string{
		"one hundred days at sea",
		"one hundred days at sea",
		"one hundred days at sea",
		"one hundred days at sea",
		"one hundred days at sea",
	}
	count := 0
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for _, line := range myData {
			if line == "one hundred days at sea" {
				count++
			}
		}
	}
	fmt.Println(count)
}

// 0.000001 ns/op
// 0.000001 ns/op
func BenchmarkColRow(b *testing.B) {
	myData := []string{
		"one hundred days at sea",
		"one hundred days at sea",
		"one hundred days at sea",
		"one hundred days at sea",
		"one hundred days at sea",
	}
	letter := byte('a')
	count := 0
	for col := 0; col < 23; col++ {
		for row := 0; row < len(myData); row++ {
			if myData[row][col] == letter {
				count++
			}
		}
	}
}

//BenchmarkAtomic-8   	18676130	       109 ns/op	       0 B/op	       0 allocs/op
//PASS
//ok  	af.sorry.gasprice	3.009s

func BenchmarkMutex(bch *testing.B) {
	bch.ResetTimer()
	for j := 0; j < bch.N; j++ {
		mu.Lock()
		e := a
		f := b
		g := c
		h := d
		a = h
		b = g
		c = f
		d = e
		mu.Unlock()
	}
}
