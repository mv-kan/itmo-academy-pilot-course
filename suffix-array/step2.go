package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/edu/course/2/lesson/2/1/practice/contest/269100/problem/A
// suffix array with sorting
// O(n log^2n)
func Solution2(s string) []int {
	s += "$"
	p := make([]int, len(s))
	c := make([]int, len(s))
	{
		type pair struct {
			idx  int
			char byte
		}
		a := make([]pair, len(s))
		for i := 0; i < len(s); i++ {
			a[i] = pair{
				idx:  i,
				char: s[i],
			}
		}
		sort.Slice(a, func(i, j int) bool {
			if a[i].char == a[j].char {
				return i > j
			}
			return a[i].char < a[j].char
		})
		for i := 0; i < len(s); i++ {
			p[i] = a[i].idx
		}
		c[p[0]] = 0
		for i := 1; i < len(s); i++ {
			if a[i-1].char == a[i].char {
				c[p[i]] = c[p[i-1]]
			} else {
				c[p[i]] = c[p[i-1]] + 1
			}
		}
	}
	k := 0
	for 1<<k < len(s) {
		type elem struct {
			idx int
			a   int
			b   int
		}
		radixSort := func(a []elem) []elem {
			var result []elem
			{
				cnt := make([]int, len(a))
				for _, v := range a {
					cnt[v.b]++
				}
				aCopy := make([]elem, len(a))
				pos := make([]int, len(a))
				pos[0] = 0
				for i := 1; i < len(a); i++ {
					pos[i] = pos[i-1] + cnt[i-1]
				}
				for _, v := range a {
					i := v.b
					aCopy[pos[i]] = v
					pos[i]++
				}
				result = aCopy
			}
			{
				cnt := make([]int, len(a))
				for _, v := range a {
					cnt[v.a]++
				}
				aCopy := make([]elem, len(a))
				pos := make([]int, len(a))
				pos[0] = 0
				for i := 1; i < len(a); i++ {
					pos[i] = pos[i-1] + cnt[i-1]
				}
				for _, v := range result {
					i := v.a
					aCopy[pos[i]] = v
					pos[i]++
				}
				result = aCopy
			}
			return result
		}
		a := make([]elem, len(s))
		for i := 0; i < len(s); i++ {
			nextIdx := (i + (1 << k)) % len(s)
			a[i] = elem{
				idx: i,
				a:   c[i],
				b:   c[nextIdx],
			}
		}
		fmt.Printf("Before:\t%+v\n", a)
		a = radixSort(a)
		fmt.Printf("After:\t%+v\n", a)
		for i := 0; i < len(s); i++ {
			p[i] = a[i].idx
		}
		c[p[0]] = 0
		for i := 1; i < len(s); i++ {
			if a[i-1].a == a[i].a && a[i-1].b == a[i].b {
				c[p[i]] = c[p[i-1]]
			} else {
				c[p[i]] = c[p[i-1]] + 1
			}
		}
		k++
	}
	return p
}

// simple radix sort
// for efficiency this is a must: max(a) <= len(a) -1
// Usage
// arr := radixSort([]int{6, 0, 1, 5, 9, 4, 2, 3, 7, 8})
// fmt.Printf("%v\n", arr)
func radixSort(a []int) []int {
	count := make([]int, len(a))
	for _, v := range a {
		count[v]++
	}

	pos := make([]int, len(a))
	pos[0] = 0
	for i := 1; i < len(a); i++ {
		pos[i] = pos[i-1] + count[i-1]
	}
	aCopy := make([]int, len(a))
	for _, v := range a {
		aCopy[pos[v]] = v
		pos[v]++
	}
	return aCopy
}
func Solution2Main() {

	s := "ababba"
	// fmt.Scanf("%s", &s)
	a := Solution2(s)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
	// show with substr
	s += "$"
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d\t%s", a[i], s[a[i]:])
		fmt.Println()
	}
	fmt.Println()
}
