package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"Zero size", 0},
		{"Size 1", 1},
		{"Size 10", 10},
		{"Size 100", 100},
	}

	for _, g := range tests {
		t.Run(g.name, func(t *testing.T) {
			res := generateRandomElements(g.size)
			if len(res) != g.size {
				t.Errorf("Ожидалась длина %d, получено %d", g.size, len(res))
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expexted int
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{1}, 1},
		{"All positive", []int{1, 2, 3, 4, 5}, 5},
		{"All negative", []int{-1, -2, -3, -4, -5}, -1},
		{"Mixed", []int{1, -2, 3, -4, 5}, 5},
		{"Max at beginning", []int{5, 4, 3, 2, 1}, 5},
		{"Max at end", []int{1, 2, 3, 4, 5}, 5},
	}

	for _, g := range tests {
		t.Run(g.name, func(t *testing.T) {
			res := maximum(g.data)
			if res != g.expexted {
				t.Errorf("Ожидалось %d, получено %d", g.expexted, res)
			}
		})
	}
}
