package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {

	//1 - настройка
	testTable := []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{1, 45, 78, 75, 3},
			expected: 7,
		}, {
			numbers:  []int{47, -5, 7, 85, -22},
			expected: 85,
		}, {
			numbers:  []int{},
			expected: 7,
		},
	}

	//2 - вызов

	for _, test_case := range testTable {
		res := Max(test_case.numbers)

		t.Logf("Calling Max(%v), result %d\n", test_case.numbers, res)

		//3 - проверка результатов
		assert.Equal(t, test_case.expected, res, fmt.Sprintf("incorrect result. Expected %d, got %d", test_case.expected, res))

	}

}

func TestMaxIndex(t *testing.T) {
	//1 - настройка
	numbers := []int{1, 45, 78, 75, 3}
	expected := 2

	//2 - вызов
	res := MaxIndex(numbers)

	//3 - проверка результатов
	if res != expected {
		t.Errorf("incorrect result. Expected %d, got %d", expected, res)
	}
}
