// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/database/seeds"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/kyokomi/emoji"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"net/url"
)

var (
	installCmd = &cobra.Command{
		Use:   "install",
		Short: "Install will run the doctor command and then run database schema and insert any data dependant on Verbis.",
		Long: `This command will install first run Verbis doctor to see if the database,
exists and is passable. Install will then run the migration to insert into the schema.
Seeds are also run, inserting options and any necessary configuration into the 
database.`,
		Run: Install,
	}
)

func Install(cmd *cobra.Command, args []string) {
	// Run doctor
	cfg, db, err := doctor(false)
	if err != nil {
		printError(err.Error())
	}
	d := deps.New(*cfg)

	// Check if the database exists.
	// TODO NOT WORKING
	//err = db.CheckExists()
	//if err != nil {
	//	printError(fmt.Sprintf("A database with the name %s has already been installed. \nPlease run verbis uninstall if you want to delete it.", cfg.Env.DbDatabase))
	//}

	// Get the user & site variables
	user := createOwner()
	fmt.Println()
	uri := setURL()

	// Start the spinner
	printSpinner("Installing Verbis...")

	// Install the database
	err = db.Install()
	if err != nil {
		printError(fmt.Sprintf("Error installing the Verbis database: %v", err))
	}

	// Run the seeds
	seeder := seeds.New(db, cfg.Store)
	if err := seeder.Seed(); err != nil {
		printError(err.Error())
	}

	// Create the owner user
	_, err = d.Store.User.Create(*user)
	if err != nil {
		printError(fmt.Sprintf("Error creating the owner: %s", err.Error()))
	}

	// Insert the site uri
	fmt.Println()
	err = d.Store.Options.Update("site_url", uri)
	if err != nil {
		printError(fmt.Sprintf("Error not inserting the site uri: %s", err.Error()))
	}

	// Print success
	printSuccess("Successfully installed verbis")
}

// setURL
func setURL() string {
	emoji.Println(":backhand_index_pointing_right: Enter the url will sit on:")
	fmt.Println("If in development, be sure to append a port (for example: http://127.0.0.1:8080):")

	prompt := promptui.Prompt{
		Label: "URL",
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("enter url")
			}
			_, err := url.ParseRequestURI("http://google.com/")
			if err != nil {
				return fmt.Errorf("enter a valid url")
			}
			return nil
		},
	}
	homeURL, err := prompt.Run()
	if err != nil {
		printError(fmt.Sprintf("Install failed: %v\n", err))
	}

	return homeURL
}

// createOwner Create's the owner of the site for the install.
func createOwner() *domain.UserCreate {
	emoji.Print(":backhand_index_pointing_right: Enter the owner's details:")

	promptFirstName := promptui.Prompt{
		Label: "First name",
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("enter a first name")
			}
			return nil
		},
	}
	firstName, err := promptFirstName.Run()
	if err != nil {
		printError(fmt.Sprintf("Install failed: %v\n", err))
	}

	promptLastName := promptui.Prompt{
		Label: "Last name",
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("enter a last name")
			}
			return nil
		},
	}
	lastName, err := promptLastName.Run()
	if err != nil {
		printError(fmt.Sprintf("install failed: %v\n", err))
	}

	promptEmail := promptui.Prompt{
		Label: "Email",
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("enter a email address")
			}
			if !isEmailValid(input) {
				return fmt.Errorf("enter a valid email address")
			}
			return nil
		},
	}
	email, err := promptEmail.Run()
	if err != nil {
		printError(fmt.Sprintf("Install failed: %v\n", err))
	}

	promptPassword := promptui.Prompt{
		Label: "Password",
		Validate: func(input string) error {
			if len(input) < 8 { //nolint
				return fmt.Errorf("password must have more than 8 characters")
			}
			return nil
		},
		Mask: '*',
	}
	password, err := promptPassword.Run()
	if err != nil {
		printError(fmt.Sprintf("Install failed: %v\n", err))
	}

	prompConfirmPassword := promptui.Prompt{
		Label: "Password",
		Validate: func(input string) error {
			if input != password {
				return fmt.Errorf("password and confirm password must match")
			}
			return nil
		},
		Mask: '*',
	}
	confirmPassword, err := prompConfirmPassword.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		printError(fmt.Sprintf("Install failed: %v\n", err))
	}

	user := domain.UserCreate{
		User: domain.User{
			UserPart: domain.UserPart{
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Role: domain.Role{
					Id: domain.OwnerRoleID,
				},
			},
		},
		Password:        password,
		ConfirmPassword: confirmPassword,
	}

	return &user
}
