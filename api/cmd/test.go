package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {

			doctor()
			//
			//log.WithFields(log.Fields{
			//	"code": errors.INTERNAL,
			//	"message": "Test",
			//	"operation": "op",
			//	"error": fmt.Errorf("This is the error"),
			//}).Panic()

			log.WithFields(log.Fields{
				"error": errors.Error{Code: errors.INTERNAL, Message: "Test", Operation: "op", Err: fmt.Errorf("dfdasfdasf")},
			}).Error()

		},
	}
)



// CreateUser creates a new user in the system.
// Returns INVALID if the username is blank or already exists.
// Returns CONFLICT if the username is already in use.






