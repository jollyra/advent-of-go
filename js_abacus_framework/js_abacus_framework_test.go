package main

import "testing"

func TestSum(t *testing.T) {
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

func TestSumJSONWithIgnore(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"[1,2,3]", 6},
		{`[1,{"c":"red","b":2},3]`, 4},
		{`{"d":"red","e":[1,2,3,4],"f":5}`, 0},
		{`[1,"red",5]`, 6},
	}
	for _, c := range cases {
		got := sumJSONWithIgnore([]byte(c.in))
		if got != c.want {
			t.Errorf("sumJSON(%v) == %d, want %d", c.in, got, c.want)
		}
	}
}
