/*
Credits: https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go#comment56305515_1760757

AMD Athlon64 X2
2GB DDR2

Go 1.5.1 linux/amd64

Ubuntu 15.10
Linux asus 4.2.0-19-generic #23-Ubuntu SMP Wed Nov 11 11:39:30 UTC 2015 x86_64 x86_64 x86_64 GNU/Linux

BenchmarkStringSprintf-2	     100	  10140559 ns/op	 5518024 B/op	    4616 allocs/op
BenchmarkStringJoin-2   	     200	   6962496 ns/op	 5388737 B/op	    2000 allocs/op
BenchmarkStringAdd-2    	     500	   3440360 ns/op	 2694352 B/op	     999 allocs/op
BenchmarkStringWrite-2  	   30000	     54597 ns/op	   18656 B/op	       8 allocs/op
BenchmarkStringAppend-2 	   30000	     41325 ns/op	   20736 B/op	      13 allocs/op
BenchmarkBytesJoin-2    	     300	   4883126 ns/op	 2694449 B/op	    1001 allocs/op
BenchmarkBytesAppend-2  	   30000	     53491 ns/op	   20736 B/op	      13 allocs/op
BenchmarkBytesWrite-2   	   20000	     64233 ns/op	   18656 B/op	       8 allocs/op

Feel free to change the LIMIT.
*/
package test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var (
	s1 = "local"
	b1 = []byte("local")
)

const LIMIT = 1000

func BenchmarkStringSprintf(b *testing.B) {
	var q string
	for i := 0; i < b.N; i++ {
		for j := 0; j < LIMIT; j++ {
			q = fmt.Sprintf("%s%s", q, s1)
		}
		q = ""
	}
	b.ReportAllocs()
}

func BenchmarkStringJoin(b *testing.B) {
	var q string
	for i := 0; i < b.N; i++ {
		for j := 0; j < LIMIT; j++ {
			q = strings.Join([]string{q, s1}, "")
		}
		q = ""
	}
	b.ReportAllocs()
}

func BenchmarkStringAdd(b *testing.B) {
	var q string
	for i := 0; i < b.N; i++ {
		for j := 0; j < LIMIT; j++ {
			q = q + s1
		}
		q = ""
	}
	b.ReportAllocs()
}

func BenchmarkStringWrite(b *testing.B) {
	q := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		for j := 0; j < LIMIT; j++ {
			q.WriteString(s1)
		}
		q = new(bytes.Buffer)
	}
}

func BenchmarkStringAppend(b *testing.B) {
	var q []byte
	for i := 0; i < b.N; i++ {
		for j := 0; j < LIMIT; j++ {
			q = append(q, s1...)
		}
		q = nil
	}
	b.ReportAllocs()
}

func BenchmarkBytesJoin(b *testing.B) {
	var q []byte
	for i := 0; i < b.N; i++ {
		for j := 0; j < LIMIT; j++ {
			q = bytes.Join([][]byte{q, b1}, nil)
		}
		q = nil
	}
	b.ReportAllocs()
}

func BenchmarkBytesAppend(b *testing.B) {
	var q []byte
	for i := 0; i < b.N; i++ {
		for j := 0; j < LIMIT; j++ {
			q = append(q, b1...)
		}
		q = nil
	}
	b.ReportAllocs()
}

func BenchmarkBytesWrite(b *testing.B) {
	q := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		for j := 0; j < LIMIT; j++ {
			q.Write(b1)
		}
		q = new(bytes.Buffer)
	}
}
