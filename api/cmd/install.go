package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/ainsleyclark/verbis/api/config"
	"github.com/ainsleyclark/verbis/api/database/seeds"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/errors"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	"github.com/ainsleyclark/verbis/api/helpers/webp"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// Add child commands
func init() {
	v = validation.New()
}

var (
	v validation.Validator
	installCmd = &cobra.Command{
		Use:   "install",
		Short: "Install will run the doctor command and then run database schema and insert any data dependant on Verbis.",
		Long:  `This command will install first run Verbis doctor to see if the database,
exists and is passable. Install will then run the migration to insert into the schema.
Seeds are also run, inserting options and any necessary configuration into the 
database.`,
		Run: Install,
	}
)

// Add child commands/init
func init() {
	v = validation.New()
}

func Install(cmd *cobra.Command, args []string) {

	// Run doctor
	db, err := doctor()
	if err != nil {
		printError(err.Error())
	}

	// Start the spinner
	printSpinner("Installing Verbis...")

	// Install the database
	if err := db.Install(); err != nil {
		fmt.Println(err)
		printError(fmt.Sprintf("A database with the name %s has already been installed. \nPlease run verbis uninstall if you want to delete it.", environment.GetDatabaseName()))
	}

	// Init Config
	con, err := config.New()
	if err != nil {
		printError(errors.Message(err))
	}

	// Set up stores & pass the database.
	store := models.New(db, *con)
	if err != nil {
		printError(err.Error())
	}

	// Run the seeds
	seeder := seeds.New(db.Sqlx, store)
	if err := seeder.Seed(); err != nil {
		printError(err.Error())
	}

	// Create the owner user
	user := createOwner()
	if _, err := store.User.Create(user); err != nil {
		printError(fmt.Sprintf("Error creating the owner", err.Error()))
	}

	// Insert the site url
	fmt.Println()
	url := setUrl()
	mUrl, _ := json.Marshal(url)
	if err := store.Options.Update("site_url", mUrl); err != nil {
		printError(fmt.Sprintf("Error not inserting the site url:", err.Error()))
	}

	// Get webp executables
	bin := webp.CreateBinWrapper()
	bin.ExecPath("cwebp")
	if err := bin.Run(); err != nil {
		printError(fmt.Sprintf("Error downloading webp executables, are you connected to the internet? ", err.Error()))
	}

	// Print success
	printSuccess("Successfully installed verbis")

	return
}

// setUrl
func setUrl() string {
	fmt.Println("Enter the url the site will sit on, if in development, be sure to append a port (for example: http://127.0.0.1:8080):")
	reader := bufio.NewReader(os.Stdin)

	// Url
	url := ""
	for {
		fmt.Print("Url: ")
		url, _ = reader.ReadString('\n')
		if  url == "" {
			fmt.Println("Enter a valid url")
		} else {
			break
		}
	}

	url = strings.TrimSuffix(url, "\n")

	return url
}

// createOwner Create's the owner of the site for the install.
func createOwner() *domain.UserCreate {
	fmt.Println("Enter the owner's details:")

	reader := bufio.NewReader(os.Stdin)
	var user domain.UserCreate

	// First name
	firstName := ""
	for {
		fmt.Print("First name: ")
		firstName, _ = reader.ReadString('\n')
		user.FirstName = strings.TrimSuffix(firstName, "\n")
		err := v.CmdCheck("firstname", user)
		if  err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	// Last name
	lastName := ""
	for {
		fmt.Print("Last name: ")
		lastName, _ = reader.ReadString('\n')
		user.LastName = strings.TrimSuffix(lastName, "\n")
		err := v.CmdCheck("lastname", user)
		if  err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	// Email
	email := ""
	for {
		fmt.Print("Email address: ")
		email, _ = reader.ReadString('\n')
		user.Email = strings.TrimSuffix(email, "\n")
		err := v.CmdCheck("email", user)
		if  err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	// Password
	password := ""
	for {
		fmt.Print("Password: ")
		password, _ = reader.ReadString('\n')
		user.Password = strings.TrimSuffix(password, "\n")
		err := v.CmdCheck("password", user)
		if  err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	// Password
	confirmPassword := ""
	for {
		fmt.Print("Confirm Password: ")
		confirmPassword, _ = reader.ReadString('\n')
		user.ConfirmPassword = strings.TrimSuffix(confirmPassword, "\n")
		if confirmPassword != password {
			fmt.Println("Password's dont match! Exit and try again.")
		} else {
			break
		}
	}
	
	user.Role = domain.UserRole{
		Id: 6,
	}

	return &user
}