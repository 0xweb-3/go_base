package main

import "fmt"

type Set map[int]struct{}

func (s Set) Put(x int) {
	s[x] = struct{}{}
}

func (s Set) Has(x int) (exists bool) {
	_, exists = s[x]
	return
}

func (s Set) Delete(x int) {
	delete(s, x)
}

func main() {
	// 在一些场合，空结构体经常被用作集合中的占位符，尤其是当你只关心某些元素的存在性，而不关心它们的具体值时。
	s := make(Set)
	s.Put(1)
	s.Put(2)
	fmt.Println(s.Has(1))
	s.Delete(1)
	fmt.Println(s.Has(1))
}
