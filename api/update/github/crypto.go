// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func ChecksumFileSHA256(path string) string {
	fmt.Println(path)
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer f.Close()

	h := sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return ""
	}

	/// THISSSS
	return hex.EncodeToString(h.Sum(nil))
}

func ChecksumFileSHA256Bytes(path string) []byte {
	fmt.Println(path)
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	h := sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return nil
	}

	return h.Sum(nil)
}
