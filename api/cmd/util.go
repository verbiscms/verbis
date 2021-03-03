// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/gookit/color"
	"os"
	"regexp"
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

// Print error to terminal with no exit
func printErrorNoExit(msg string) {
	fmt.Println()
	errMsg := color.New(color.LightWhite, color.BgRed, color.OpBold)
	errMsg.Print(" ERROR ")
	fmt.Print(" ")
	color.Red.Print(msg)
	fmt.Println()
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
	fmt.Println()
	s := spinner.New(spinner.CharSets[14], 50*time.Millisecond) //nolint
	s.Suffix = " " + msg
	s.Start()
	fmt.Printf("\n")
	s.Stop()
}

// isEmailValid checks if the email provided passes the required structure and length.
//nolint
func isEmailValid(e string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
