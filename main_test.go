package main

import "testing"

// Пишите тесты в этом файле
type testMax []struct {
	name     string
	date     []int
	expected int
}

func TestGenerateRandomElements(t *testing.T) {
	t.Run("Стандартный случай", func(t *testing.T) {
		result := generateRandomElements(100)
		if len(result) != 100 {
			t.Errorf("Ожидаемая длина слайса %d != полученной длине слайса %d", 100, len(result))
		}
	})
	t.Run("Диапазон значений", func(t *testing.T) {
		result := generateRandomElements(10)
		for _, value := range result {
			if value < 0 || value > 10 {
				t.Errorf("Ожидаемые значения в диапазоне от 0 до 9 != полученному значению %d", value)
			}
		}
	})
	t.Run("Длина слайса 0", func(t *testing.T) {
		result := generateRandomElements(0)
		if len(result) != 0 {
			t.Errorf("Ожидаемая длина слайса 0 != полученному значению %d", len(result))
		}
	})
	t.Run("Слишком большой размер", func(t *testing.T) {
		result := generateRandomElements(100_000_001)
		if len(result) != 0 {
			t.Errorf("Ожидаемая длина слайса 0 != полученному значению %d", len(result))
		}
	})
}

func TestMaximum(t *testing.T) {
	tests := testMax{
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
				t.Errorf("Получено %d != ожидалось %d", result, test.expected)
			}
		})
	}
}
