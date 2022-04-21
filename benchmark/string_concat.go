package benchmark

import (
	"bytes"
	"strings"
)

// Cat Catは+=演算子を使って文字列を結合する
func Cat(ss ...string) string {
	var r string
	for _, s := range ss {
		r += s
	}
	return r
}

// Buf Bufはbytes.Bufferを使って文字列を結合する
func Buf(ss ...string) string {
	var b bytes.Buffer
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}

// Join Joinはstrings.Joinを使って文字列を結合する
func Join(ss ...string) string {
	return strings.Join(ss, "")
}
