package main

import (
	"fmt"
	"sort"
)

func Solution4(s string) ([]int, []int) {
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

	lcp := make([]int, n)
	k = 0
	for i := 0; i < n-1; i++ {
		pi := c[i]
		j := p[pi-1]
		for s[i+k] == s[j+k] {
			k++
		}
		lcp[pi] = k
		k = max(k-1, 0)
	}

	return p, lcp
}

func Solution4Main() {
	// s := "ababba"
	s := ""
	n := 0
	fmt.Scanf("%s\n%d\n", &s, &n)
	substrs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &substrs[i])
	}
	p, lcp := Solution4(s)
	s += "$"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d. lcp=%d %d - %s\n", i, lcp[i], p[i], s[p[i]:])
	}
}
