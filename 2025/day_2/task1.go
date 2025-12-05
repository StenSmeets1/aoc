package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(data))
	fmt.Println(sumDoubleIDs(input))
}

func sumDoubleIDs(line string) int64 {
	ranges := parseRanges(line)

	maxR := int64(0)
	for _, r := range ranges {
		if r[1] > maxR {
			maxR = r[1]
		}
	}

	maxDigits := numDigits(maxR)
	maxK := maxDigits / 2
	
	found := make(map[int64]bool)

	for k := 1; k <= maxK; k++ {
		base := int64Pow(10, k)
		mult := base + 1

		sMin := int64Pow(10, k-1)
		sMax := base -1

		for _, r := range ranges {
			L, R := r[0], r[1]

			sLow := ceilDiv(L, mult)
			sHigh := R / mult

			sStart := maxInt64(sLow, sMin)
			sEnd := minInt64(sHigh, sMax)

			if sStart <= sEnd {
				for s := sStart; s <= sEnd; s++ {
					N := s * mult
					found[N] = true
				}
			}
		}
	}

	var total int64 = 0
	for v := range found {
		total += v
	}
	return total
}

func parseRanges(line string) [][2]int64 {
	parts := strings.Split(line, ",")
	ranges := make([][2]int64, 0)
	for _, p := range parts {
		if p == "" {
			continue
		}
		pair := strings.Split(p, "-")
		L, _ := strconv.ParseInt(pair[0], 10, 64)
		R, _ := strconv.ParseInt(pair[1], 10, 64)
		ranges = append(ranges, [2]int64{L, R})
	}
	return ranges
}

func numDigits(n int64) int {
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func int64Pow(a, b int) int64 {
	result := int64(1)
	for i := 0; i < b; i++ {
		result *= int64(a)
	}
	return result
}

func ceilDiv(a, b int64) int64 {
	return (a + b - 1) / b
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
