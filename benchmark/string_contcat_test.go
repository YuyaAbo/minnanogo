package benchmark

import (
	"fmt"
	"testing"
)

// seedはベンチマーク用のトークンをつくる
// 長さを受け取り、指定された長さの文字列のスライスを生成する
// 今回は、単純に"a"をn個並べたスライスを生成する
func seed(n int) []string {
	s := make([]string, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, "a")
	}
	return s
}

// benchはベンチマーク用のヘルパ
// テストしたい文字列の組み合わせ長と、文字列結合のための
// 手続きを渡す。それについてベンチマークを実行させる
func bench(b *testing.B, n int, f func(...string) string) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f(seed(n)...)
	}
}

func BenchmarkConcatenate(b *testing.B) {
	benchCases := []struct {
		name string
		n    int
		f    func(...string) string
	}{
		{"Cat", 3, Cat},
		{"Buf", 3, Buf},
		{"Join", 3, Join},
		{"Cat", 100, Cat},
		{"Buf", 100, Buf},
		{"Join", 100, Join},
		{"Cat", 10000, Cat},
		{"Buf", 10000, Buf},
		{"Join", 10000, Join},
	}
	for _, c := range benchCases {
		b.Run(fmt.Sprintf("%s%d", c.name, c.n),
			func(b *testing.B) {
				bench(b, c.n, c.f)
			})
	}
}
