package linked_list

import (
	"math"
	"testing"
)

const initsize = 1 << 10 // for `contains` `1Remove9Add90Contains` `1Range9Remove90Add900Contains`
const randN = math.MaxUint32

func BenchmarkInsert(b *testing.B) {
	b.Run("intlist", func(b *testing.B) {
		l := NewInt()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				l.Insert(int(fastrandn(randN)))
			}
		})
	})
	b.Run("simpleintlist", func(b *testing.B) {
		l := NewSimpleInt()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				l.Insert(int(fastrandn(randN)))
			}
		})
	})
}

func BenchmarkContains100Hits(b *testing.B) {
	b.Run("intlist", func(b *testing.B) {
		l := NewInt()
		for i := 0; i < initsize; i++ {
			l.Insert(i)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = l.Contains(int(fastrandn(initsize)))
			}
		})
	})
	b.Run("simpleintlist", func(b *testing.B) {
		l := NewSimpleInt()
		for i := 0; i < initsize; i++ {
			l.Insert(i)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = l.Contains(int(fastrandn(initsize)))
			}
		})
	})
}

func BenchmarkContains50Hits(b *testing.B) {
	const rate = 2
	b.Run("intlist", func(b *testing.B) {
		l := NewInt()
		for i := 0; i < initsize*rate; i++ {
			if fastrandn(rate) == 0 {
				l.Insert(i)
			}
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = l.Contains(int(fastrandn(initsize * rate)))
			}
		})
	})
	b.Run("simpleintlist", func(b *testing.B) {
		l := NewSimpleInt()
		for i := 0; i < initsize*rate; i++ {
			if fastrandn(rate) == 0 {
				l.Insert(i)
			}
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = l.Contains(int(fastrandn(initsize * rate)))
			}
		})
	})
}

func BenchmarkContainsNoHits(b *testing.B) {
	b.Run("intlist", func(b *testing.B) {
		l := NewInt()
		invalid := make([]int, 0, initsize)
		for i := 0; i < initsize*2; i++ {
			if i%2 == 0 {
				l.Insert(i)
			} else {
				invalid = append(invalid, i)
			}
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = l.Contains(invalid[fastrandn(uint32(len(invalid)))])
			}
		})
	})
	b.Run("simpleintlist", func(b *testing.B) {
		l := NewSimpleInt()
		invalid := make([]int, 0, initsize)
		for i := 0; i < initsize*2; i++ {
			if i%2 == 0 {
				l.Insert(i)
			} else {
				invalid = append(invalid, i)
			}
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = l.Contains(invalid[fastrandn(uint32(len(invalid)))])
			}
		})
	})
}

func Benchmark50Add50Contains(b *testing.B) {
	b.Run("intlist", func(b *testing.B) {
		l := NewInt()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				u := fastrandn(10)
				if u < 5 {
					l.Insert(int(fastrandn(randN)))
				} else {
					l.Contains(int(fastrandn(randN)))
				}
			}
		})
	})
	b.Run("simpleintlist", func(b *testing.B) {
		l := NewSimpleInt()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				u := fastrandn(10)
				if u < 5 {
					l.Insert(int(fastrandn(randN)))
				} else {
					l.Contains(int(fastrandn(randN)))
				}
			}
		})
	})
}

func Benchmark30Add70Contains(b *testing.B) {
	b.Run("intlist", func(b *testing.B) {
		l := NewInt()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				u := fastrandn(10)
				if u < 3 {
					l.Insert(int(fastrandn(randN)))
				} else {
					l.Contains(int(fastrandn(randN)))
				}
			}
		})
	})
	b.Run("simpleintlist", func(b *testing.B) {
		l := NewSimpleInt()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				u := fastrandn(10)
				if u < 3 {
					l.Insert(int(fastrandn(randN)))
				} else {
					l.Contains(int(fastrandn(randN)))
				}
			}
		})
	})
}

func Benchmark1Remove9Add90Contains(b *testing.B) {
	b.Run("intlist", func(b *testing.B) {
		l := NewInt()
		for i := 0; i < initsize; i++ {
			l.Insert(i)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				u := fastrandn(100)
				if u < 9 {
					l.Insert(int(fastrandn(randN)))
				} else if u == 10 {
					l.Delete(int(fastrandn(randN)))
				} else {
					l.Contains(int(fastrandn(randN)))
				}
			}
		})
	})
	b.Run("simpleintlist", func(b *testing.B) {
		l := NewSimpleInt()
		for i := 0; i < initsize; i++ {
			l.Insert(i)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				u := fastrandn(100)
				if u < 9 {
					l.Insert(int(fastrandn(randN)))
				} else if u == 10 {
					l.Delete(int(fastrandn(randN)))
				} else {
					l.Contains(int(fastrandn(randN)))
				}
			}
		})
	})
}

func Benchmark1Range9Remove90Add900Contains(b *testing.B) {
	b.Run("intlist", func(b *testing.B) {
		l := NewInt()
		for i := 0; i < initsize; i++ {
			l.Insert(int(i))
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				u := fastrandn(1000)
				if u == 0 {
					l.Range(func(score int) bool {
						return true
					})
				} else if u > 10 && u < 20 {
					l.Delete(int(fastrandn(randN)))
				} else if u >= 100 && u < 190 {
					l.Insert(int(fastrandn(randN)))
				} else {
					l.Contains(int(fastrandn(randN)))
				}
			}
		})
	})
	b.Run("simpleintlist", func(b *testing.B) {
		l := NewSimpleInt()
		for i := 0; i < initsize; i++ {
			l.Insert(i)
		}
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				u := fastrandn(1000)
				if u == 0 {
					l.Range(func(score int) bool {
						return true
					})
				} else if u > 10 && u < 20 {
					l.Delete(int(fastrandn(randN)))
				} else if u >= 100 && u < 190 {
					l.Insert(int(fastrandn(randN)))
				} else {
					l.Contains(int(fastrandn(randN)))
				}
			}
		})
	})
}
