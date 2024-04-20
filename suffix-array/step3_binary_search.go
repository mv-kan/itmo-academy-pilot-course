package main

import "fmt"

func Solution3(s string, substrs ...string) []bool {
	idxs := Solution2(s)
	compare := func(s1 string, s2 string, n int) int {
		for i := 0; i < n; i++ {
			if s1[i] > s2[i] {
				return -1
			} else if s1[i] < s2[i] {
				return 1
			}
		}
		return 0
	}
	contains := func(idxs []int, s string, substr string) bool {
		n := len(substr)
		idx := int((len(s)+1)/2) - 1
		for i := 1; i < len(s); i *= 2 {
			comp := compare(s[idxs[idx]:], substr, n)
			if comp == -1 {
				idx = idx + idx/2
			} else if comp == 1 {
				idx = idx - idx/2
			} else {
				return true
			}
		}
		return false
	}
	return []bool{contains(idxs, s, substrs[0])}
}

func Solution3Main() {
	fmt.Printf("%t", Solution3("ababba", "aba"))
}
