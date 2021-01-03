package fields

// GetField
//
// Returns the value of a specific field.
// Returns errors.NOTFOUND if the field was not found by the given key.
func (s *Service) GetField(name string, args ...interface{}) (interface{}, error) {
	fields, format := s.handleArgs(args)

	field, err := s.findFieldByName(name, fields)
	if err != nil {
		return nil, err
	}

	if !format {
		return field.Value, nil
	}

	return s.resolveField(field).Value, nil
}


//// GetFieldLayout
////
//func (s *Service) GetFieldLayout(name string, args ...interface{}) (*domain.Field, error) {
//	layout, err := walker.ByName(name, s.layout)
//	if err != nil {
//		return nil, err
//	}
//
//	return &layout, nil
//}