## Introduction

intlist is a high-performance concurrent list.

## Features

- Concurrent safe API whit high-performance.
- Wait-free Contains and Range Operations.
- Sorted items.

## QuickStart

```go
package main

import (
	"fmt"
	"github.com/yangyunfeng007/linked_list"
)

func main() {
	l := linked_list.NewInt()
	for _, v := range []int{10, 12, 15} {
		if l.Insert(v) {
			fmt.Println("int list add", v)
		}
	}
	if l.Contains(10) {
		fmt.Println("int list contains 10")
	}
	l.Range(func(value int) bool {
		fmt.Println("int list find", value)
		return true
	})
	l.Delete(15)
	fmt.Printf("int list contains %d items\r\n", l.Len())
}
```

## Benchmark

```shell
$ go test -run=NOTEST -bench=. -benchtime=100000x -benchmem -count=10 -timeout=60m  > x.txt
$ benchstat x.txt
```

```shell
name                                            time/op
Insert/intlist-12                               27.1µs ± 5%
Insert/simpleintlist-12                          271µs ± 3%
Contains100Hits/intlist-12                      48.3ns ± 4%
Contains100Hits/simpleintlist-12                61.2ns ± 3%
Contains50Hits/intlist-12                       49.6ns ± 5%
Contains50Hits/simpleintlist-12                 63.6ns ± 9%
ContainsNoHits/intlist-12                       53.3ns ± 9%
ContainsNoHits/simpleintlist-12                 71.1ns ±11%
50Add50Contains/intlist-12                      12.8µs ± 6%
50Add50Contains/simpleintlist-12                99.4µs ± 2%
30Add70Contains/intlist-12                      6.67µs ± 3%
30Add70Contains/simpleintlist-12                37.2µs ± 1%
1Remove9Add90Contains/intlist-12                1.45µs ± 5%
1Remove9Add90Contains/simpleintlist-12          4.39µs ± 4%
1Range9Remove90Add900Contains/intlist-12        1.43µs ± 6%
1Range9Remove90Add900Contains/simpleintlist-12  4.65µs ± 5%
```