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
		start := 0
		end := len(s) - 1
		mid := int((len(s) + 1) / 2)
		for {
			sub := s[idxs[mid]:]
			comp := compare(sub, substr, n)
			if comp == -1 {
				if mid == end {
					break
				}
				end = mid
				mid -= (end - start + 1) / 2
			} else if comp == 1 {
				if mid == start {
					break
				}
				start = mid
				mid += (end - start + 1) / 2
			} else {
				return true
			}
		}
		return false
	}
	s += "$"
	answs := make([]bool, len(substrs))
	for i, substr := range substrs {
		answs[i] = contains(idxs, s, substr)
	}
	return answs
}

func Solution3_2(s string, substrs ...string) []int {
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
	findBegin := func(idxs []int, s string, substr string) int {
		n := len(substr)
		start := 0
		end := len(s) - 1
		mid := int((len(s) + 1) / 2)
		resultIdx := -1
		for {
			sub := s[idxs[mid]:]
			comp := compare(sub, substr, n)
			if comp == -1 {
				if mid == end {
					break
				}
				end = mid
				mid -= (end - start + 1) / 2
			} else if comp == 1 {
				if mid == start {
					break
				}
				start = mid
				mid += (end - start + 1) / 2
			} else {
				if resultIdx == mid {
					break
				}
				start = 0
				end = mid
				resultIdx = mid
				mid -= (end - start + 1) / 2
			}
		}
		return resultIdx
	}
	findEnd := func(idxs []int, s string, substr string) int {
		n := len(substr)
		start := 0
		end := len(s) - 1
		mid := int((len(s) + 1) / 2)
		resultIdx := -1
		for {
			sub := s[idxs[mid]:]
			comp := compare(sub, substr, n)
			if comp == -1 {
				if mid == end {
					break
				}
				end = mid
				mid -= (end - start + 1) / 2
			} else if comp == 1 {
				if mid == start {
					break
				}
				start = mid
				mid += (end - start + 1) / 2
			} else {
				if resultIdx == mid {
					break
				}
				start = mid
				end = len(s) - 1
				resultIdx = mid
				mid += (end - start + 1) / 2
			}
		}
		return resultIdx
	}
	count := func(idxs []int, s string, substr string) int {
		begin := findBegin(idxs, s, substr)
		if begin == -1 {
			return 0
		}
		end := findEnd(idxs, s, substr)
		return end - begin + 1
	}
	answs := make([]int, len(substrs))
	s += "$"
	for i, substr := range substrs {
		answs[i] = count(idxs, s, substr)
	}
	return answs
}

func Solution3_2Main() {
	// s := "ababba"
	// substrs := []string{"baba", "ba", "abba"}
	s := ""
	n := 0
	fmt.Scanf("%s\n%d\n", &s, &n)
	substrs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &substrs[i])
	}
	answs := Solution3_2(s, substrs...)
	for _, answ := range answs {
		fmt.Printf("%d\n", answ)
	}
}

func Solution3Main() {
	// s := "a"
	// substrs := []string{"a", "a", "ab"}
	// s := "codeforces"
	// substrs := []string{"s", "t", "c", "d", "math", "code", "forces"}
	s := ""
	n := 0
	fmt.Scanf("%s\n%d\n", &s, &n)
	substrs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &substrs[i])
	}
	answs := Solution3(s, substrs...)
	for _, answ := range answs {
		if answ {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
	// fmt.Printf("%s\n%d\n%v\n", s, n, substrs)
}
