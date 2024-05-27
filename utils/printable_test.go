package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintSliceString(t *testing.T) {
	tests := []struct {
		name     string
		header   string
		input    []string
		expected string
	}{
		{
			name:     "single_element",
			header:   "Header",
			input:    []string{"element1"},
			expected: "Header\nelement1\n",
		},
		{
			name:     "multiple_elements",
			header:   "Header",
			input:    []string{"element1", "element2", "element3"},
			expected: "Header\nelement1\nelement2\nelement3\n",
		},
		{
			name:     "empty_slice",
			header:   "Header",
			input:    []string{},
			expected: "Header\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := CaptureOutput(func() {
				PrintSliceString(test.header, test.input)
			})
			assert.Equal(t, test.expected, output)
		})
	}
}
