package cmd

import (
	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)



// CreateUser creates a new user in the system.
// Returns INVALID if the username is blank or already exists.
// Returns CONFLICT if the username is already in use.






