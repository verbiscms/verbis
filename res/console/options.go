package console

import  (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/gookit/color"
)

// Welcome message
func Welcome() {
	figure := figure.NewColorFigure("Reddico CMS", "cybermedium", "reset", true)
	figure.Print()
	fmt.Printf("\n%s version %s\n", color.Green.Render("Reddico CMS"), color.Yellow.Render("0.0.1"))
	fmt.Printf("\n")
}

// Display options/help
func Options() {

	color.Yellow.Println("- Options")

	// Help
	color.Green.Print("    help")
	fmt.Print("\t\tDisplay this help message\n")
	color.Green.Print("    exit")
	fmt.Print("\t\tExit the shell\n\n")

	color.Yellow.Println("- Available Commands")

	// Serve
	color.Yellow.Println("serve")
	color.Green.Print("    serve\t\t")
	fmt.Print("Serve the CMS\n")
	color.Green.Print("    serve:live\t\t")
	fmt.Print("Serve the CMS with Gin\n")

	// Make
	color.Yellow.Println("make")
	color.Green.Print("    make:controller\t")
	fmt.Print("Create a new controller instance\n")
	color.Green.Print("    make:migration\t")
	fmt.Print("Create a new migration file\n")
	color.Green.Print("    make:model\t")
	fmt.Print("\tCreate a new model instance\n")

	// Migrate
	color.Yellow.Println("migrate")
	color.Green.Print("    migrate\t\t")
	fmt.Print("Run all database migrations\n")
	color.Green.Print("    migrate:fresh\t")
	fmt.Print("Drop all tables and re-run all migrations\n")
	color.Green.Print("    migrate:fresh\t")
	fmt.Print("Drop all tables and re-run all migrations\n")

	fmt.Println()
}
