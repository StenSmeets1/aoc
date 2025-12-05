package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Range struct {
    start, end int
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var ranges []Range
    var available []int
    parsingRanges := true

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" {
            parsingRanges = false
            continue
        }

        if parsingRanges {
            parts := strings.Split(line, "-")
                        start, err1 := strconv.Atoi(parts[0])
            end, err2 := strconv.Atoi(parts[1])
            if err1 != nil || err2 != nil {
                panic(err1)
		panic(err2)
            }
            ranges = append(ranges, Range{start, end})
        } else {
            id, err := strconv.Atoi(line)
            if err != nil {
                panic(err)
	    }
            available = append(available, id)
        }
    }

   

    freshCount := 0
    for _, id := range available {
        isFresh := false
        for _, r := range ranges {
            if id >= r.start && id <= r.end {
                isFresh = true
                break
            }
        }
        if isFresh {
            freshCount++
        }
    }

    fmt.Println(freshCount)
}

