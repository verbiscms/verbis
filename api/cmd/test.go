package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test Command",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("in test")

			//opts, err := app.store.Options.GetStruct()
			//if err != nil {
			//	log.Debug("in err")
			//	log.Error(err)
			//}
			//log.Debug(opts)


		},
	}
)
