package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	installCmd = &cobra.Command{
		Use:   "install",
		Short: "Runs database migrations and setups up the database.",
		Long:  `Migrate will run all the database migrations, IMplement`,
		Run: func(cmd *cobra.Command, args []string) {

			//err := doctorCmd.Execute(); if err != nil {
			//	fmt.Println(err)
			//	return
			//}

			fmt.Println("in here")

			//if err := app.db.Install(); err != nil {
			//	color.Red.Println("->" + err.Error())
			//	return
			//}

		},
	}
)

func init() {

}
