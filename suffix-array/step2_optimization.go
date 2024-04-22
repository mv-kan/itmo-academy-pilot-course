package main

import (
	"fmt"
	"sort"
)

// suffix array with radix sorting
// O(n log n)
func Solution2(s string) []int {
	s += "$"
	n := len(s)
	p := make([]int, n)
	c := make([]int, n)
	{
		type pair struct {
			idx  int
			char byte
		}
		a := make([]pair, n)
		for i := 0; i < n; i++ {
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
		for i := 0; i < n; i++ {
			p[i] = a[i].idx
		}
		c[p[0]] = 0
		for i := 1; i < n; i++ {
			if a[i-1].char == a[i].char {
				c[p[i]] = c[p[i-1]]
			} else {
				c[p[i]] = c[p[i-1]] + 1
			}
		}
	}
	pSort := func(p []int, c []int) []int {
		n := len(p)
		count := make([]int, n)
		for _, v := range c {
			count[v]++
		}
		pos := make([]int, n)
		pos[0] = 0
		for i := 1; i < n; i++ {
			pos[i] = pos[i-1] + count[i-1]
		}
		pCopy := make([]int, n)
		for _, v := range p {
			i := c[v]
			pCopy[pos[i]] = v
			pos[i]++
		}
		return pCopy
	}
	k := 0
	for 1<<k < n {
		for i := 0; i < n; i++ {
			p[i] = (p[i] - 1<<k + n) % n
		}

		p = pSort(p, c)

		cNew := make([]int, n)
		cNew[p[0]] = 0
		for i := 1; i < n; i++ {
			type pair struct {
				a, b int
			}
			prev := pair{
				a: c[p[i-1]],
				b: c[(p[i-1]+1<<k)%n],
			}
			now := pair{
				a: c[p[i]],
				b: c[(p[i]+1<<k)%n],
			}
			if now == prev {
				cNew[p[i]] = cNew[p[i-1]]
			} else {
				cNew[p[i]] = cNew[p[i-1]] + 1
			}
		}
		c = cNew
		k++
	}
	return p
}

// simple radix sort
// for efficiency this is a must: max(a) <= n -1
// Usage
// arr := radixSort([]int{6, 0, 1, 5, 9, 4, 2, 3, 7, 8})
// fmt.Printf("%v\n", arr)
func radixSort(a []int) []int {
	n := len(a)
	count := make([]int, n)
	for _, v := range a {
		count[v]++
	}

	pos := make([]int, n)
	pos[0] = 0
	for i := 1; i < n; i++ {
		pos[i] = pos[i-1] + count[i-1]
	}
	aCopy := make([]int, n)
	for _, v := range a {
		aCopy[pos[v]] = v
		pos[v]++
	}
	return aCopy
}

func Solution2Main() {
	s := ""
	fmt.Scanf("%s", &s)
	a := Solution2(s)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
	// // show with substr
	// s += "$"
	// for i := 0; i < n+1; i++ {
	// 	fmt.Printf("%d\t%s", a[i], s[a[i]:])
	// 	fmt.Println()
	// }
	// fmt.Println()
}
