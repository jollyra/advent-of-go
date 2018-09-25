package main

import "testing"

func TestLookSay(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		// {"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	}

	for _, c := range cases {
		got := LookSayFast(c.in)
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

func TestChomps(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"1", []string{
			"1",
		}},
		{"11", []string{
			"11",
		}},
		{"21", []string{
			"2",
			"1",
		}},
		{"1211", []string{
			"1",
			"2",
			"11",
		}},
		{"111221", []string{
			"111",
			"22",
			"1",
		}},
	}

	for _, c := range cases {
		got := chomps(c.in)
		if !equal(got, c.want) {
			t.Errorf("Chomps(%s) == %s, want %s", c.in, got, c.want)
		}
	}
}
