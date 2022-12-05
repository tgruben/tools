package lists

import (
	"math/rand"
	"sort"
)

type signedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type unsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type float interface {
	~float32 | ~float64
}

type number interface {
	signedInt | unsignedInt | float
}

type enumerable interface {
	number | ~string
}

func Shuffle[T any](list []T) {
	e := len(list)
	for i := range list {
		j := rand.Intn(e)
		list[i], list[j] = list[j], list[i]
	}
}

func First[T any](list []T) T {
	return list[0]
}

func Rest[T any](list []T) []T {
	if len(list) == 0 {
		return list
	}
	return list[1:]
}

func Rotate[T any](list []T){
	if len(list) == 1 {
		return 
	}
	first:=list[0]
	copy(list,list[1:])
	list[len(list)-1]=first
}

func Sum[T number](list []T) T {
	var sum T
	for _, v := range list {
		sum += v
	}
	return sum
}

func Sort[T enumerable](list []T) {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
}

func SortDesc[T enumerable](list []T) {
	sort.Slice(list, func(i, j int) bool {
		return list[i] > list[j]
	})
}

func Min[T number](list []T) T {
	if len(list) == 0 {
		return T(0)
	}

	var min T = list[0]

	for _, v := range list {
		if v < min {
			min = v
		}
	}
	return min
}

func Max[T number](list []T) T {
	if len(list) == 0 {
		return T(0)
	}

	var max T = list[0]

	for _, v := range list {
		if v > max {
			max = v
		}
	}
	return max
}

func Fold[T any](list []T, f func(T, T) T) T {
	if len(list) == 1 {
		return list[0]
	}
	var accum T
	for i := range list {
		if i == 0 {
			accum = list[i]
		} else {
			accum = f(accum, list[i])
		}
	}
	return accum
}

func Contains[T comparable](list []T, value T) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}

	return false
}


func Filter[T any](list []T, f func(T) bool) []T {
	result := make([]T, 0, len(list))
	for _, v := range list {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func Diff[T comparable](s1, s2 []T) []T {
	if len(s1) == 0 {
		return []T{}
	}

	var result []T

	for _, v := range s1 {
		if !Contains(s2, v) {
			result = append(result, v)
		}
	}

	return result
}

func Intersect[T comparable](s1, s2 []T) []T {
	if len(s1) == 0 || len(s2) == 0 {
		return []T{}
	}

	var result []T

	for _, v := range s1 {
		if Contains(s2, v) {
			result = append(result, v)
		}
	}

	return result
}

func Any[T comparable](list []T, f func(T) bool) bool {
	for _, v := range list {
		if f(v) {
			return true
		}
	}

	return false
}

func All[T comparable](list []T, f func(T) bool) bool {
	for _, v := range list {
		if !f(v) {
			return false
		}
	}

	return true
}

func Select[T, K comparable](list []T, f func(T) K) []K {
	result := make([]K, 0, len(list))
	for _, v := range list {
		result = append(result, f(v))
	}

	return result
}

func SelectI[T, K comparable](list []T, f func(T) bool) []int {
	var result []int
	for i, v := range list {
		if f(v) {
			result = append(result, i)
		}
	}
	return result
}

func Push[T any](list []T,item T)[]T {
	return append([]T{item},list...)
}

func Pop[T any](list []T)(T,[]T) {
	item := list[0]
	list = list[1:]
	return item,list
}
