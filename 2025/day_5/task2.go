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

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" {
            break
        }
        parts := strings.Split(line, "-")
        if len(parts) != 2 {
            panic("invalid range line: " + line)
        }
        start, err1 := strconv.Atoi(parts[0])
        end, err2 := strconv.Atoi(parts[1])
        if err1 != nil || err2 != nil {
            panic(err1)
	    panic(err2)
        }
        ranges = append(ranges, Range{start, end})
    }


    merged := mergeRanges(ranges)

    totalFresh := 0
    for _, r := range merged {
        totalFresh += r.end - r.start + 1
    }

    fmt.Println(totalFresh)
}

func mergeRanges(ranges []Range) []Range {
    if len(ranges) == 0 {
        return nil
    }

    for i := 0; i < len(ranges); i++ {
        for j := i + 1; j < len(ranges); j++ {
            if ranges[j].start < ranges[i].start {
                ranges[i], ranges[j] = ranges[j], ranges[i]
            }
        }
    }

    merged := []Range{ranges[0]}
    for _, r := range ranges[1:] {
        last := &merged[len(merged)-1]
        if r.start <= last.end+1 { 
            if r.end > last.end {
                last.end = r.end
            }
        } else {
            merged = append(merged, r)
        }
    }
    return merged
}

