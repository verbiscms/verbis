// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/importer/wordpress"
	"github.com/kyokomi/emoji"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	importCmd = &cobra.Command{
		Use:   "import",
		Short: "Import XML files from Wordpress and migrate content to your Verbis installation",
		Long: `This command will accept an XML file from a Wordpress installation
and convert the data into Verbis content. `,
		Run: func(cmd *cobra.Command, args []string) {

			// Run doctor
			cfg, _, err := doctor(false)
			if err != nil {
				printError(err.Error())
			}

			fmt.Println()

			file := getXMLFile()

			wp, err := wordpress.New(file, cfg.Store, true)
			if err != nil {
				printError(err.Error())
			}

			//"/Users/ainsley/Desktop/Reddico/websites/reddico-website/theme/res/import-xml/test.xml"

			wp.Import()
		},
	}
)

func getXMLFile() string {

	emoji.Println(":backhand_index_pointing_right: Enter the absolute path of the XML file to be imported")
	fmt.Println()

	promptXML := promptui.Prompt{
		Label: "XML File",
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("Enter a the XML file path")
			}
			return nil
		},
	}

	xmlFile, err := promptXML.Run()
	if err != nil {
		printError(fmt.Sprintf("Install failed: %v\n", err))
	}

	return xmlFile
}
