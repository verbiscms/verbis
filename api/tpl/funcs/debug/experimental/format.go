// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package experimental

//type formatState struct {
//	value          interface{}
//	wrt            *bytes.Buffer
//	depth          int
//	pointers       map[uintptr]int
//	ignoreNextType bool
//}
//
//func newFormatState(i interface{}) *formatState {
//	return &formatState{
//		value:    i,
//		wrt:      new(bytes.Buffer),
//		depth:    0,
//		pointers: make(map[uintptr]int),
//	}
//}
//
//func (f *formatState) Get() string {
//
//	f.formatPtr(reflect.ValueOf(f.value))
//
//	return f.wrt.String()
//}
//
//// formatPtr handles formatting of pointers by indirecting them as necessary.
//func (f *formatState) formatPtr(v reflect.Value) {
//	// Display nil if top level pointer is nil.
//	if v.IsNil() && f.ignoreNextType {
//		f.wrt.Write(debug.nilAngleBytes)
//		return
//	}
//
//	// Remove pointers at or below the current depth from map used to detect
//	// circular refs.
//	for k, depth := range f.pointers {
//		if depth >= f.depth {
//			delete(f.pointers, k)
//		}
//	}
//
//	// Keep list of all dereferenced pointers to possibly show later.
//	pointerChain := make([]uintptr, 0)
//
//	// Figure out how many levels of indirection there are by derferencing
//	// pointers and unpacking interfaces down the chain while detecting circular
//	// references.
//	nilFound := false
//	cycleFound := false
//	indirects := 0
//	ve := v
//	for ve.Kind() == reflect.Ptr {
//		if ve.IsNil() {
//			nilFound = true
//			break
//		}
//		indirects++
//		addr := ve.Pointer()
//		pointerChain = append(pointerChain, addr)
//		if pd, ok := f.pointers[addr]; ok && pd < f.depth {
//			cycleFound = true
//			indirects--
//			break
//		}
//		f.pointers[addr] = f.depth
//
//		ve = ve.Elem()
//		if ve.Kind() == reflect.Interface {
//			if ve.IsNil() {
//				nilFound = true
//				break
//			}
//			ve = ve.Elem()
//		}
//	}
//
//	// Display type or indirection level depending on flags.
//	if nilFound || cycleFound {
//		indirects += strings.Count(ve.Type().String(), "*")
//	}
//	f.wrt.Write(debug.openAngleBytes)
//	f.wrt.Write([]byte(strings.Repeat("*", indirects)))
//	f.wrt.Write(debug.closeAngleBytes)
//
//	// Display dereferenced value.
//	switch {
//	case nilFound:
//		f.wrt.Write(debug.nilAngleBytes)
//
//	case cycleFound:
//		f.wrt.Write(debug.circularShortBytes)
//
//	default:
//		f.ignoreNextType = true
//		f.Format(ve)
//	}
//}
//
//func (f *formatState) Format(v reflect.Value) {
//	kind := v.Kind()
//
//	// Handle pointers
//	if kind == reflect.Ptr {
//		//f.formatPtr(v)
//		return
//	}
//
//	switch kind {
//	case reflect.String:
//		f.wrt.Write([]byte(v.String()))
//	case reflect.Invalid:
//		f.wrt.Write(debug.invalidAngleBytes)
//	case reflect.Bool:
//		debug.printBool(f.wrt, v.Bool())
//
//	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
//		debug.printInt(f.wrt, v.Int(), 10)
//
//	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
//		debug.printUint(f.wrt, v.Uint(), 10)
//
//	case reflect.Float32:
//		debug.printFloat(f.wrt, v.Float(), 32)
//
//	case reflect.Float64:
//		debug.printFloat(f.wrt, v.Float(), 64)
//
//	case reflect.Complex64:
//		debug.printComplex(f.wrt, v.Complex(), 32)
//
//	case reflect.Complex128:
//		debug.printComplex(f.wrt, v.Complex(), 64)
//
//	//case reflect.Slice:
//	//	if v.IsNil() {
//	//		f.wrt.Write(nilAngleBytes)
//	//		break
//	//	}
//	//	fallthrough
//	//
//	//case reflect.Array:
//	//	f.wrt.Write(openBracketBytes)
//	//	f.depth++
//	//	numEntries := v.Len()
//	//
//	//	for i := 0; i < numEntries; i++ {
//	//		if i > 0 {
//	//			f.wrt.Write(spaceBytes)
//	//		}
//	//		f.Format(f.unpackValue(v.Index(i)))
//	//	}
//	//
//	//	f.depth--
//	//	f.wrt.Write(closeBracketBytes)
//	//
//	//case reflect.Interface:
//	//	// The only time we should get here is for nil interfaces due to
//	//	// unpackValue calls.
//	//	if v.IsNil() {
//	//		f.wrt.Write(nilAngleBytes)
//	//	}
//	//
//	//case reflect.Ptr:
//	//	// Do nothing.  We should never get here since pointers have already
//	//	// been handled above.
//	//
//	case reflect.Map:
//		// nil maps should be indicated as different than empty maps
//		if v.IsNil() {
//			f.wrt.Write(debug.nilAngleBytes)
//			break
//		}
//
//		f.wrt.Write(debug.openMapBytes)
//		f.depth++
//		keys := v.MapKeys()
//
//		// Sorting herte
//
//		for i, key := range keys {
//			if i > 0 {
//				f.wrt.Write(debug.spaceBytes)
//			}
//			f.Format(f.unpackValue(key))
//			f.wrt.Write(debug.colonBytes)
//			f.Format(f.unpackValue(v.MapIndex(key)))
//		}
//
//		f.depth--
//		f.wrt.Write(debug.closeMapBytes)
//
//		//case reflect.Struct:
//		//	numFields := v.NumField()
//		//	f.wrt.Write(openBraceBytes)
//		//	f.depth++
//		//
//		//	for i := 0; i < numFields; i++ {
//		//		if i > 0 {
//		//			f.wrt.Write(spaceBytes)
//		//		}
//		//		f.Format(f.unpackValue(v.Field(i)))
//		//	}
//		//
//		//	f.depth--
//		//	f.wrt.Write(closeBraceBytes)
//		//
//		//case reflect.Uintptr:
//		//	printHexPtr(f.wrt, uintptr(v.Uint()))
//		//
//		//case reflect.UnsafePointer, reflect.Chan, reflect.Func:
//		//	printHexPtr(f.wrt, v.Pointer())
//		//}
//	}
//}
//
//// unpackValue returns values inside of non-nil interfaces when possible and
//// ensures that types for values which have been unpacked from an interface
//// are displayed when the show types flag is also set.
//// This is useful for data types like structs, arrays, slices, and maps which
//// can contain varying types packed inside an interface.
//func (f *formatState) unpackValue(v reflect.Value) reflect.Value {
//	if v.Kind() == reflect.Interface {
//		if !v.IsNil() {
//			v = v.Elem()
//		}
//	}
//	return v
//}
//
//func sortValues(m map[string]interface{}) []string {
//	keys := make([]string, len(m))
//	i := 0
//
//	for k := range m {
//		keys[i] = k
//		i++
//	}
//	sort.Strings(keys)
//
//	return keys
//}
