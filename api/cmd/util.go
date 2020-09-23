package cmd

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/gookit/color"
	"os"
	"time"
)

// Print error to terminal
func printError(msg string) {
	fmt.Println()
	errMsg := color.New(color.LightWhite, color.BgRed, color.OpBold)
	errMsg.Print(" ERROR ")
	fmt.Print(" ")
	color.Red.Print(msg)
	fmt.Println()
	os.Exit(1)
}

// Print success to terminal
func printSuccess(msg string) {
	fmt.Println()
	successMsg := color.New(color.LightWhite, color.BgGreen)
	successMsg.Print(" SUCCESS ")
	fmt.Print(" ")
	color.Green.Print(msg)
	fmt.Println()
}

// Print spinner
func printSpinner(msg string) {
	fmt.Printf("\n")
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " " + msg
	s.Start()
	time.Sleep(1 * time.Second)
	fmt.Printf("\n")
}