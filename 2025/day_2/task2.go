package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type interval struct{ L, R int64 }

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	line := strings.TrimSpace(string(data))
	ranges := parseRanges(line)
	if len(ranges) == 0 {
		fmt.Println(0)
		return
	}

	merged := mergeIntervals(ranges)
	minR, maxR := merged[0].L, merged[len(merged)-1].R
	maxDigits := numDigits(maxR)

	found := make(map[int64]bool)

	for k := 1; k <= maxDigits; k++ {
		sMin := int64Pow(10, k-1)
		sMax := int64Pow(10, k) - 1
		maxT := maxDigits / k
		for t := 2; t <= maxT; t++ {
			for s := sMin; s <= sMax; s++ {
				sStr := strconv.FormatInt(s, 10)
				repStr := strings.Repeat(sStr, t)

				if len(repStr) > 19 { 
					break
				}
				N, err := strconv.ParseInt(repStr, 10, 64)
				if err != nil {
					break
				}

				if N > maxR {
					break
				}
				if N < minR {
					continue
				}
				if inAnyInterval(N, merged) {
					found[N] = true
				}
			}
		}
	}

	var total int64 = 0
	for v := range found {
		total += v
	}
	fmt.Println(total)
}

func parseRanges(line string) []interval {
	if strings.TrimSpace(line) == "" {
		return nil
	}
	parts := strings.Split(line, ",")
	out := make([]interval, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		pr := strings.Split(p, "-")
		if len(pr) != 2 {
			continue
		}
		L, err1 := strconv.ParseInt(pr[0], 10, 64)
		R, err2 := strconv.ParseInt(pr[1], 10, 64)
		if err1 != nil || err2 != nil {
			continue
		}
		if L > R {
			L, R = R, L
		}
		out = append(out, interval{L, R})
	}
	return out
}

func mergeIntervals(iv []interval) []interval {
	if len(iv) == 0 {
		return nil
	}
	sort.Slice(iv, func(i, j int) bool { return iv[i].L < iv[j].L })
	res := make([]interval, 0, len(iv))
	cur := iv[0]
	for i := 1; i < len(iv); i++ {
		if iv[i].L <= cur.R+1 {
			if iv[i].R > cur.R {
				cur.R = iv[i].R
			}
		} else {
			res = append(res, cur)
			cur = iv[i]
		}
	}
	res = append(res, cur)
	return res
}

func inAnyInterval(x int64, intervals []interval) bool {
	i := sort.Search(len(intervals), func(i int) bool { return intervals[i].L > x })
	if i > 0 {
		c := intervals[i-1]
		if x >= c.L && x <= c.R {
			return true
		}
	}
	return false
}

func numDigits(n int64) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func int64Pow(a, b int) int64 {
	if b == 0 {
		return 1
	}
	res := int64(1)
	base := int64(a)
	for b > 0 {
		if b&1 == 1 {
			res *= base
		}
		base *= base
		b >>= 1
	}
	return res
}
