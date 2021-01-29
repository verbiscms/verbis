package cast

// ToSlice
//
// Casts an interface to a []interface{} type.
//
// Example: {{ toSlice 1 }}
// Returns: `[1]`
func (ns *Namespace) ToSlice(i interface{}) []interface{} {
	var s []interface{}

	if i == nil {
		return nil
	}

	switch v := i.(type) {
	case []interface{}:
		return append(s, v...)
	case []map[string]interface{}:
		for _, u := range v {
			s = append(s, u)
		}
		return s
	default:
		s = append(s, i)
		return s
	}
}
