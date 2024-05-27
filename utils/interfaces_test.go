package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapValueToYAMLString(t *testing.T) {
	tests := []struct {
		name           string
		input          interface{}
		expected       string
		expectingError bool
	}{
		{
			name: "simple_ map",
			input: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			expected:       "key1: value1\nkey2: value2\n",
			expectingError: false,
		},
		{
			name: "nested_map",
			input: map[string]interface{}{
				"key1": map[string]interface{}{
					"subkey1": "subvalue1",
				},
				"key2": "value2",
			},
			expected:       "key1:\n    subkey1: subvalue1\nkey2: value2\n",
			expectingError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := MapValueToYAMLString(test.input)
			if test.expectingError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, output)
			}
		})
	}
}

func TestSortEntries(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]interface{}
	}{
		{
			name: "simple_map",
			input: `
key2: value2
key1: value1
`,
			expected: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "nested_map",
			input: `
key2: value2
key1:
  subkey1: subvalue1
`,
			expected: map[string]interface{}{
				"key1": map[string]interface{}{
					"subkey1": "subvalue1",
				},
				"key2": "value2",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := SortEntries([]byte(test.input))
			assert.Equal(t, test.expected, output)
		})
	}
}

func TestSortMapKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected []string
	}{
		{
			name: "simple_map",
			input: map[string]interface{}{
				"key2": "value2",
				"key1": "value1",
			},
			expected: []string{"key1", "key2"},
		},
		{
			name: "unordered_map",
			input: map[string]interface{}{
				"b": "valueB",
				"a": "valueA",
				"c": "valueC",
			},
			expected: []string{"a", "b", "c"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := sortMapKeys(test.input)
			assert.Equal(t, test.expected, output)
		})
	}
}
