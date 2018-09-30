package main

import (
	"fmt"
	"strings"

	"github.com/jollyra/stringutil"
)

func increment(s string) string {
	var b strings.Builder
	var i int
	for i = len(s) - 1; i >= 0; i-- {
		r := s[i]
		var rNew rune
		if r == 'z' {
			rNew = 'a'
			fmt.Fprintf(&b, "%s", string(rNew))
		} else {
			rNew = rune(r + 1)
			fmt.Fprintf(&b, "%s", string(rNew))
			break
		}
	}

	var next string
	if i > 0 {
		next = s[:i] + stringutil.Reverse(b.String())
	} else {
		next = stringutil.Reverse(b.String())
	}

	return next

}

func increasingStraight(s string, length int) bool {
	for i := 0; i <= len(s)-length; i++ {
		r := int(s[i])
		straight := true
		for j := 1; j < length; j++ {
			if r != int(s[i+j])-j {
				straight = false
			}
		}
		if straight == true {
			return true
		}
	}
	return false
}

func countDifferentNonOverlappingPairs(s string) int {
	count := 0
	seen := ""
	i := 0
	for i < len(s)-1 {
		if s[i] == s[i+1] && !strings.ContainsRune(seen, rune(s[i])) {
			seen += string(s[i])
			count++
			i += 2
		} else {
			i++
		}
	}
	return count
}

func validatePassword(s string) bool {
	return increasingStraight(s, 3) &&
		!strings.ContainsAny(s, "iol") &&
		countDifferentNonOverlappingPairs(s) >= 2
}

func nextPassword(oldPassword string, validator func(string) bool) string {
	candidatePassword := oldPassword
	for {
		candidatePassword = increment(candidatePassword)
		if validator(candidatePassword) {
			return candidatePassword
		}
	}
}

func main() {
	oldPassword := "cqjxjnds"
	fmt.Println(nextPassword(oldPassword, validatePassword))
}
