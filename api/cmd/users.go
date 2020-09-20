package cmd

import (
	"bufio"
	"github.com/ainsleyclark/verbis/api/domain"
	validation "github.com/ainsleyclark/verbis/api/helpers/vaidation"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

var (
	v validation.Validator
	usersCmd = &cobra.Command{
		Use:   "users",
		Short: "Access & modify the CMS users.",
		Long: `Calling users will enable you to add, update and delete users,
you will be able to assign role ids, and modify existing users that have
been stored in the cms database.`,
	}
)

// Add child commands
func init() {
	v = validation.New()
	usersCmd.AddCommand(listUsersCmd)
	usersCmd.AddCommand(createUserCmd)
}

// listCmd represents the list command
var listUsersCmd = &cobra.Command{
	Use:   "list",
	Short: "List's all users in the database.",
	Run: func(cmd *cobra.Command, args []string) {

		//users, err := app.store.User.GetAll()
		//if err != nil {
		//	fmt.Print(err)
		//}
		//
		//t := table.NewWriter()
		//t.SetOutputMirror(os.Stdout)
		//t.AppendHeader(table.Row{"# ID", "First Name", "Last Name", "Email", "Email Verified At", "Created At", "Updated At"})
		//
		//for _, v := range users {
		//	t.AppendRows([]table.Row{
		//		{v.Id, v.FirstName, v.LastName, v.Email, v.EmailVerifiedAt, v.CreatedAt, v.UpdatedAt},
		//	})
		//}
		//
		//t.Render()
	},
}

// createCmd represents the create command
var createUserCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a user in the database.",
	RunE: func(cmd *cobra.Command, args []string) error {

		reader := bufio.NewReader(os.Stdin)
		var user domain.User

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

		// Access Level
		accessLevel := ""
		for {
			fmt.Println("User a number to select an access level")
			fmt.Println("0) Banned")
			fmt.Println("1) Operator")
			fmt.Println("2) Administrator")
			fmt.Print("Access Level: ")
			accessLevel, _ = reader.ReadString('\n')

			_, err := strconv.Atoi(strings.TrimSuffix(accessLevel, "\n"))
			if err != nil {
				return err
			}

			//user.Role.Id = l
			err = v.CmdCheck("accesslevel", user)
			if  err != nil {
				fmt.Println(err)
			} else {
				break
			}
		}

		// Create
		_, err := app.store.User.Create(&user)
		if err != nil {
			return err
		}


		fmt.Println(user.FirstName + " created successfully!")

		return nil
	},
}