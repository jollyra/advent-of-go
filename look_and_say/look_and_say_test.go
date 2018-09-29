package main

import "testing"

func TestLookSay(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	}

	for _, c := range cases {
		got := lookSay(c.in)
		if got != c.want {
			t.Errorf("LookSay(%s) == %s, want %s", c.in, got, c.want)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
