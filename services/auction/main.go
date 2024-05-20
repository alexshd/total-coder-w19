package main

import (
	"log"
	"net/http"

	"github.com/alexshd/total-coder-w19/services/auction/api"
)

func main() {
	handler := http.HandlerFunc(api.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}

type Tree[V any] struct {
	left  *Tree[V]
	right *Tree[V]
	value V
}
type Iter[V any] struct {
	stk []*Tree[V]
}

func (t *Tree[V]) NewIter() *Iter[V] {
	it := new(Iter[V])
	for ; t != nil; t = t.left {
		it.stk = append(it.stk, t)
	}

	return it
}

func (it *Iter[V]) Next() (v V, ok bool) {
	if len(it.stk) == 0 {
		return v, false
	}

	t := it.stk[len(it.stk)-1]
	v = t.value
	it.stk = it.stk[:len(it.stk)-1]
	for t = t.right; t != nil; t = t.left {
		it.stk = append(it.stk, t)
	}

	return v, true
}

func (t *Tree[V]) All(f func(v V)) {
	if t != nil {
		t.left.All(f)
		f(t.value)
		t.right.All(f)
	}
}

func SameValues[V comparable](t1, t2 *Tree[V]) bool {
	c1 := make(chan V)
	c2 := make(chan V)
	go gopher(c1, t1.All)
	go gopher(c2, t2.All)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 && !ok2 {
			return true
		}
	}
}

func gopher[V any](c chan<- V, all func(func(V))) {
	all(func(v V) { c <- v })

	close(c)
}
