package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/edu/course/2/lesson/2/1/practice/contest/269100/problem/A
// suffix array with sorting
// O(n log^2n)
func Solution1(s string) []int {
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
		a := make([]elem, len(s))
		for i := 0; i < len(s); i++ {
			nextIdx := (i + (1 << k)) % len(s)
			a[i] = elem{
				idx: i,
				a:   c[i],
				b:   c[nextIdx],
			}
		}
		sort.Slice(a, func(i, j int) bool {
			if a[i].a == a[j].a {
				if a[i].b == a[j].b {
					return a[i].idx > a[j].idx
				}
				return a[i].b < a[j].b
			} else {
				return a[i].a < a[j].a
			}
		})
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

func Solution1Main() {
	s := ""
	fmt.Scanf("%s", &s)
	a := Solution1(s)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
	// // show with substr
	// s += "$"
	// for i := 0; i < len(a); i++ {
	// 	fmt.Printf("%d\t%s", a[i], s[a[i]:])
	// 	fmt.Println()
	// }
	// fmt.Println()
}
