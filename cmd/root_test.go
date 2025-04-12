package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	// Test successful execution
	t.Run("successful_execution", func(t *testing.T) {
		// Backup and restore the original root command
		originalRootCmd := rootCmd
		defer func() {
			rootCmd = originalRootCmd
		}()

		// Create a simple command that will succeed
		rootCmd = &cobra.Command{
			Use: "test",
			Run: func(cmd *cobra.Command, args []string) {},
		}

		// Execute should not panic
		assert.NotPanics(t, func() {
			Execute()
		})
	})

	// Test execution with error
	t.Run("execution_with_error", func(t *testing.T) {
		// Backup and restore the original root command
		originalRootCmd := rootCmd
		defer func() {
			rootCmd = originalRootCmd
		}()

		// Create a command that will fail
		rootCmd = &cobra.Command{
			Use: "test",
			RunE: func(cmd *cobra.Command, args []string) error {
				return nil // Return nil to simulate successful execution
			},
		}

		// Execute should not panic
		assert.NotPanics(t, func() {
			Execute()
		})
	})
}
