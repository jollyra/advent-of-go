package main

import "testing"

func TestIncrement(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"a", "b"},
		{"aa", "ab"},
		{"az", "ba"},
		{"azzz", "baaa"},
	}
	for _, c := range cases {
		got := increment(c.in)
		if got != c.want {
			t.Errorf("increment(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestIncreasingStraight(t *testing.T) {
	cases := []struct {
		in     string
		length int
		want   bool
	}{
		{"abc", 3, true},
		{"abd", 3, false},
		{"aabcdd", 3, true},
		{"ab", 3, false},
	}
	for _, c := range cases {
		got := increasingStraight(c.in, c.length)
		if got != c.want {
			t.Errorf("increaingStraight(%q, %d) == %t, want %t",
				c.in, c.length, got, c.want)
		}
	}
}

func TestCountNonOverlappingPairs(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"ab", 0},
		{"aa", 1},
		{"aaa", 1},
		{"aaaa", 1},
		{"aabbccc", 3},
	}
	for _, c := range cases {
		got := countDifferentNonOverlappingPairs(c.in)
		if got != c.want {
			t.Errorf("countDifferentNonOverlappingPairs(%q) == %d, "+
				"want %d", c.in, got, c.want)
		}
	}
}

func TestNextPassword(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
	}
	for _, c := range cases {
		got := nextPassword(c.in, validatePassword)
		if got != c.want {
			t.Errorf("nextPassword(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
