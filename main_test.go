package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
type tests []struct {
	name     string
	date     []int
	size     int
	expected int
}

func TestGenerateRandomElements(t *testing.T) {
	tests := tests{
		{name: "Стандартный случай",
			size:     100,
			expected: 100,
		},
		{name: "Длина слайса 0",
			size:     0,
			expected: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := generateRandomElements(test.size)
			assert.Len(t, result, test.expected)
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := tests{
		{
			name:     "Стандартный случай",
			date:     []int{1, 100, 200, 500, 0},
			expected: 500,
		},
		{
			name:     "Пустой слайс",
			date:     []int{},
			expected: 0,
		},
		{
			name:     "Слайс с одним элементом",
			date:     []int{100},
			expected: 100,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := maximum(test.date)
			if result != test.expected {
				assert.Equal(t, result, test.expected)
			}
		})
	}
}
