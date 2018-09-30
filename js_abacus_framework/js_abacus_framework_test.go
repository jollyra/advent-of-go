package main

import "testing"

func TestIncrement(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"[1,2,3]", 6},
		{"[10,20,-31]", -1},
		{`{"a":2,"b":4}`, 6},
		{"[[[3]]]", 3},
		{`{"a":{"b":4},"c":-1}`, 3},
		{`{"a":[-1,1]}`, 0},
		{`[-1,{"a":1}]`, 0},
		{"[]", 0},
		{"{}", 0},
	}
	for _, c := range cases {
		got := sum(c.in)
		if got != c.want {
			t.Errorf("sum(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
