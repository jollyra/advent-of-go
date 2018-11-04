package main

import "testing"

// Part 1
func TestCalibrate(t *testing.T) {
	var mutations = []mutation{
		mutation{"H", "HO"},
		mutation{"H", "OH"},
		mutation{"O", "HH"},
	}

	cases := []struct {
		molecule string
		want     int
	}{
		{"HOH", 4},
		{"HOHOHO", 7},
	}

	for _, c := range cases {
		got := calibrate(mutations, c.molecule)
		if got != c.want {
			t.Errorf("calibrate(%v, %v) == %d, want %d",
				mutations, c.molecule, got, c.want)
		}
	}
}

// Part 2
func TestReverseEngineer(t *testing.T) {
	var mutations = []mutation{
		mutation{"e", "H"},
		mutation{"e", "O"},
		mutation{"H", "HO"},
		mutation{"H", "OH"},
		mutation{"O", "HH"},
	}

	cases := []struct {
		molecule string
		want     int
	}{
		{"HOH", 3},
		{"HOHOHO", 6},
	}

	for _, c := range cases {
		got := reverseEngineer(mutations, c.molecule)
		if got != c.want {
			t.Errorf("reverseEngineer(%v, %v) == %d, want %d",
				mutations, c.molecule, got, c.want)
		}
	}
}
