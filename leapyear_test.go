package leapyear

import "testing"

func TestIsLeapYear(t *testing.T) {
	testCases := []struct {
		year     int
		expected bool
	}{
		{2000, true},
		{2400, true},
		{1800, false},
		{2100, false},
		{2023, false},
		{2024, true},
		{1, false},
		{0, true},
	}

	for _, testCase := range testCases {
		if result := Is(testCase.year); result != testCase.expected {
			t.Errorf("Is(%d) = %v, want %v", testCase.year, result, testCase.expected)
		}
	}
}
