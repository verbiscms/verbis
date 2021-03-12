// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package options

// Find
//
// Returns a option by searching with the given name.
// Returns errors.INTERNAL if there was an error executing the query.
// Returns errors.NOTFOUND if the option was not found by the given name.
func (s *Store) Find(name string) (interface{}, error) {
	const op = "OptionStore.Find"

	//q := s.Builder().
	//	From(s.Schema()+TableName).
	//	Where("id", "=", id).
	//	Limit(1)
	//
	//var category domain.Category
	//err := s.DB().Get(&category, q.Build())
	//if err == sql.ErrNoRows {
	//	return domain.Category{}, &errors.Error{Code: errors.NOTFOUND, Message: fmt.Sprintf("No category exists with the ID: %d", id), Operation: op, Err: err}
	//} else if err != nil {
	//	return domain.Category{}, &errors.Error{Code: errors.INTERNAL, Message: database.ErrQueryMessage, Operation: op, Err: err}
	//}

	return nil, nil
}
