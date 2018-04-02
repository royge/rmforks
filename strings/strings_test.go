package strings

import "testing"

func TestContains(t *testing.T) {
	l := []string{"abc", "def", "123", "456"}
	tt := []struct {
		s     string
		found bool
	}{
		{
			s:     "abc",
			found: true,
		},
		{
			s:     "abcd",
			found: false,
		},
		{
			s:     "def",
			found: true,
		},
		{
			s:     "123",
			found: true,
		},
		{
			s:     "1",
			found: false,
		},
		{
			s:     "456",
			found: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.s, func(t *testing.T) {
			found := Contains(l, tc.s)
			if found != tc.found {
				t.Fatalf("expected found to %v, got %v", tc.found, found)
			}
		})
	}
}
