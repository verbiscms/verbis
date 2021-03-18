// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"testing"
)

//nolint
var wg sync.WaitGroup
var errCountPhoto int
var errCountPage int
var photoErr error

func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}

func Test_Hammer(t *testing.T) {
	skipCI(t)
	t.Skip("Skipping hammer, manual test")

	conn := 2000

	for i := 0; i < conn; i++ {
		wg.Add(1)

		//time.Sleep(time.Millisecond * 1)
		//go runHammer(i)

		// Needs to be removed here!
		wg.Done()
	}

	wg.Wait()

	fmt.Println("-------- PHOTO --------")
	fmt.Println(errCountPhoto)
	fmt.Println("-------- PAGE --------")
	fmt.Println(errCountPage)
	fmt.Println("-------- ERROR --------")
	fmt.Println(photoErr)
}

//nolint
func runHammer(i int) {
	defer func() {
		wg.Done()
	}()

	fmt.Println(i)

	//req, err := http.Get("https://staging.reddico.co.uk")
	req, err := http.Get("https://staging.reddico.co.uk/uploads/2021/02/untitled-design-4.png")
	if err != nil {
		fmt.Println("NUMBER: ")
		fmt.Println(i)
		photoErr = err
		errCountPhoto++
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	reqHtml, err := http.Get("https://staging.reddico.co.uk")
	//reqHtml, err := http.Get("http://127.0.0.1:8080/casestudies/case-study")
	if err != nil {
		fmt.Println(err)
		errCountPage++
		return
	}

	bodyHtml, err := ioutil.ReadAll(reqHtml.Body)
	if err != nil {
		fmt.Println(err)
	}

	var b = body
	var bHtml = bodyHtml
	_ = fmt.Errorf("%s%s", b, bHtml)
}
