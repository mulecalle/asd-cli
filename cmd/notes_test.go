package cmd

import (
	"asd-cli/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCommands(t *testing.T) {
	// Set up test cases
	tests := []struct {
		name       string
		domain     string
		note       string
		expected   string
		entryValue string
		setupFunc  func()
	}{
		{
			name:     "domain_and_note_are_set",
			domain:   "k8s",
			note:     "get-jobs-by-namespace",
			expected: "k get jobs -n {namespace}\n\n",
			setupFunc: func() {
				Commands = map[string]command{
					"k8s": {
						entries: map[string]interface{}{
							"get-jobs-by-namespace": "k get jobs -n {namespace}",
						},
					},
				}
			},
		},
		{
			name:     "only_domain_is_set",
			domain:   "k8s",
			expected: "delete-jobs-in-namespace: k delete job --all -n {namespace}\ndelete-pod: k delete pod {pod-name} --grace-period=0 --force -n {namespace}\nget-jobs-by-namespace: k get jobs -n {namespace}\n\n",
			setupFunc: func() {
				Commands = map[string]command{
					"k8s": {
						entries: map[string]interface{}{
							"get-jobs-by-namespace":    "k get jobs -n {namespace}",
							"delete-jobs-in-namespace": "k delete job --all -n {namespace}",
							"delete-pod":               "k delete pod {pod-name} --grace-period=0 --force -n {namespace}",
						},
					},
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Set up the test environment
			test.setupFunc()
			domain = test.domain
			note = test.note

			// Run the command and capture the output
			output := utils.CaptureOutput(runCommands)

			// Check the output
			assert.Equal(t, test.expected, output)
		})
	}
}
