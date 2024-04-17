package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/edu/course/2/lesson/2/1/practice/contest/269100/problem/A
func Solution(s string) []int {
	type pair struct {
		idx    int
		str    string
		cleft  int
		cright int
	}
	s += "$"
	p := []pair{}
	for i, ch := range s {
		p = append(p, pair{
			idx: i,
			str: string(ch),
		})
	}
	// k := 0
	sort.Slice(p, func(i, j int) bool {
		return byte(p[i].str[0]) < byte(p[j].str[0])
	})
	c := make([]int, 0, len(p))
	class := 0
	prev := p[0]
	c = append(c, class)
	for i := 1; i < len(p); i++ {
		if prev.str != p[i].str {
			class++
		}
		c = append(c, class)
		prev = p[i]
	}
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", c)
	return nil
}

func main() {
	s := ""
	fmt.Scanf("%s", &s)
	Solution(s)
}
