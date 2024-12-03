package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
    // Open the file.
    f, _ := os.Open("input.txt")
    // Create a new Scanner for the file.
    scanner := bufio.NewScanner(f)
    // Loop over all lines in the file and print them.
	
	numberOfLines := 0

	safeReports := 0

    for scanner.Scan() {
      line := scanner.Text()
      numberOfLines++

	  slice := strings.Fields(line)

	   if (isReportSafe(slice)) {
		  safeReports++
	   }
    }

	fmt.Printf("SAFE REPORTS: %d", safeReports)
}

func isReportSafe(line []string) bool {
	safeReport := true

	// if it has repeats, it's not safe
	lineCompact := slices.Compact(line)
	if(len(line) != len(lineCompact)) {
		safeReport = false
	}

	// determine if it's ascending or descending in order
	ascending := false
	descending := false

    if (slices.IsSorted(line)) {
        ascending = true
	} else {
		descending = slices.IsSortedFunc(line, func(a, b string) int {
			return strings.Compare(b, a)
		})
	}

	if(!ascending && !descending) {
		// not safe
		safeReport = false;
	}

    // ensure adjacent numbers are within a diff of 1-3
    maxLength := len(line)

    for idx, num := range line {
		if(idx < maxLength - 1) {

			currentNum, _ := strconv.Atoi(num)
			nextNum, _ := strconv.Atoi(line[idx + 1])

			diff := currentNum - nextNum
			diffFloat := float64(diff)
            
			 if (math.Abs(diffFloat) > 3) {
                safeReport = false
			 }
		}
	}

	return safeReport
}