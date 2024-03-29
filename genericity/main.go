package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// 切片定义泛型
type Slice[T int | float32 | float64] []T

// map 定义泛型
type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

type Test2[T int | string | float32] struct {
	test11 T
	Test21 T
	test31 T
}

func main() {
	strs := []string{"Hao", "Chen", "MegaEase"}
	upstrs := gMap(strs, func(s string) string  {
		return strings.ToUpper(s)
	})
	fmt.Println(upstrs)
	nums := []int {0,1,2,3,4,5,6,7,8,9}
	dict := []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
	strs =  gMap(nums, func (elem int) string  {
		return  dict[elem]
	})
	fmt.Println(dict)



	var f Tree[int]
	f.cmp = test1
	for i := 0; i < 100; i++ {
		f.Insert(rand.Intn(100))
	}
	fmt.Println(f.root)

	var q Queue[int]

	for i := 0; i < 5; i++ {
		q.Put(rand.Intn(100))
	}
	fmt.Println(q.elements)
	fmt.Println(q.Size())
	q.Pop()
	fmt.Println(q.elements)

	fmt.Println(q.Size())
}


func gFilter[T any] (arr []T, in bool, f func(T) bool) []T {
	result := []T{}
	
}


func gMap[T1 any , T2 any] (arr []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(arr))

	for i , elem := range arr {
		result[i] = f(elem)
	}
	return result
}

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Put(value T) {
	q.elements = append(q.elements, value)
}

func (q *Queue[T]) Pop() (T, bool) {
	var value T
	if len(q.elements) == 0 {
		return value, true
	}

	value = q.elements[0]
	q.elements = q.elements[1:]
	return value, len(q.elements) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}

func test1(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func MapKeys[Key comparable, Val any](m map[Key]Val) []Key {
	s := make([]Key, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

type Tree[T any] struct {
	cmp  func(T, T) int
	root *node[T]
}

type node[T any] struct {
	left, right *node[T]
	val         T
}

func (bt *Tree[T]) find(val T) **node[T] {
	pl := &bt.root
	for *pl != nil {
		switch cmp := bt.cmp(val, (*pl).val); {
		case cmp < 0:
			pl = &(*pl).left
		case cmp > 0:
			pl = &(*pl).right
		default:
			return pl
		}
	}
	return pl
}

func (bt *Tree[T]) Insert(val T) bool {
	pl := bt.find(val)
	if *pl != nil {
		return false
	}
	*pl = &node[T]{val: val}

	return true
}
