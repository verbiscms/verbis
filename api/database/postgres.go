// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package database

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestPostgres(env *environment.Env) *sqlx.DB {
	//connString := "postgresql://" + env.DbHost + ":" + env.DbPort + "/" + env.DbDatabase + "?user=" + env.DbUser + "&password=" + env.DbPassword + "&statement_cache_mode=describe"
	connString := "postgresql://127.0.0.1:5432/verbis?user=ainsley&password=&sslmode=disable"

	driver, err := sqlx.Connect("postgres", connString)

	fmt.Println(err)

	return driver
}
