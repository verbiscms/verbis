// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run ../generator/main.go

package updates

import (
	sm "github.com/hashicorp/go-version"
)

type UpdateCallBackFn func() error

type registry []*Update

type Update struct {
	Version   string
	Migration string
	Callback  UpdateCallBackFn
}

var UpdateRegistry = make(registry, 0)

func (r *registry) Add(update Update) {
	*r = append(*r, &update)
}

func (r *registry) Get(version string) *Update {
	for _, v := range *r {
		if v.Version == version {
			return v
		}
	}
	return nil
}

func (u Update) ToSemVer() (*sm.Version, error) {
	return sm.NewVersion(u.Version)
}

func test() {

	//for _, update := range u {
	//	ver, err := update.ToSemVer()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	needsUpdate := ver.LessThan(version.SemVer)
	//	if !needsUpdate {
	//		return
	//	}
	//
	//	fmt.Println("hello")
	//}
}
