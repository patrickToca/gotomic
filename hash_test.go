package gotomic

import (
	"testing"
	"fmt"
)

type key string
func (self key) HashCode() uint32 {
	var sum uint32
	for c := range self {
		sum += uint32(c)
	}
	return sum
}
func (self key) Equals(t thing) bool {
	if s, ok := t.(key); ok {
		return s == self
	}
	return false
}

func TestPutGet(t *testing.T) {
	h := newHash()
	fmt.Println(h)
	h.put(key("a"), "b")
	fmt.Println(h)
}
