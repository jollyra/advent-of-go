package main

import "testing"

func TestWhatFloor(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, c := range cases {
		got := WhatFloor(c.in)
		if got != c.want {
			t.Errorf("WhatFloor(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestDistanceToFloor(t *testing.T) {
	cases := []struct {
		instructions string
		floor        int
		want         int
	}{
		{")", -1, 1},
		{"()())", -1, 5},
	}

	for _, c := range cases {
		got := DistanceToFloor(c.instructions, c.floor)
		if got != c.want {
			t.Errorf("DistanceToFloor(%q, %d) == %d, want %d",
				c.instructions, c.floor, got, c.want)
		}
	}
}
